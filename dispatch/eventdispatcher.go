package dispatch

import (
	"errors"
	"sync"
)

var instance = NewEventDispatcher()

// Ins 获取单例
func Ins() IEventDispatcher {
	if instance == nil {
		instance = NewEventDispatcher()
	}
	return instance
}

// IEventDispatcher 事件调度接口
type IEventDispatcher interface {
	// 注册事件
	Register(key interface{}, pSender interface{}, handler TEventHandler) IListener
	// 派发事件
	Dispatch(key interface{}, data interface{}) (int, error)
	// 移除事件
	Remove(pSender interface{})
}

// NewEventDispatcher 创建事件派发器
func NewEventDispatcher() *TEventDispatcher {
	p := &TEventDispatcher{}
	return p
}

// TEventDispatcher  事件派发器
type TEventDispatcher struct {
	events  sync.Map // 事件表
	senders sync.Map // 发送表
}

// Register 注册
func (m *TEventDispatcher) Register(key interface{}, pSender interface{}, handler TEventHandler) IListener {
	//
	pListener := &TListener{
		parent:  m,
		key:     key,
		pSender: pSender,
		handler: handler,
	}

	// 创建事件表
	if v, ok := m.events.Load(key); ok {
		v.(*TListener).Prev = pListener
		pListener.Next = v.(*TListener)
	}

	m.events.Store(key, pListener)

	// todo 根据sender 来创建表， 主要为了移除用
	v1, _ := m.senders.LoadOrStore(pSender, []*TListener{})
	senders := append(v1.([]*TListener), pListener)
	m.senders.Store(pSender, senders)

	return pListener
}

// Dispatch 事件派发
func (m *TEventDispatcher) Dispatch(key interface{}, data interface{}) (int, error) {
	v, ok := m.events.Load(key)

	if !ok {
		return 0, errors.New("未注册的事件")
	}

	// 进行循环发送
	pListener := v.(*TListener)
	nCount := 0

	for {
		if pListener == nil {
			break
		}

		bContinue := pListener.handler(&TEvent{
			target: m,
			key:    key,
			data:   data,
		})

		nCount++

		if !bContinue {
			break
		}

		pListener = pListener.Next
	}

	return nCount, nil
}

// Remove 移除
func (m *TEventDispatcher) Remove(pSender interface{}) {
	v1, ok := m.senders.LoadAndDelete(pSender)
	if ok {
		for _, v := range v1.([]*TListener) {
			v.RemoveSelf()
		}
	}
}
