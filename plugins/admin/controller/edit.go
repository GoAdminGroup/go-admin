package controller

import (
	"fmt"
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

func ShowForm(ctx *context.Context) {
	param := guard.GetShowFormParam(ctx)
	showForm(ctx, "", param.Panel, param.Id, param.GetUrl(), param.GetInfoUrl())
}

func showForm(ctx *context.Context, alert template2.HTML, panel table.Table, id string, url, infoUrl string) {

	formData, groupFormData, groupHeaders, title, description := panel.GetDataFromDatabaseWithId(id)

	fmt.Println("groupHeaders", groupHeaders)

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + aForm().
			SetContent(formData).
			SetGroupContent(groupFormData).
			SetGroupHeaders(groupHeaders).
			SetPrefix(config.PrefixFixSlash()).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetUrl(url).
			SetToken(auth.TokenHelper.AddToken()).
			SetInfoUrl(infoUrl).
			SetHeader(panel.GetForm().HeaderHtml).
			SetFooter(panel.GetForm().FooterHtml).
			GetContent(),
		Description: description,
		Title:       title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))

	ctx.Html(http.StatusOK, buf.String())
}

func EditForm(ctx *context.Context) {

	param := guard.GetEditFormParam(ctx)

	if param.HasAlert() {
		showForm(ctx, param.Alert, param.Panel, param.Id, param.GetUrl(), param.GetInfoUrl())
		return
	}

	// process uploading files, only support local storage for now.
	if len(param.MultiForm.File) > 0 {
		_, _ = file.GetFileEngine("local").Upload(param.MultiForm)
	}

	if param.IsManage() { // manager edit
		editManager(param.Value())
	} else if param.IsRole() { // role edit
		editRole(param.Value())
	} else {
		param.Panel.UpdateDataFromDatabase(param.Value())
	}

	table.RefreshTableList()

	panelInfo := param.Panel.GetDataFromDatabase(param.Path, param.Param)

	dataTable := aDataTable().
		SetInfoList(panelInfo.InfoList).
		SetPrimaryKey(param.Panel.GetPrimaryKey().Name).
		SetThead(panelInfo.Thead).
		SetNewUrl(param.GetNewUrl())

	if panelInfo.Editable {
		dataTable.SetEditUrl(param.GetEditUrl())
	}
	if panelInfo.Deletable {
		dataTable.SetDeleteUrl(param.GetDeleteUrl())
	}

	box := aBox().
		SetBody(dataTable.GetContent()).
		SetHeader(dataTable.GetDataTableHeader() + param.Panel.GetInfo().HeaderHtml).
		WithHeadBorder(false).
		SetFooter(param.Panel.GetInfo().FooterHtml + panelInfo.Paginator.GetContent()).
		GetContent()

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(true)
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(param.PreviousPath)))

	ctx.Html(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
}
