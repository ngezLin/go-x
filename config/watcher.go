package config

import (
	"context"
	"log"
	"sync"
	"time"
)

// type Watcher struct {
// 	value   map[string]interface{}
// 	channel chan []string
// }

// func NewWatcher(keys []string, cfg Config) Watcher {
// 	cfgMap := make(map[string]interface{})
// 	for _, key := range keys {
// 		cfgMap[key] = cfg.Get(key)
// 	}

// 	return Watcher{value: cfgMap, channel: make(chan []string)}
// }

// func (w Watcher) Change() <-chan []string {
// 	return w.channel
// }

// func (w Watcher) Update(cfg Config) {
// 	change := false
// 	var keys []string
// 	for key, val := range w.value {
// 		if newVal := cfg.Get(key); newVal != val {
// 			w.value[key] = newVal
// 			keys = append(keys, key)
// 			change = true
// 		}
// 	}

// 	if change {
// 		w.channel <- keys
// 	}
// }

// func (w Watcher) Close() {
// 	close(w.channel)
// }

type Watcher struct {
	Load         func(ctx context.Context) (interface{}, error)
	mutex        sync.RWMutex
	refreshTimer *time.Ticker
	cancel       context.CancelFunc
	data         interface{}
}

// data expected as pointer
func (dep *Watcher) load(ctx context.Context, data interface{}) (err error) {
	secret, err := dep.Load(ctx)
	if err != nil {
		return err
	}

	dep.mutex.Lock()
	defer dep.mutex.Unlock()
	dep.data = secret
	data = dep.data
	return nil
}

// Stop stops the auto-refresh mechanism.
func (dep *Watcher) Stop(ctx context.Context) error {
	dep.cancel()
	dep.refreshTimer.Stop()
	return nil
}

// autoRefresh periodically refreshes the configuration.
func (dep *Watcher) autoRefresh(ctx context.Context, data interface{}) {
	for {
		select {
		case <-dep.refreshTimer.C:
			if err := dep.load(ctx, data); err != nil {
				log.Printf("Error refreshing config: %v", err)
			}
		case <-ctx.Done():
			return
		}
	}
}
