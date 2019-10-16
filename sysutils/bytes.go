package sysutils

import (
	"bytes"
	"errors"
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

// BlockCopy 拷贝
func BlockCopy(src []byte, srcOffset int, dst []byte, dstOffset, count int) (bool, error) {
	srcLen := len(src)
	if srcOffset > srcLen || count > srcLen || srcOffset+count > srcLen {
		return false, errors.New("源缓冲区 索引超出范围")
	}
	dstLen := len(dst)
	if dstOffset > dstLen || count > dstLen || dstOffset+count > dstLen {
		return false, errors.New("目标缓冲区 索引超出范围")
	}
	index := 0
	for i := srcOffset; i < srcOffset+count; i++ {
		dst[dstOffset+index] = src[srcOffset+index]
		index++
	}
	return true, nil
}
