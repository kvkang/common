package pool

import (
	"sync"

	"github.com/kvkang/common/pkg/pool"
)

const _size = 8 // max len for bits to byte

var bytesOnce sync.Once
var bytesPoolInstance *pool.BytesPool

// BytesPoolInstance return an global instance of bytes pool
func BytesPoolInstance() *pool.BytesPool {
	bytesOnce.Do(func() {
		bytesPoolInstance = pool.NewBytesPool(_size, nil, nil)
	})
	return bytesPoolInstance
}

var intsOnce sync.Once
var intsPoolInstance *pool.IntsPool

// IntsPoolInstance return an global instance of ints pool
func IntsPoolInstance() *pool.IntsPool {
	intsOnce.Do(func() {
		intsPoolInstance = pool.NewIntsPool(_size, nil, nil)
	})
	return intsPoolInstance
}
