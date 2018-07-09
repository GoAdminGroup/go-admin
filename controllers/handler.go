package controller

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
	"goAdmin/menu"
	"goAdmin/template"
	"goAdmin/transform"
	"runtime/debug"
	"github.com/mgutz/ansi"
	"strconv"
	"log"
	"github.com/go-sql-driver/mysql"
)

// 显示列表
func ShowInfo(ctx *fasthttp.RequestCtx) {

	defer handle(ctx)

	user := ctx.UserValue("cur_user").(auth.User)
	prefix := ctx.UserValue("prefix").(string)
	path := string(ctx.Path())

	page := string(ctx.QueryArgs().Peek("page")[:])
	if page == "" {
		page = "1"
	}
	pageSize := string(ctx.QueryArgs().Peek("pageSize")[:])
	if pageSize == "" {
		pageSize = "10"
	}

	thead, infoList, paginator, title, description := transform.TransfromData(page, pageSize, path, prefix)

	menu.GlobalMenu.SetActiveClass(path)

	buffer := new(bytes.Buffer)

	if string(ctx.Request.Header.Peek("X-PJAX")[:]) == "true" {
		template.InfoListPjax(infoList, (*menu.GlobalMenu).GlobalMenuList, thead, paginator, title, description, buffer)
	} else {
		template.InfoList(infoList, (*menu.GlobalMenu).GlobalMenuList, thead, paginator, title, description, user, buffer)
	}

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

func DeleteData(ctx *fasthttp.RequestCtx) {

	defer handle(ctx)

	prefix := ctx.UserValue("prefix").(string)

	id := string(ctx.FormValue("id")[:])

	transform.DeleteDataFromDatabase(prefix, id)

	// TODO: 增加反馈

	ctx.WriteString(`{"code":200, "msg":"删除成功"`)
	return
}


// 全局错误处理
func handle(ctx *fasthttp.RequestCtx) {

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
