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
	getMenuInfoPanel(ctx)
	return
}

func ShowEditMenu(ctx *context.Context) {

	formData, title, description := table.List["menu"].GetDataFromDatabaseWithId(ctx.Query("id"))

	user := auth.Auth(ctx)

	js := `<script>
$('.icon').iconpicker({placement: 'bottomLeft'});
</script>`

	tmpl, tmplName := aTemplate().GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content: aForm().
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

	menu.SetGlobalMenu(auth.Auth(ctx).WithRoles().WithMenus())
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

	menu.GetGlobalMenu(user.WithRoles().WithMenus()).AddMaxOrder()
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

func getMenuInfoPanel(ctx *context.Context) {
	user := auth.Auth(ctx)

	menu.GlobalMenu.SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1))

	editUrl := config.Url("/menu/edit/show")
	deleteUrl := config.Url("/menu/delete")
	orderUrl := config.Url("/menu/order")

	tree := aTree().
		SetTree((menu.GetGlobalMenu(user)).GlobalMenuList).
		SetEditUrl(editUrl).
		SetDeleteUrl(deleteUrl).
		SetOrderUrl(orderUrl).
		GetContent()

	header := aTree().GetTreeHeader()
	box := aBox().SetHeader(header).SetBody(tree).GetContent()
	col1 := aCol().SetSize(map[string]string{"md": "6"}).SetContent(box).GetContent()

	newForm := aForm().
		SetPrefix(config.PREFIX).
		SetUrl(config.Url("/menu/new")).
		SetInfoUrl(config.Url("/menu")).
		SetTitle("New").
		SetContent(table.GetNewFormList(table.List["menu"].GetForm().FormList)).
		GetContent()

	col2 := aCol().SetSize(map[string]string{"md": "6"}).SetContent(newForm).GetContent()

	row := aRow().SetContent(col1 + col2).GetContent()

	menu.GlobalMenu.SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1))

	tmpl, tmplName := aTemplate().GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     row,
		Description: "Menus Manage",
		Title:       "Menus Manage",
	}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buf.String())
}
