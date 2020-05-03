package fast

import (
	"sync"

	"github.com/valyala/fasthttp"
)

// ! import "yangtizi/http/fast"

var mpHandler sync.Map

// Register 注册回调
func Register(strPath string, cb func(*fasthttp.RequestCtx)) {
	mpHandler.Store(strPath, cb)
}
