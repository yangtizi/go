package fast

import (
	"github.com/valyala/fasthttp"
)

// ! import "yangtizi/http/fast"

func init() {
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	strPath := string(ctx.Path())

	if v, ok := mpHandler.Load(strPath); ok {
		v.(func(*fasthttp.RequestCtx))(ctx)
		return
	}

	// log.Println("ctx.Path = ", strPath)
	// log.Println("ctx.PostBody()", ctx.PostBody())
	// log.Println("ctx.Path()", string(ctx.Path()))
	// log.Println("ctx.Method()", string(ctx.Method()))
	// log.Println("ctx.Request.Header", ctx.Request.Header)
	// log.Println("ctx.Host()", string(ctx.Host()))
	// log.Println("ctx.PostArgs()", ctx.PostArgs())

	// log.Println("ctx.RemoteAddr()", ctx.RemoteAddr())
	// log.Println("ctx.RequestURI()", string(ctx.RequestURI()))
	// log.Println("ctx.IsTLS()", ctx.IsTLS())
	// log.Println("ctx.Referer()", string(ctx.Referer()))
	// log.Println("ctx.UserAgent()", string(ctx.UserAgent()))

	// log.Println(ctx.QueryArgs())

	// args := ctx.QueryArgs()

	// agent := args.Peek("agent")
	// log.Println(string(agent))
	// agent1 := args.Peek("agent1")
	// log.Println(string(agent1))

	// // -----------------------------------------------------------------
	// log.Println("---------------------------------------------------")
	// ctx.Write([]byte(`{"err":"test success"}`))
}
