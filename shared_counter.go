package sharedcounter

import (
	"sync/atomic"
)

type SharedCounter struct {
	counter int64
}

func NewSharedCounter() *SharedCounter {
	return &SharedCounter{counter: 0}
}

func (c *SharedCounter) Add(num int64) {
	atomic.AddInt64(&c.counter, num)
}

func (c *SharedCounter) Read() int64 {
	return atomic.LoadInt64(&c.counter)
}

func (c *SharedCounter) Reset() {
	atomic.SwapInt64(&c.counter, 0)
}
