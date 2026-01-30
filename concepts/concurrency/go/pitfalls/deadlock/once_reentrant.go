package deadlock

import "sync"

func ReentrantDeadlock() {
	var once sync.Once
	once.Do(func() {
		once.Do(func() {
		})
	})
}
