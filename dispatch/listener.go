package dispatch

// TListener 监听器
type TListener struct {
	nIndex  int64
	key     interface{}
	pSender interface{}
	handler TEventHandler
}
