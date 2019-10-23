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

func ShowForm(ctx *context.Context) {
	param := guard.GetShowFormParam(ctx)
	showForm(ctx, "", param.Panel, param.Id, param.GetUrl(), param.GetInfoUrl())
}

func showForm(ctx *context.Context, alert template2.HTML, panel table.Table, id string, url, infoUrl string) {

	formData, groupFormData, groupHeaders, title, description, err := panel.GetDataFromDatabaseWithId(id)

	if err != nil && alert == "" {
		alert = aAlert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
	}

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + aForm().
			SetContent(formData).
			SetTabContents(groupFormData).
			SetTabHeaders(groupHeaders).
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
		editManager(param.Value())
	} else if param.IsRole() { // role edit
		editRole(param.Value())
	} else {
		err := param.Panel.UpdateDataFromDatabase(param.Value())
		if err != nil {
			alert := aAlert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
				SetTheme("warning").
				SetContent(template2.HTML(err.Error())).
				GetContent()
			showForm(ctx, alert, param.Panel, param.Id, param.GetUrl(), param.GetInfoUrl())
			return
		}
	}

	table.RefreshTableList()

	editUrl := modules.AorB(param.Panel.GetEditable(), param.GetEditUrl(), "")
	deleteUrl := modules.AorB(param.Panel.GetDeletable(), param.GetDeleteUrl(), "")
	exportUrl := modules.AorB(param.Panel.GetExportable(), param.GetExportUrl(), "")
	newUrl := modules.AorB(param.Panel.GetCanAdd(), param.GetNewUrl(), "")
	infoUrl := param.GetInfoUrl()

	buf := showTable(ctx, param.Panel, param.Path, param.Param, exportUrl, newUrl, deleteUrl, infoUrl, editUrl)

	ctx.Html(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, param.PreviousPath)
}
