package sysutils

import (
	"math/rand"
	"time"

	uuid "github.com/yangtizi/go/sysutils/uuid"
)

// GenUUID 获取uuid xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
func GenUUID() string {
	guid := uuid.NewV4()
	strGUID := guid.String()
	strGUID = StringReplaceAll(strGUID, "-", "")
	strGUID = UpperCase(strGUID)
	return strGUID
}

// GenGUID 生成GUID xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func GenGUID() string {
	guid := uuid.NewV4()
	strGUID := guid.String()
	strGUID = UpperCase(strGUID)
	return strGUID
}

// GenRandomNumber 生成随机的字符串
func GenRandomNumber(nLen int) string {
	bytes := []byte("0123456789")
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < nLen; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GenRandomHex 生成随机字符串
func GenRandomHex(nLen int) string {
	s := GenUUID()
	for {
		if nLen < len(s) {
			break
		}
		s += GenUUID()
	}
	return s[0:nLen]
}
