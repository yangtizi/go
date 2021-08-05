package main

import (
	"fmt"

	"github.com/yangtizi/go/crypto/aes"
	"github.com/yangtizi/go/crypto/gzip"
	"github.com/yangtizi/go/crypto/zlib"
)

func main() {
	aesDemo1()
	aesDemo2()
	zlibDemo()
	gzipDemo()
}

func aesDemo1() {
	fmt.Println("aseDemo1 -----------------")
	strWary := "warrially"          // 加密原文
	strAESKey := "1234567812345678" // AES KEY

	bufCrypto, _ := aes.CoAES([]byte(strWary), []byte(strAESKey), make([]byte, 16, 16), "PKCS5")
	fmt.Println("加密后是: ", bufCrypto)

	strWary2, _ := aes.UnAES(bufCrypto, []byte(strAESKey), make([]byte, 16, 16), "PKCS5")
	fmt.Println("解密", string(strWary2)) //
	fmt.Println("aseDemo1 -----------------")
}

func aesDemo2() {
	fmt.Println("aseDemo2 -----------------")
	strWary := "warrially"          // 加密原文
	strAESKey := "1234567812345678" // AES KEY

	bufCrypto, _ := aes.CoAES([]byte(strWary), []byte(strAESKey), []byte(strAESKey), "PKCS7")
	fmt.Println("加密后是: ", bufCrypto)

	strWary2, _ := aes.UnAES(bufCrypto, []byte(strAESKey), []byte(strAESKey), "PKCS7")
	fmt.Println("解密", string(strWary2)) //
	fmt.Println("aseDemo2 -----------------")
}

func zlibDemo() {
	fmt.Println("zlibDemo -----------------")
	strWary := "warrially"

	bufCrypto := zlib.Compress([]byte(strWary))
	fmt.Println("加密后是: ", bufCrypto)

	strWary2 := zlib.UnCompress(bufCrypto)
	fmt.Println("解密", string(strWary2)) //
	fmt.Println("zlibDemo -----------------")
}

func gzipDemo() {
	a := []byte("xiaonini")

	b := gzip.Compress(a)
	fmt.Println(b)

	c := gzip.UnCompress(b)
	fmt.Println(c)
	fmt.Println(string(c))
}
