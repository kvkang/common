package utils

import (
	"github.com/kvkang/common/pkg"
	"github.com/kvkang/common/pkg/internal/pool"
)

func bit8ToByte(bits []int) byte {
	return byte(bits[0]<<7 | bits[1]<<6 | bits[2]<<5 | bits[3]<<4 | bits[4]<<3 | bits[5]<<2 | bits[6]<<1 | bits[7])
}

// BitsToByte return return byte and nil if len(bits) <= 8, return byte(from bits[0:8]) and ErrOutOfRange if len(bits) > 8
func BitsToByte(bits []int) (byte, error) {
	if len(bits) == 8 {
		return bit8ToByte(bits), nil
	} else if len(bits) > 8 {
		return bit8ToByte(bits[0:8]), pkg.ErrSliceOutOfRange
	}

	// len(bits) < 8
	bs := pool.IntsPoolInstance().Get()
	defer func() {
		pool.IntsPoolInstance().Put(bs)
	}()

	copy(bs, bits)
	IntsAllSet(bs[len(bits):8], 0)
	return bit8ToByte(bs[0:8]), nil
}
