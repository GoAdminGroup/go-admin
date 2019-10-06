package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/modules/constant"
	"github.com/chenhg5/go-admin/plugins/admin/modules/file"
	"github.com/chenhg5/go-admin/plugins/admin/modules/guard"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	template2 "html/template"
	"net/http"
)

func ShowNewForm(ctx *context.Context) {
	param := guard.GetShowNewFormParam(ctx)
	showNewForm(ctx, "", param.Panel, param.GetUrl(), param.GetInfoUrl())
}

func showNewForm(ctx *context.Context, alert template2.HTML, panel table.Table, url, infoUrl string) {

	user := auth.Auth(ctx)

	table.RefreshTableList()

	formList, groupFormList, groupHeaders := table.GetNewFormList(panel.GetForm().GroupHeaders, panel.GetForm().Group,
		panel.GetForm().FormList, panel.GetPrimaryKey().Name)

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + aForm().
			SetPrefix(config.PrefixFixSlash()).
			SetContent(formList).
			SetGroupContent(groupFormList).
			SetGroupHeaders(groupHeaders).
			SetUrl(url).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetToken(auth.TokenHelper.AddToken()).
			SetTitle("New").
			SetInfoUrl(infoUrl).
			SetHeader(panel.GetForm().HeaderHtml).
			SetFooter(panel.GetForm().FooterHtml).
			GetContent(),
		Description: panel.GetForm().Description,
		Title:       panel.GetForm().Title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))
	ctx.Html(http.StatusOK, buf.String())
}

func NewForm(ctx *context.Context) {

	param := guard.GetNewFormParam(ctx)

	table.RefreshTableList()

	if param.HasAlert() {
		showNewForm(ctx, param.Alert, param.Panel, param.GetUrl(), param.GetInfoUrl())
		return
	}

	// process uploading files, only support local storage
	if len(param.MultiForm.File) > 0 {
		_, _ = file.GetFileEngine("local").Upload(param.MultiForm)
	}

	if param.IsManage() { // manager edit
		newManager(param.Value())
	} else if param.IsRole() { // role edit
		newRole(param.Value())
	} else {
		param.Panel.InsertDataFromDatabase(param.Value())
	}

	panelInfo := param.Panel.GetDataFromDatabase(param.Path, param.Param)

	dataTable := aDataTable().
		SetInfoList(panelInfo.InfoList).
		SetThead(panelInfo.Thead).
		SetPrimaryKey(param.Panel.GetPrimaryKey().Name).
		SetEditUrl(param.GetEditUrl()).
		SetNewUrl(param.GetNewUrl()).
		SetDeleteUrl(param.GetDeleteUrl())

	box := aBox().
		SetBody(dataTable.GetContent()).
		SetHeader(dataTable.GetDataTableHeader() + param.Panel.GetInfo().HeaderHtml).
		WithHeadBorder(false).
		SetFooter(param.Panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
		GetContent()

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(true)
	buffer := template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(param.PreviousPath)))

	ctx.Html(http.StatusOK, buffer.String())
	ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
}
