package controller

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
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

func ShowMenu(ctx *context.Context) {
	getMenuInfoPanel(ctx, "")
}

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

	formData, groupFormData, groupHeaders, title, description, _ := table.List["menu"].GetDataFromDatabaseWithId(ctx.Query("id"))

	user := auth.Auth(ctx)

	js := `<script>
$('.icon').iconpicker({placement: 'bottomLeft'});
</script>`

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: aForm().
			SetContent(formData).
			SetTabContents(groupFormData).
			SetTabHeaders(groupHeaders).
			SetPrefix(config.PrefixFixSlash()).
			SetPrimaryKey(table.List["menu"].GetPrimaryKey().Name).
			SetUrl(config.Url("/menu/edit")).
			SetToken(auth.TokenHelper.AddToken()).
			SetInfoUrl(config.Url("/menu")).
			GetContent() + template2.HTML(js),
		Description: description,
		Title:       title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))

	ctx.Html(http.StatusOK, buf.String())
}

func DeleteMenu(ctx *context.Context) {
	models.MenuWithId(guard.GetMenuDeleteParam(ctx).Id).Delete()
	table.RefreshTableList()
	response.Ok(ctx)
}

func EditMenu(ctx *context.Context) {

	param := guard.GetMenuEditParam(ctx)

	if param.HasAlert() {
		getMenuInfoPanel(ctx, param.Alert)
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
		return
	}

	menuModel := models.MenuWithId(param.Id)

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

func NewMenu(ctx *context.Context) {

	param := guard.GetMenuNewParam(ctx)

	if param.HasAlert() {
		getMenuInfoPanel(ctx, param.Alert)
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
		return
	}

	user := auth.Auth(ctx)

	menuModel := models.Menu().New(param.Title, param.Icon, param.Uri, param.Header, param.ParentId, (menu.GetGlobalMenu(user)).MaxOrder+1)

	for _, roleId := range param.Roles {
		menuModel.AddRole(roleId)
	}

	menu.GetGlobalMenu(user).AddMaxOrder()
	table.RefreshTableList()

	getMenuInfoPanel(ctx, "")
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/menu"))
}

func MenuOrder(ctx *context.Context) {

	var data []map[string]interface{}
	_ = json.Unmarshal([]byte(ctx.FormValue("_order")), &data)

	models.Menu().ResetOrder(data)

	response.Ok(ctx)
}

func getMenuInfoPanel(ctx *context.Context, alert template2.HTML) {
	user := auth.Auth(ctx)

	table.RefreshTableList()

	editUrl := config.Url("/menu/edit/show")
	deleteUrl := config.Url("/menu/delete")
	orderUrl := config.Url("/menu/order")

	tree := aTree().
		SetTree((menu.GetGlobalMenu(user)).List).
		SetEditUrl(editUrl).
		SetDeleteUrl(deleteUrl).
		SetOrderUrl(orderUrl).
		GetContent()

	header := aTree().GetTreeHeader()
	box := aBox().SetHeader(header).SetBody(tree).GetContent()
	col1 := aCol().SetSize(map[string]string{"md": "6"}).SetContent(box).GetContent()

	list := table.List["menu"]

	formList, groupFormList, groupHeaders := table.GetNewFormList(list.GetForm().TabHeaders, list.GetForm().TabGroups,
		list.GetForm().FieldList, list.GetPrimaryKey().Name)

	newForm := aForm().
		SetPrefix(config.PrefixFixSlash()).
		SetUrl(config.Url("/menu/new")).
		SetPrimaryKey(table.List["menu"].GetPrimaryKey().Name).
		SetToken(auth.TokenHelper.AddToken()).
		SetInfoUrl(config.Url("/menu")).
		SetTitle("New").
		SetContent(formList).
		SetTabContents(groupFormList).
		SetTabHeaders(groupHeaders).
		GetContent()

	col2 := aCol().SetSize(map[string]string{"md": "6"}).SetContent(newForm).GetContent()

	row := aRow().SetContent(col1 + col2).GetContent()

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     alert + row,
		Description: "Menus Manage",
		Title:       "Menus Manage",
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))

	ctx.Html(http.StatusOK, buf.String())
}
