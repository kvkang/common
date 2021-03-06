package utils

import (
	"github.com/kvkang/common/pkg"
)

// IntsAll all the item is same input value, if not same, return the diff index
func IntsAll(bs []int, b int) (int, bool) {
	for i := range bs {
		if bs[i] != b {
			return i, false
		}
	}
	return pkg.ErrIndexReturn, true
}

// IntsAllSet set all the item as input value
func IntsAllSet(bs []int, b int) {
	for i := range bs {
		bs[i] = b
	}
}

func sameLenIntsEqual(bs1, bs2 []int) (int, bool) {
	for i, end := 0, len(bs1)-1; i < end; i++ {
		if bs1[i] != bs2[i] {
			return i, false
		}
	}
	return pkg.ErrIndexReturn, true
}

// IntsEqual compare the two ints, return the diff index and false, if all same, return -1, true
func IntsEqual(bs1, bs2 []int) (int, bool) {
	if len(bs1) == len(bs2) {
		return sameLenIntsEqual(bs1, bs2)
	}
	minLen := MinInt(len(bs1), len(bs2))
	if i, same := sameLenIntsEqual(bs1[0:minLen], bs2[0:minLen]); !same {
		return i, false
	}
	return minLen, false
}
