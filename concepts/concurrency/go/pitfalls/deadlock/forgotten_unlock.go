package deadlock

import "sync"

func ForgottenUnlock() {
	var mu sync.Mutex
	for range 10 {
		go func() {
			mu.Lock()
		}()
	}
}
