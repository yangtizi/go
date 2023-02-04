package sysutils

import (
	"encoding/binary"
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

/*
ipInt := 3232235777
ip := IntToIP(ipInt)
fmt.Println(ip.String())
*/
func IntToIP(ipInt int) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, uint32(ipInt))
	return ip
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

/*
	ip := net.ParseIP("192.168.1.1")
	ipInt := IPToInt(ip)
	fmt.Println(ipInt)
*/

func IPToInt(ip net.IP) int {
	if len(ip) == 16 {
		return int(binary.BigEndian.Uint32(ip[12:16]))
	}
	return int(binary.BigEndian.Uint32(ip))
}
