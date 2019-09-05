package fast

import (
	"github.com/valyala/fasthttp"
	log "github.com/yangtizi/go/log/zaplog"
)

// ! import "github.com/yangtizi/go/http/fast"

// StartServer (地址)
func StartServer(strAddress string) {
	if len(strAddress) <= 0 { //
		log.Println("错误的Address", strAddress) //
	} //
	log.Println("服务器", strAddress, "正在开启")                      //
	err := fasthttp.ListenAndServe(strAddress, fastHTTPHandler) //
	if err != nil {                                             //
		log.Println("fasthttp", err) //
	} //
	log.Println("服务器", strAddress, "已经关闭") //
}
