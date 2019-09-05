package openssl

import (
	"crypto/cipher"
)

// ECBEncrypt ()
func ECBEncrypt(block cipher.Block, src []byte, padding string) ([]byte, error) {
	blockSize := block.BlockSize()
	src = Padding(padding, src, blockSize)

	encryptData := make([]byte, len(src))

	ecb := NewECBEncrypter(block)
	ecb.CryptBlocks(encryptData, src)

	return encryptData, nil
}

// ECBDecrypt ()
func ECBDecrypt(block cipher.Block, src []byte, padding string) ([]byte, error) {
	dst := make([]byte, len(src))

	mode := NewECBDecrypter(block)
	mode.CryptBlocks(dst, src)

	dst = UnPadding(padding, dst)

	return dst, nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter 返回在电子代码簿中加密的BlockMode
//模式，使用给定的块。
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}

// BlockSize
func (x *ecbEncrypter) BlockSize() int {
	return x.blockSize
}

// CryptBlocks ()
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter 返回在电子代码簿中解密的块模式
//模式，使用给定的块。
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}

// BlockSize ()
func (x *ecbDecrypter) BlockSize() int {
	return x.blockSize
}

// CryptBlocks ()
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
