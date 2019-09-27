package main

import (
	"github.com/valyala/fasthttp"
	"github.com/yangtizi/go/http/fast"
)

func main() {
	fasthttpStartDemo()
}

// 启动fast服务器
func fasthttpStartDemo() {
	fast.Register("/fast/Demo", onFastDemo)
	fast.StartServer(":8080")
}

func onFastDemo(ctx *fasthttp.RequestCtx) {
}
