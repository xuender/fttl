package fttl

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	_fileMode   os.FileMode = 0o664
	_DirFileMod os.FileMode = 0o771
)

type DB struct {
	dir  string
	cfg  map[uint64]*Policy
	lock sync.RWMutex
	once sync.Once
	done chan struct{}
	wait sync.WaitGroup
}

func New(dir string) *DB {
	ret := &DB{dir: dir, done: make(chan struct{})}
	ret.cfg = ret.loadConfig()

	return ret
}

func (p *DB) Get(key []byte) ([]byte, error) {
	num := Hash(key)

	p.lock.RLock()

	item, has := p.cfg[num]
	p.lock.RUnlock()

	if has {
		now := time.Now()

		if now.After(item.Expire) {
			p.lock.Lock()
			defer p.lock.Unlock()

			_ = p.delete(num)

			return nil, ErrNotFound
		}

		if item.Access > 0 {
			expire := now.Add(item.Access)
			if expire.After(item.Expire) {
				item.Expire = expire
			}
		}
	}

	path := p.path(num)

	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return os.ReadFile(path)
	}

	if err == nil {
		return nil, ErrIsDir
	}

	return nil, err
}

func (p *DB) Put(key, value []byte) error {
	_, err := p.put(key, value)

	return err
}

func (p *DB) PutTTL(key, value []byte, expire, access time.Duration) error {
	sum, err := p.put(key, value)
	if err != nil {
		return err
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	p.cfg[sum] = &Policy{
		Expire: time.Now().Add(expire),
		Access: access,
	}

	p.once.Do(p.ttlInit)

	return nil
}

func (p *DB) Delete(key []byte) error {
	num := Hash(key)

	p.lock.Lock()
	defer p.lock.Unlock()

	return p.delete(num)
}

func (p *DB) Close() {
	p.wait.Add(1)
	close(p.done)
	p.wait.Wait()
}

func (p *DB) delete(num uint64) error {
	path := p.path(num)

	delete(p.cfg, num)

	err := os.Remove(path)
	dir := filepath.Dir(path)

	if info, err := os.ReadDir(dir); err == nil {
		if len(info) == 0 {
			_ = os.Remove(dir)
		}
	}

	return err
}

func (p *DB) path(num uint64) string {
	dir, base := Path(num)

	return filepath.Join(p.dir, dir, base)
}

func (p *DB) put(key, value []byte) (uint64, error) {
	num := Hash(key)
	path := p.path(num)
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, _DirFileMod); err != nil {
			return num, err
		}
	}

	return num, os.WriteFile(path, value, _fileMode)
}
