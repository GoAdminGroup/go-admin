package controller

import (
	"encoding/json"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/plugins/admin/modules/constant"
	"github.com/chenhg5/go-admin/plugins/admin/modules/response"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"strings"
)

func ShowMenu(ctx *context.Context) {
	GetMenuInfoPanel(ctx)
	return
}

func ShowEditMenu(ctx *context.Context) {

	formData, title, description := table.List["menu"].GetDataFromDatabaseWithId(ctx.Query("id"))

	user := auth.Auth(ctx)

	js := `<script>
$('.icon').iconpicker({placement: 'bottomLeft'});
</script>`

	tmpl, tmplName := template.Get(config.THEME).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content: template.Get(config.THEME).Form().
			SetContent(formData).
			SetPrefix(config.PREFIX).
			SetUrl(config.Url("/menu/edit")).
			SetToken(auth.TokenHelper.AddToken()).
			SetInfoUrl(config.Url("/menu")).
			GetContent() + template2.HTML(js),
		Description: description,
		Title:       title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buf.String())
}

func DeleteMenu(ctx *context.Context) {

	models.MenuWithId(ctx.Query("id")).Delete()

	menu.SetGlobalMenu(auth.Auth(ctx).UpdateMenus())
	table.RefreshTableList()
	response.Ok(ctx)
}

func EditMenu(ctx *context.Context) {
	id := ctx.FormValue("id")
	title := ctx.FormValue("title")
	parentId := ctx.FormValue("parent_id")
	if parentId == "" {
		parentId = "0"
	}
	icon := ctx.FormValue("icon")
	uri := ctx.FormValue("uri")

	menuModel := models.MenuWithId(id)

	for _, roleId := range ctx.Request.Form["roles[]"] {
		menuModel.AddRole(roleId)
	}

	menuModel.Update(title, parentId, icon, uri)

	menu.SetGlobalMenu(auth.Auth(ctx))

	GetMenuInfoPanel(ctx)
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
}

func NewMenu(ctx *context.Context) {

	title := ctx.FormValue("title")
	parentId := ctx.FormValue("parent_id")
	if parentId == "" {
		parentId = "0"
	}
	icon := ctx.FormValue("icon")
	uri := ctx.FormValue("uri")

	user := auth.Auth(ctx)

	menuModel := models.Menu().New(title, parentId, icon, uri, (menu.GetGlobalMenu(user)).MaxOrder+1)

	roles := ctx.Request.Form["roles[]"]

	for _, roleId := range roles {
		menuModel.AddRole(roleId)
	}

	menu.GetGlobalMenu(user.UpdateMenus()).AddMaxOrder()
	table.RefreshTableList()

	GetMenuInfoPanel(ctx)
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
}

func MenuOrder(ctx *context.Context) {

	var data []map[string]interface{}
	_ = json.Unmarshal([]byte(ctx.FormValue("_order")), &data)

	models.Menu().ResetOrder(data)
	menu.SetGlobalMenu(auth.Auth(ctx))

	response.Ok(ctx)
	return
}

func GetMenuInfoPanel(ctx *context.Context) {
	user := auth.Auth(ctx)

	menu.GlobalMenu.SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1))

	editUrl := config.Url("/menu/edit/show")
	deleteUrl := config.Url("/menu/delete")
	orderUrl := config.Url("/menu/order")

	tree := template.Get(config.THEME).Tree().
		SetTree((menu.GetGlobalMenu(user)).GlobalMenuList).
		SetEditUrl(editUrl).
		SetDeleteUrl(deleteUrl).
		SetOrderUrl(orderUrl).
		GetContent()

	header := template.Get(config.THEME).Tree().GetTreeHeader()
	box := template.Get(config.THEME).Box().SetHeader(header).SetBody(tree).GetContent()
	col1 := template.Get(config.THEME).Col().SetSize(map[string]string{"md": "6"}).SetContent(box).GetContent()

	newForm := template.Get(config.THEME).Form().
		SetPrefix(config.PREFIX).
		SetUrl(config.Url("/menu/new")).
		SetInfoUrl(config.Url("/menu")).
		SetTitle("New").
		SetContent(table.GetNewFormList(table.List["menu"].GetForm().FormList)).
		GetContent()

	col2 := template.Get(config.THEME).Col().SetSize(map[string]string{"md": "6"}).SetContent(newForm).GetContent()

	row := template.Get(config.THEME).Row().SetContent(col1 + col2).GetContent()

	menu.GlobalMenu.SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1))

	tmpl, tmplName := template.Get(config.THEME).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     row,
		Description: "Menus Manage",
		Title:       "Menus Manage",
	}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buf.String())
}
