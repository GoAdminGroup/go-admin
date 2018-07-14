package controller

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"runtime/debug"
	"github.com/mgutz/ansi"
	"strconv"
	"log"
	"github.com/go-sql-driver/mysql"
)

// 全局错误处理
func GlobalDeferHandler(ctx *fasthttp.RequestCtx) {

	log.Println("[GoAdmin]",
		ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode())+" ", "white:blue"),
		ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
		string(ctx.Path()))

	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println(string(debug.Stack()[:]))

		var (
			errMsg string
			mysqlError *mysql.MySQLError
			ok bool
		)
		if errMsg, ok = err.(string); ok {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetContentType("application/json")
			ctx.WriteString(`{"code":500, "msg":"`+ errMsg + `"}`)
			return
		} else if mysqlError, ok = err.(*mysql.MySQLError); ok {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetContentType("application/json")
			ctx.WriteString(`{"code":500, "msg":"`+ mysqlError.Error() + `"}`)
			return
		} else {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			ctx.SetContentType("application/json")
			ctx.WriteString(`{"code":500, "msg":"系统错误"}`)
			return
		}
	}
}
