package main

import (
	"fmt"
	"github.com/yangtizi/go/crypto/gzip"
)


func main() {
	a := []byte("xiaonini")

	b := gzip.Compress(a)
	fmt.Println(b)

	c := gzip.UnCompress(b)
	fmt.Println(c)
	fmt.Println(string(c))
}