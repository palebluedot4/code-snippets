package deadlock

import "sync"

func RecursiveRWMutex() {
	var mu sync.RWMutex
	mu.RLock()
	defer mu.RUnlock()
	go func() {
		mu.Lock()
		defer mu.Unlock()
	}()
	mu.RLock()
	defer mu.RUnlock()
}
