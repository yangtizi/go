package sysutils

import (
	"bytes"
	"math/rand"
)

// NewBytes 新建byte
func NewBytes(n int) []byte {
	b := make([]byte, n)
	return b
}

// NewRandomBytes 新建随机byte
func NewRandomBytes(n int) []byte {
	b := make([]byte, n)

	for i := 0; i < n; i++ {
		b[i] = byte(rand.Intn(256)) // 随机进去
	}

	return b
}

// Concat 连接
func Concat(a []byte, b []byte) []byte {
	buff := bytes.NewBuffer([]byte{})

	buff.Write(a)
	buff.Write(b)
	return buff.Bytes()
}
