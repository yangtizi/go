package dispatch

// IEvent 事件接口
type IEvent interface {
	Target() IEventDispatcher
	Key() interface{}
	Data() interface{}
}

// TEventHandler 监听器函数Handler
type TEventHandler func(event IEvent)
