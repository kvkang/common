package netutils

import (
	"errors"
	"math"
	"net"

	"github.com/kvkang/common/pkg"
	"github.com/kvkang/common/pkg/internal/pool"
	"github.com/kvkang/common/pkg/utils"
)

var v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}

// ErrIPv6NotSupport same function not support ipv6 yet.
var ErrIPv6NotSupport = errors.New("not support ipv6")

// IPv4 parse value i to net.IP
func IPv4(i int) (net.IP, error) {
	if i < 0 || i > math.MaxUint32 {
		return nil, pkg.ErrSliceOutOfRange
	}
	return net.IPv4(byte(i>>24&0xFF), byte(i>>16&0xFF), byte(i>>8&0xFF), byte(i&0xFF)), nil
}

// Int parse an ip to int
func Int(ip net.IP) (int, error) {
	if len(ip) == net.IPv4len {
		return int(ip[0])<<24 | int(ip[1])<<16 | int(ip[2])<<8 | int(ip[3]), nil
	} else if len(ip) == net.IPv6len {
		if _, ok := utils.BytesEqual(ip[0:12], v4InV6Prefix); ok {
			return int(ip[12])<<24 | int(ip[13])<<16 | int(ip[14])<<8 | int(ip[15]), nil
		}
		return pkg.ErrIntReturn, ErrIPv6NotSupport
	}
	return pkg.ErrIntReturn, pkg.ErrSliceLength
}

// BitsToIPv4 return ipv4 by bits
func BitsToIPv4(bits []int) (net.IP, error) {
	if len(bits) > 32 {
		return nil, pkg.ErrSliceOutOfRange
	}

	// bs := []byte{0, 0, 0, 0}
	bs := pool.BytesPoolInstance().Get()
	defer func() { pool.BytesPoolInstance().Put(bs) }()
	utils.BytesAllSet(bs[0:4], 0)

	for i := range bits {
		bs[i], _ = utils.BitsToByte(bits)
		if len(bits) <= 8 {
			break
		}
		bits = bits[8:]
	}
	return net.IPv4(bs[0], bs[1], bs[2], bs[3]), nil
}

// IPv4Add return the ip difference between input ip with add, if the dst value out of range, mark err ErrLoopRange
func IPv4Add(ip net.IP, add int) (got net.IP, err error) {
	var i int
	if i, err = Int(ip); err == nil {
		i = i + add
		if i < 0 {
			err = pkg.ErrLoopByRange
			i = -i
		}
		if i > math.MaxUint32 {
			err = pkg.ErrLoopByRange
		}
		got, _ = IPv4(i % (math.MaxUint32 + 1))
	}
	return
}

// IPv4MaskEnd return the end of ip range, which the input ip in the range with input mask
func IPv4MaskEnd(ip net.IP, mask net.IPMask) net.IP {
	if ipRangeStart := ip.Mask(mask); ipRangeStart != nil {
		off, _ := mask.Size()
		ip, _ := IPv4Add(ipRangeStart, math.MaxUint32>>off)
		return ip
	}
	return nil
}
