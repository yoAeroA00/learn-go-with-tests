package sync

import (
	"sync"
	"sync/atomic"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Inc() {
	(*c).mu.Lock()
	defer (*c).mu.Unlock()
	(*c).count++
}

func (c *Counter) Value() int {
	return (*c).count
}

type ACounter struct {
	value atomic.Int64
}

func (c *ACounter) Inc() {
	(*c).value.Add(1)
}

func (c *ACounter) Value() int64 {
	return (*c).value.Load()
}
