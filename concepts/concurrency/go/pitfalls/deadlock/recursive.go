package deadlock

import "sync"

type RecursiveResource struct {
	mu sync.Mutex
}

func (r *RecursiveResource) EntryPoint() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.internalMethod()
}

func (r *RecursiveResource) internalMethod() {
	r.mu.Lock()
	defer r.mu.Unlock()
}
