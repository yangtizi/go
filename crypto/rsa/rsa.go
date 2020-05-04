package rsa

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/yangtizi/go/crypto/zlib"
)

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}

// WxRSA RSA加密
func WxRSA(origData []byte, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey) //将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData) //RSA算法加密
}

// WxCoRSA 压缩加密 (对于比较大的必须要这样, 不然会报错)
func WxCoRSA(origData []byte, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey) //将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)

	buffer := bytes.NewBufferString("")

	partLen := pub.N.BitLen()/8 - 11
	chunks := split(zlib.Compress(origData), partLen)

	for _, chunk := range chunks {
		by, err := rsa.EncryptPKCS1v15(rand.Reader, pub, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(by)
	}

	return buffer.Bytes(), nil //RSA算法加密
}
