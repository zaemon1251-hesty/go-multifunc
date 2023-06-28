package mysync

import (
	"sync/atomic"
)

type Multi struct {
	count atomic.Int64
	n     int64
	wait  chan struct{}
	done  chan struct{}
}

func NewMulti(n int64) *Multi {
	return &Multi{
		wait: make(chan struct{}, int(n)),
		done: make(chan struct{}),
		n:    n,
	}
}

func (m *Multi) Do(f func()) {
	if m.count.Load() >= m.n {
		return
	}

	select {
	case m.wait <- struct{}{}:
	default:
		<-m.done
	}

	if m.count.Load() < m.n {
		defer func() {
			if m.count.Add(1) == m.n {
				close(m.done)
			}
		}()
		f()
	}
}
