package dispatch

// IEvent 事件接口
type IEvent interface {
	Target() IEventDispatcher
	Key() interface{}
	Data() interface{}
}

// TEventHandler 监听器函数Handler
// ? 回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false
type TEventHandler func(event IEvent) bool

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
