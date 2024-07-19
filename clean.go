package fttl

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

func (p *DB) ttlInit() {
	go p.tickClean()
}

func (p *DB) tickClean() {
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		case <-p.done:
			ticker.Stop()
			p.saveConfig()
			p.wait.Done()

			return
		case <-ticker.C:
			now := time.Now()

			p.lock.Lock()

			for sum, val := range p.cfg {
				if val.Expire.Before(now) {
					_ = p.delete(sum)
				}
			}

			p.lock.Unlock()
			p.saveConfig()
		}
	}
}

func (p *DB) loadConfig() map[uint64]*Policy {
	cfg := filepath.Join(p.dir, "fttl.json")
	ret := map[uint64]*Policy{}

	if data, err := os.ReadFile(cfg); err == nil {
		_ = json.Unmarshal(data, &ret)
	}

	return ret
}

func (p *DB) saveConfig() {
	p.lock.RLock()
	defer p.lock.RUnlock()

	cfg := filepath.Join(p.dir, "fttl.json")
	if len(p.cfg) == 0 {
		_ = os.Remove(cfg)

		return
	}
	// nolint: errchkjson
	data, _ := json.MarshalIndent(p.cfg, "\t", "\t")

	_ = os.WriteFile(cfg, data, _fileMode)
}
