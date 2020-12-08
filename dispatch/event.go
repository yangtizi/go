package dispatch

// TEvent 事件类型基类
type TEvent struct {
	target IEventDispatcher
	key    interface{}
	data   interface{}
}

// Target .
func (m *TEvent) Target() IEventDispatcher {
	return m.target
}

// Key .
func (m *TEvent) Key() interface{} {
	return m.key
}

// Data .
func (m *TEvent) Data() interface{} {
	return m.data
}
