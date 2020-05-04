package main

import (
	"github.com/yangtizi/go/http/fast"

	"github.com/valyala/fasthttp"
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
