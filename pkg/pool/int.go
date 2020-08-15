package pool

import (
	"sync"
)

// NewIntsPool build an IntsPool
func NewIntsPool(size int, getAfter, putBefore func([]int) []int) *IntsPool {
	return &IntsPool{p: &sync.Pool{New: func() interface{} { return make([]int, size) }}}
}

// IntsPool is an ints pool
type IntsPool struct {
	p         *sync.Pool
	getAfter  func([]int) []int
	putBefore func([]int) []int
}

// Get get an []int from pool
func (p *IntsPool) Get() []int {
	bs := p.p.Get().([]int)
	if p.getAfter != nil {
		bs = p.getAfter(bs)
	}
	return bs
}

// Put an []int to an pool
func (p *IntsPool) Put(bs []int) bool {
	if p.putBefore != nil {
		bs = p.putBefore(bs)
	}
	if bs != nil {
		p.p.Put(bs)
		return true
	}
	return false
}
