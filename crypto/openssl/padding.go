package openssl

import (
	"bytes"
)

// PKCS5PADDING = PKCS
const PKCS5PADDING = "PKCS5"

// PKCS7PADDING = PKCS
const PKCS7PADDING = "PKCS7"

// ZEROSPADDING = ZERO
const ZEROSPADDING = "ZEROS"

// Padding ()
func Padding(padding string, src []byte, blockSize int) []byte {
	switch padding {
	case PKCS5PADDING:
		src = PKCS5Padding(src, blockSize)
	case PKCS7PADDING:
		src = PKCS7Padding(src, blockSize)
	case ZEROSPADDING:
		src = ZerosPadding(src, blockSize)
	}
	return src
}

// UnPadding ()
func UnPadding(padding string, src []byte) []byte {
	switch padding {
	case PKCS5PADDING:
		src = PKCS5Unpadding(src)
	case PKCS7PADDING:
		src = PKCS7UnPadding(src)
	case ZEROSPADDING:
		src = ZerosUnPadding(src)
	}
	return src
}

// PKCS5Padding ()
func PKCS5Padding(src []byte, blockSize int) []byte {
	return PKCS7Padding(src, blockSize)
}

// PKCS5Unpadding ()
func PKCS5Unpadding(src []byte) []byte {
	return PKCS7UnPadding(src)
}

// PKCS7Padding ()
func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7UnPadding ()
func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

// ZerosPadding ()
func ZerosPadding(src []byte, blockSize int) []byte {
	paddingCount := blockSize - len(src)%blockSize
	if paddingCount == 0 {
		return src
	}

	return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
}

// ZerosUnPadding ()
func ZerosUnPadding(src []byte) []byte {
	for i := len(src) - 1; ; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
}
