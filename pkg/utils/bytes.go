package utils

import (
	"github.com/kvkang/common/pkg"
)

// BytesAll all the item is same input value, if not same, return the diff index
func BytesAll(bs []byte, b byte) (int, bool) {
	for i := range bs {
		if bs[i] != b {
			return i, false
		}
	}
	return pkg.ErrIndexReturn, true
}

// BytesAllSet set all the item as input value
func BytesAllSet(bs []byte, b byte) {
	for i := range bs {
		bs[i] = b
	}
}

func sameLenBytesEqual(bs1, bs2 []byte) (int, bool) {
	for i, end := 0, len(bs1)-1; i < end; i++ {
		if bs1[i] != bs2[i] {
			return i, false
		}
	}
	return pkg.ErrIndexReturn, true
}

// BytesEqual compare the two bytes, return the diff index and false, if all same, return -1, true
func BytesEqual(bs1, bs2 []byte) (int, bool) {
	if len(bs1) == len(bs2) {
		return sameLenBytesEqual(bs1, bs2)
	}
	minLen := MinInt(len(bs1), len(bs2))
	if i, same := sameLenBytesEqual(bs1[0:minLen], bs2[0:minLen]); !same {
		return i, false
	}
	return minLen, false
}
