package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
	"goAdmin/connections/mysql"
	"goAdmin/menu"
	"goAdmin/template"
)

// 显示菜单
func ShowMenu(ctx *fasthttp.RequestCtx) {
	defer GlobalDeferHandler(ctx)

	path := string(ctx.Path())
	user := ctx.UserValue("cur_user").(auth.User)

	menu.GlobalMenu.SetActiveClass(path)

	buffer := new(bytes.Buffer)

	//if string(ctx.Request.Header.Peek("X-PJAX")[:]) == "true" {
	//template.MenuPanel((*menu.GlobalMenu).GetEditMenuList(), (*menu.GlobalMenu).GlobalMenuList, user, buffer)
	//} else {
	template.MenuPanel((*menu.GlobalMenu).GetEditMenuList(), (*menu.GlobalMenu).GlobalMenuList, (*menu.GlobalMenu).GlobalMenuOption, user, buffer)
	//}

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

// 显示编辑菜单
func ShowEditMenu(ctx *fasthttp.RequestCtx) {
	id := string(ctx.QueryArgs().Peek("id")[:])
	user := ctx.UserValue("cur_user").(auth.User)

	buffer := new(bytes.Buffer)

	if string(ctx.Request.Header.Peek("X-PJAX")[:]) == "true" {
		template.MenuEditPanelPjax(menu.GetMenuItemById(id), (*menu.GlobalMenu).GlobalMenuOption, buffer)
	} else {
		template.MenuPanel((*menu.GlobalMenu).GetEditMenuList(), (*menu.GlobalMenu).GlobalMenuList, (*menu.GlobalMenu).GlobalMenuOption, user, buffer)
	}

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

// 删除菜单
func DeleteMenu(ctx *fasthttp.RequestCtx) {
	id := string(ctx.QueryArgs().Peek("id")[:])

	buffer := new(bytes.Buffer)

	mysql.Exec("delete from goadmin_menu where id = ?", id)

	menu.SetGlobalMenu()
	template.MenuPanelPjax((*menu.GlobalMenu).GetEditMenuList(), (*menu.GlobalMenu).GlobalMenuOption, buffer)

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

// 编辑菜单
func EditMenu(ctx *fasthttp.RequestCtx) {
	defer GlobalDeferHandler(ctx)

	buffer := new(bytes.Buffer)

	id := string(ctx.FormValue("id")[:])
	title := string(ctx.FormValue("title")[:])
	parent_id := string(ctx.FormValue("parent_id")[:])
	if parent_id == "" {
		parent_id = "0"
	}
	icon := string(ctx.FormValue("icon")[:])
	uri := string(ctx.FormValue("uri")[:])

	mysql.Exec("update goadmin_menu set title = ?, parent_id = ?, icon = ?, uri = ? where id = ?",
		title, parent_id, icon, uri, id)

	menu.SetGlobalMenu()

	template.MenuPanelPjax((*menu.GlobalMenu).GetEditMenuList(), (*menu.GlobalMenu).GlobalMenuOption, buffer)

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

// 新建菜单
func NewMenu(ctx *fasthttp.RequestCtx) {
	defer GlobalDeferHandler(ctx)

	buffer := new(bytes.Buffer)

	title := string(ctx.FormValue("title")[:])
	parent_id := string(ctx.FormValue("parent_id")[:])
	if parent_id == "" {
		parent_id = "0"
	}
	icon := string(ctx.FormValue("icon")[:])
	uri := string(ctx.FormValue("uri")[:])

	mysql.Exec("insert into goadmin_menu (title, parent_id, icon, uri, `order`) values (?, ?, ?, ?, ?)", title, parent_id, icon, uri, (*menu.GlobalMenu).MaxOrder+1)

	(*menu.GlobalMenu).SexMaxOrder((*menu.GlobalMenu).MaxOrder + 1)
	menu.SetGlobalMenu()

	template.MenuPanelPjax((*menu.GlobalMenu).GetEditMenuList(), (*menu.GlobalMenu).GlobalMenuOption, buffer)

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}

// 修改菜单顺序
