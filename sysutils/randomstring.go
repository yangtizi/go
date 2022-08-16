package sysutils

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ23456789!@#$%^&*()_"

func init() {
	rand.Seed(time.Now().UnixNano() + 12345)
}

// RandStr(10)
func RandStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

// 随机 并且不报错
func Random(n int) int {
	if n <= 0 {
		return 0
	}

	return rand.Intn(n)
}
