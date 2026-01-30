package deadlock

import "sync"

func BadUnlockOnEarlyReturn(cond bool) {
	var mu sync.Mutex
	mu.Lock()
	if cond {
		return
	}
	mu.Unlock()
}

func GoodUnlockWithDefer(cond bool) {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	if cond {
		return
	}
}

func BadUnlockInLoop() {
	var mu sync.Mutex
	for i := range 10 {
		mu.Lock()
		if i == 5 {
			continue
		}
		mu.Unlock()
	}
}

func GoodUnlockInLoop() {
	var mu sync.Mutex
	for i := range 10 {
		func() {
			mu.Lock()
			defer mu.Unlock()
			if i == 5 {
				return
			}
		}()
	}
}
