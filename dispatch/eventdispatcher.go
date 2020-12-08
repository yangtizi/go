package dispatch

import (
	"fmt"
	"sync"
)

// NewEventDispatcher 创建事件派发器
func NewEventDispatcher() *TEventDispatcher {
	p := &TEventDispatcher{}
	return p
}

// TEventDispatcher  事件派发器
type TEventDispatcher struct {
	events sync.Map // 事件表
	// senders sync.Map // 发送表
	// regs    sync.Map // 实际表
	nIndex int64
}

// Register 注册
func (m *TEventDispatcher) Register(key interface{}, pSender interface{}, handler TEventHandler) {
	// 创建事件表
	v1, _ := m.events.LoadOrStore(key, []*TListener{})
	// v2, _ := m.senders.LoadOrStore(key, []*TListener{})

	m.nIndex++

	pListener := &TListener{
		nIndex:  m.nIndex,
		key:     key,
		pSender: pSender,
		handler: handler,
	}

	// 这个新的注册
	events := append(v1.([]*TListener), pListener)
	m.events.Store(key, events)

	// senders := append(v2.([]*TListener), pListener)
	// m.senders.Store(key, senders)

	// m.regs.Store(m.nIndex, pListener)

}

// Dispatch 事件派发
func (m *TEventDispatcher) Dispatch(key interface{}, data interface{}) {
	v, ok := m.events.Load(key)

	if !ok {
		fmt.Println("未注册的事件")
		return
	}

	events := v.([]*TListener)

	// 进行循环发送
	for _, pListener := range events {
		if pListener != nil && pListener.handler != nil {
			pListener.handler(&TEvent{
				target: m,
				key:    key,
				data:   data,
			})
		}
	}

}
