package controller

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
)

// ShowMenu show menu info page.
func ShowMenu(ctx *context.Context) {
	getMenuInfoPanel(ctx, "")
}

// ShowNewMenu show new menu page.
func ShowNewMenu(ctx *context.Context) {

	panel := table.Get("menu")

	formData, groupFormData, groupHeaders := table.GetNewFormList(panel.GetForm().TabHeaders,
		panel.GetForm().TabGroups,
		panel.GetForm().FieldList)

	user := auth.Auth(ctx)

	js := `<script>
$('.icon').iconpicker({placement: 'bottomLeft'});
</script>`

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: formContent(aForm().
			SetContent(formData).
			SetTabContents(groupFormData).
			SetTabHeaders(groupHeaders).
			SetPrefix(config.PrefixFixSlash()).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetUrl(config.Url("/menu/edit")).
			SetToken(authSrv().AddToken()).
			SetOperationFooter(formFooter()).
			SetInfoUrl(config.Url("/menu"))) +
			template2.HTML(js),
		Description: panel.GetForm().Description,
		Title:       panel.GetForm().Title,
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))

	ctx.HTML(http.StatusOK, buf.String())
}

// ShowEditMenu show edit menu page.
func ShowEditMenu(ctx *context.Context) {

	if ctx.Query("id") == "" {
		getMenuInfoPanel(ctx, template.Get(config.Theme).Alert().
			SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> `+language.Get("error")+`!`)).
			SetTheme("warning").
			SetContent(template2.HTML("wrong id")).
			GetContent())
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
		return
	}

	formData, groupFormData, groupHeaders, title, description, _ := table.Get("menu").GetDataFromDatabaseWithId(ctx.Query("id"))

	user := auth.Auth(ctx)

	js := `<script>
$('.icon').iconpicker({placement: 'bottomLeft'});
</script>`

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: formContent(aForm().
			SetContent(formData).
			SetTabContents(groupFormData).
			SetTabHeaders(groupHeaders).
			SetPrefix(config.PrefixFixSlash()).
			SetPrimaryKey(table.Get("menu").GetPrimaryKey().Name).
			SetUrl(config.Url("/menu/edit")).
			SetOperationFooter(formFooter()).
			SetToken(authSrv().AddToken()).
			SetInfoUrl(config.Url("/menu"))) + template2.HTML(js),
		Description: description,
		Title:       title,
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))

	ctx.HTML(http.StatusOK, buf.String())
}

// DeleteMenu delete the menu of given id.
func DeleteMenu(ctx *context.Context) {
	models.MenuWithId(guard.GetMenuDeleteParam(ctx).Id).SetConn(db.GetConnection(services)).Delete()
	table.RefreshTableList()
	response.Ok(ctx)
}

// EditMenu edit the menu of given id.
func EditMenu(ctx *context.Context) {

	param := guard.GetMenuEditParam(ctx)

	if param.HasAlert() {
		getMenuInfoPanel(ctx, param.Alert)
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
		return
	}

	menuModel := models.MenuWithId(param.Id).SetConn(db.GetConnection(services))

	menuModel.DeleteRoles()
	for _, roleId := range param.Roles {
		menuModel.AddRole(roleId)
	}
	table.RefreshTableList()

	menuModel.Update(param.Title, param.Icon, param.Uri, param.Header, param.ParentId)

	getMenuInfoPanel(ctx, "")
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
}

// NewMenu create a new menu item.
func NewMenu(ctx *context.Context) {

	param := guard.GetMenuNewParam(ctx)

	if param.HasAlert() {
		getMenuInfoPanel(ctx, param.Alert)
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
		return
	}

	user := auth.Auth(ctx)

	menuModel := models.Menu().SetConn(db.GetConnection(services)).
		New(param.Title, param.Icon, param.Uri, param.Header, param.ParentId, (menu.GetGlobalMenu(user, conn)).MaxOrder+1)

	for _, roleId := range param.Roles {
		menuModel.AddRole(roleId)
	}

	menu.GetGlobalMenu(user, conn).AddMaxOrder()
	table.RefreshTableList()

	getMenuInfoPanel(ctx, "")
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
}

// MenuOrder change the order of menu items.
func MenuOrder(ctx *context.Context) {

	var data []map[string]interface{}
	_ = json.Unmarshal([]byte(ctx.FormValue("_order")), &data)

	models.Menu().SetConn(db.GetConnection(services)).ResetOrder(data)

	response.Ok(ctx)
}

func getMenuInfoPanel(ctx *context.Context, alert template2.HTML) {
	user := auth.Auth(ctx)

	table.RefreshTableList()

	editUrl := config.Url("/menu/edit/show")
	deleteUrl := config.Url("/menu/delete")
	orderUrl := config.Url("/menu/order")

	tree := aTree().
		SetTree((menu.GetGlobalMenu(user, conn)).List).
		SetEditUrl(editUrl).
		SetUrlPrefix(config.Prefix()).
		SetDeleteUrl(deleteUrl).
		SetOrderUrl(orderUrl).
		GetContent()

	header := aTree().GetTreeHeader()
	box := aBox().SetHeader(header).SetBody(tree).GetContent()
	col1 := aCol().SetSize(map[string]string{"md": "6"}).SetContent(box).GetContent()

	list := table.Get("menu")

	formList, groupFormList, groupHeaders := table.GetNewFormList(list.GetForm().TabHeaders, list.GetForm().TabGroups,
		list.GetForm().FieldList)

	newForm := menuFormContent(aForm().
		SetPrefix(config.PrefixFixSlash()).
		SetUrl(config.Url("/menu/new")).
		SetPrimaryKey(table.Get("menu").GetPrimaryKey().Name).
		SetToken(authSrv().AddToken()).
		SetInfoUrl(config.Url("/menu")).
		SetOperationFooter(formFooter()).
		SetTitle("New").
		SetContent(formList).
		SetTabContents(groupFormList).
		SetTabHeaders(groupHeaders))

	col2 := aCol().SetSize(map[string]string{"md": "6"}).SetContent(newForm).GetContent()

	row := aRow().SetContent(col1 + col2).GetContent()

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     alert + row,
		Description: "Menus Manage",
		Title:       "Menus Manage",
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))

	ctx.HTML(http.StatusOK, buf.String())
}
