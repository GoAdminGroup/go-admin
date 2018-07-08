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
)

type EndPointFun func(ctx *fasthttp.RequestCtx, path string, prefix string, user auth.User)

// 显示列表
func ShowInfo(ctx *fasthttp.RequestCtx, path string, prefix string, user auth.User) {

	defer handle(ctx)

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

func DeleteData(ctx *fasthttp.RequestCtx, path string, prefix string, user auth.User) {

	defer handle(ctx)

	id := string(ctx.FormValue("id")[:])

	transform.DeleteDataFromDatabase(prefix, id)

	// TODO: 增加反馈

	ctx.WriteString(`{"code":200, "msg":"删除成功"`)
	return
}


// 全局错误处理
func handle(ctx *fasthttp.RequestCtx) {
	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println(string(debug.Stack()[:]))
		ctx.Error(`{"code":500, "msg":"系统错误"}`, fasthttp.StatusInternalServerError)
		return
	}
}
