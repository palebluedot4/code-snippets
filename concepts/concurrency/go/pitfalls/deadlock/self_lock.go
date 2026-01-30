package deadlock

import "sync"

func SelfLock() {
	var mu sync.Mutex
	mu.Lock()
	mu.Lock()
}

func SelfLockRWMutex() {
	var mu sync.RWMutex
	mu.Lock()
	mu.Lock()
}

func SelfLockRWMutexMixed() {
	var mu sync.RWMutex
	mu.RLock()
	mu.Lock()
}
