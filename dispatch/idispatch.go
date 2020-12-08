package dispatch

var instance = NewEventDispatcher()

// Ins 获取单例
func Ins() IEventDispatcher {
	return instance
}

// IEventDispatcher 事件调度接口
type IEventDispatcher interface {
	// 注册事件
	Register(key interface{}, pSender interface{}, handler TEventHandler)
	// 派发事件
	Dispatch(key interface{}, data interface{})
	// todo 移除事件
}
