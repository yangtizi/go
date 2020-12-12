package dispatch

// IListener 监听器接口
type IListener interface {
	RemoveSelf() // RemoveSelf 移除掉自己
}

// TListener 监听器
type TListener struct {
	parent  *TEventDispatcher // 注册到的派发器
	key     interface{}       // 事件注册的名称
	pSender interface{}       // 事件注册人（移除时可以用）
	handler TEventHandler     // 事件的handler（回调函数）
	Next    *TListener        // 链表
	Prev    *TListener        // 链表
}

// RemoveSelf 移除掉自己
func (m *TListener) RemoveSelf() {
	// 拿到上下的
	pNext := m.Next
	pPrev := m.Prev
	defer func() {
		m.Next = nil
		m.Prev = nil
	}()
	if pNext == nil && pPrev == nil {
		// log.Println("两边都空。这里估计是唯一的链")
		// 两边都空。两边都空。这里估计是唯一的链, 删掉map的记录
		v, ok := m.parent.events.Load(m.key)
		if ok && v == m {
			// log.Println("清空")
			m.parent.events.Delete(m.key)
		}
		return
	}

	if pNext == nil {
		// log.Println("链尾, 上一链的链尾清掉")
		// 链尾, 上一链的链尾清掉
		pPrev.Next = nil
		return
	} else if pPrev == nil {

		// log.Println("链头， 下一链的链头清掉")
		// 链头， 下一链的链头清掉
		pNext.Prev = nil
		// 既然是链头， 移除以后需要重新指向新的链头
		m.parent.events.Store(m.key, pNext)
		return
	}

	// 前后链握手
	pNext.Prev, pPrev.Next = pPrev, pNext
}
