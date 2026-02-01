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

func (c Counter) BadMethod() {
	c.mu.Lock()
	defer c.mu.Unlock()
}

func (c *Counter) GoodMethod() {
	c.mu.Lock()
	defer c.mu.Unlock()
}

func BadPassSyncByValue(wg sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	wg.Wait()
}

func GoodPassSyncByPointer(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	wg.Wait()
}
