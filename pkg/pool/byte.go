package pool

import (
	"sync"
)

// NewBytesPool build an BytesPool
func NewBytesPool(size int, getAfter, putBefore func([]byte) []byte) *BytesPool {
	return &BytesPool{p: &sync.Pool{New: func() interface{} { return make([]byte, size) }}}
}

// BytesPool is an bytes pool
type BytesPool struct {
	p         *sync.Pool
	getAfter  func([]byte) []byte
	putBefore func([]byte) []byte
}

// Get get an []byte from pool
func (p *BytesPool) Get() []byte {
	bs := p.p.Get().([]byte)
	if p.getAfter != nil {
		bs = p.getAfter(bs)
	}
	return bs
}

// Put an []byte to an pool
func (p *BytesPool) Put(bs []byte) bool {
	if p.putBefore != nil {
		bs = p.putBefore(bs)
	}
	if bs != nil {
		p.p.Put(bs)
		return true
	}
	return false
}
