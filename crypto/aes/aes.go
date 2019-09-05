package aes

import (
	"github.com/yangtizi/go/crypto/openssl"
)

// CoAES AES加密 aes.CoAES([]byte("hello world"), []byte("1234567890123456"), []byte("0000000000000000"), "PKCS7")
func CoAES(rawData []byte, key []byte, iv []byte, strPadding string) ([]byte, error) {
	dst, err := openssl.AesCBCEncrypt(rawData, key, iv, strPadding)
	return dst, err
}

// UnAES AES解密
func UnAES(rawData []byte, key []byte, iv []byte, strPadding string) ([]byte, error) {
	dst, err := openssl.AesCBCDecrypt(rawData, key, iv, strPadding)
	return dst, err
}
