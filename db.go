package fttl

import (
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	_fileMode os.FileMode = 0o664
	_dirMode  os.FileMode = 0o771
)

type DB struct {
	dir   string
	cfg   map[uint64]*Policy
	lock  sync.RWMutex
	once  sync.Once
	done  chan struct{}
	wait  sync.WaitGroup
	isRun bool
}

func New(dir string) *DB {
	ret := &DB{dir: dir, done: make(chan struct{})}
	ret.cfg = ret.loadConfig()

	return ret
}

func (p *DB) Get(key []byte) ([]byte, error) {
	return p.GetByHash(Hash(key))
}

func (p *DB) GetByHash(hash uint64) ([]byte, error) {
	if err := p.Refresh(hash); err != nil {
		return nil, err
	}

	return p.GetByPath(p.path(hash))
}

func (p *DB) Refresh(hash uint64) error {
	p.lock.RLock()

	item, has := p.cfg[hash]
	p.lock.RUnlock()

	if !has {
		return nil
	}

	now := time.Now()

	if now.After(item.Expire) {
		p.lock.Lock()
		defer p.lock.Unlock()

		_ = p.delete(hash)

		return ErrNotFound
	}

	if item.Access > 0 {
		expire := now.Add(item.Access)
		if expire.After(item.Expire) {
			item.Expire = expire
		}
	}

	return nil
}

func (p *DB) GetByPath(path string) ([]byte, error) {
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return os.ReadFile(path)
	}

	if err == nil {
		return nil, ErrIsDir
	}

	return nil, err
}

func (p *DB) Put(key, value []byte) *Result {
	num := Hash(key)
	path := p.path(num)
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); err != nil {
		if err := os.MkdirAll(dir, _dirMode); err != nil {
			return &Result{Error: err, Path: path, Hash: num}
		}
	}

	return &Result{Error: os.WriteFile(path, value, _fileMode), Path: path, Hash: num}
}

func (p *DB) PutTTL(key, value []byte, expire, access time.Duration) *Result {
	res := p.Put(key, value)
	if res.Error != nil {
		return res
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	p.cfg[res.Hash] = &Policy{
		Expire: time.Now().Add(expire),
		Access: access,
	}

	p.once.Do(p.ttlInit)

	return res
}

func (p *DB) Delete(key []byte) error {
	num := Hash(key)

	p.lock.Lock()
	defer p.lock.Unlock()

	return p.delete(num)
}

func (p *DB) Has(key []byte) bool {
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

			return false
		}
	}

	path := p.path(num)
	info, err := os.Stat(path)

	return err == nil && !info.IsDir()
}

func (p *DB) Close() {
	if p.isRun {
		p.wait.Add(1)
		close(p.done)
		p.wait.Wait()
	}
}

func (p *DB) delete(num uint64) error {
	path := p.path(num)

	delete(p.cfg, num)

	err := os.Remove(path)
	p.removeDir(filepath.Dir(path))

	return err
}

func (p *DB) removeDir(path string) {
	if path == p.dir {
		return
	}

	items, err := os.ReadDir(path)
	if err != nil {
		return
	}

	if len(items) > 0 {
		return
	}

	if err := os.Remove(path); err != nil {
		return
	}

	p.removeDir(filepath.Dir(path))
}

func (p *DB) path(num uint64) string {
	dir, base := Path(num)

	return filepath.Join(p.dir, dir, base)
}
