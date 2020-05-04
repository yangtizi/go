package fast

import (
	log "github.com/yangtizi/go/log/zaplog"

	"github.com/valyala/fasthttp"
)

// ! import "github.com/yangtizi/go/http/fast"

// StartServer (地址)
func StartServer(strAddress string) {
	defer log.Infof("服务器", strAddress, "已经关闭") //

	if len(strAddress) <= 0 { //
		log.Errorf("错误的Address = [%s]", strAddress) //
		return
	} //
	log.Infof("服务器 [%s] 正在开启", strAddress)                      //
	err := fasthttp.ListenAndServe(strAddress, fastHTTPHandler) //
	if err != nil {                                             //
		log.Errorf("fasthttp.ListenAndServe错误 = [%v]", err) //
	} //

}
