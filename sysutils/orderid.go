package sysutils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Order = &TOrderID{}

type TOrderID struct {
	nIndex int64 // 内置的顺序索引
	mutex  sync.Mutex
}

// 生成字符串订单号
func (m *TOrderID) GenString() string {
	str := fmt.Sprintf("%04d%02d%02d%02d%02d%02d%05d",
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		time.Now().Hour(),
		time.Now().Minute(),
		time.Now().Second(),
		time.Now().Nanosecond()%100000,
	)

	m.mutex.Lock()
	nIndex := m.nIndex
	m.nIndex++
	if m.nIndex >= 999 {
		m.nIndex = 0
	}

	str += Int64ToStr(nIndex)
	m.mutex.Unlock()

	return str
}

func (m *TOrderID) GenMerchantID() int {
	m.mutex.Lock()
	nIndex := m.nIndex
	m.nIndex++
	if m.nIndex >= 999 {
		m.nIndex = 0
	}
	m.mutex.Unlock()

	n := rand.Intn(1000000)*1000 + int(nIndex)
	return n
}
