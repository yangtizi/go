package main

import "github.com/yangtizi/go/http/fast"

func main() {
	fasthttpStartDemo()
}

// 启动fast服务器
func fasthttpStartDemo() {
	fast.StartServer(":8080")
}
