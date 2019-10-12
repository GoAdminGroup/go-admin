package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/file"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
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

	formList, groupFormList, groupHeaders := table.GetNewFormList(panel.GetForm().TabHeaders, panel.GetForm().TabGroups,
		panel.GetForm().FieldList, panel.GetPrimaryKey().Name)

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + aForm().
			SetPrefix(config.PrefixFixSlash()).
			SetContent(formList).
			SetTabContents(groupFormList).
			SetTabHeaders(groupHeaders).
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
		err := file.GetFileEngine(config.FileUploadEngine.Name).Upload(param.MultiForm)
		if err != nil {
			alert := aAlert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
				SetTheme("warning").
				SetContent(template2.HTML(err.Error())).
				GetContent()
			showForm(ctx, alert, param.Panel, param.Id, param.GetUrl(), param.GetInfoUrl())
			return
		}
	}

	if param.IsManage() { // manager edit
		newManager(param.Value())
	} else if param.IsRole() { // role edit
		newRole(param.Value())
	} else {
		err := param.Panel.InsertDataFromDatabase(param.Value())
		if err != nil {
			alert := aAlert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
				SetTheme("warning").
				SetContent(template2.HTML(err.Error())).
				GetContent()
			showForm(ctx, alert, param.Panel, param.Id, param.GetUrl(), param.GetInfoUrl())
			return
		}
	}

	editUrl := modules.AorB(param.Panel.GetEditable(), param.GetEditUrl(), "")
	deleteUrl := modules.AorB(param.Panel.GetDeletable(), param.GetDeleteUrl(), "")
	exportUrl := modules.AorB(param.Panel.GetExportable(), param.GetExportUrl(), "")
	newUrl := modules.AorB(param.Panel.GetCanAdd(), param.GetNewUrl(), "")
	infoUrl := param.GetInfoUrl()

	buf := showTable(ctx, param.Panel, param.Path, param.Param, exportUrl, newUrl, deleteUrl, infoUrl, editUrl)

	ctx.Html(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
}
