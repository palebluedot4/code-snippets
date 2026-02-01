package copylock

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func BadIncrement(c Counter) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func GoodIncrement(c *Counter) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
