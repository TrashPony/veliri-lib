package timecache

import (
	"sync"
	"time"
)

var (
	timer *TimeCache
	once  sync.Once
)

type TimeCache struct {
	currentUnix     int64
	currentUnixNano int64
	interval        time.Duration
	stopChan        chan struct{}
	mu              sync.RWMutex
}

func GetTimer() *TimeCache {
	once.Do(func() {
		timer = &TimeCache{
			interval: 32 * time.Millisecond, // Обновление каждые 32 мс (~30 FPS)
			stopChan: make(chan struct{}),
		}
		go timer.run()
	})
	return timer
}

func (tc *TimeCache) run() {
	ticker := time.NewTicker(tc.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			tc.mu.Lock()
			tc.currentUnix = now.Unix()
			tc.currentUnixNano = now.UnixNano()
			tc.mu.Unlock()
		case <-tc.stopChan:
			return
		}
	}
}

func (tc *TimeCache) Unix() int64 {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	return tc.currentUnix
}

func (tc *TimeCache) UnixNano() int64 {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	return tc.currentUnixNano
}

func (tc *TimeCache) Stop() {
	close(tc.stopChan)
}
