package sysutils

import (
	"net"
)

// IPIntToStr ip地址int转str
func IPIntToStr(n uint32) string {
	b := NewBytes(4)

	for i := 0; i < 4; i++ {
		b[i] = byte(n & 0xFF)
		n = n >> 8
	}

	return net.IPv4(b[3], b[2], b[1], b[0]).To4().String()
}

// IPStrToInt ip地址str转int
func IPStrToInt(s string) uint32 {
	ip := net.ParseIP(Trim(s))
	if ip == nil {
		return 0
	}

	to4 := ip.To4()
	if to4 == nil {
		return 0
	}
	n := uint32(0)
	n += uint32(to4[0]) << 24
	n += uint32(to4[1]) << 16
	n += uint32(to4[2]) << 8
	n += uint32(to4[3])
	return n
}
