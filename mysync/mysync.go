package mysync

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	done uint32
	mu   sync.Mutex
}

func (once *Once) Do(f func()) {
	if atomic.LoadUint32(&once.done) == 1 {
		return
	}

	once.mu.Lock()
	defer once.mu.Unlock()

	if once.done == 0 {
		defer atomic.StoreUint32(&once.done, 1)
		f()
	}
}
