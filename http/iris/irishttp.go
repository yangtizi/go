package iris

import (
	"log"

	iirriiss "github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var app = iirriiss.New()

// StartServer (地址)
func StartServer(strAddress string) {
	if len(strAddress) <= 0 {
		log.Println("错误的Address", strAddress) //
	}
	log.Println("服务器", strAddress, "正在开启") //
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	app.Run(iirriiss.Addr(strAddress))
	log.Println("服务器", strAddress, "已经关闭") //
}

// Party 拿到
func Party(relativePath string, handlers ...context.Handler) router.Party {
	return app.Party(relativePath, handlers...)
}
