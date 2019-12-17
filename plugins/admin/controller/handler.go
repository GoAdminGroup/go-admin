package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"
)

// GlobalDeferHandler is a global error handler of admin plugin.
func GlobalDeferHandler(ctx *context.Context) {

	logger.Access(ctx)

	RecordOperationLog(ctx)

	if err := recover(); err != nil {
		logger.Error(err)
		logger.Error(string(debug.Stack()[:]))

		var (
			errMsg string
			ok     bool
			e      error
		)

		if errMsg, ok = err.(string); !ok {
			if e, ok = err.(error); ok {
				errMsg = e.Error()
			}
		}

		if errMsg == "" {
			errMsg = "system error"
		}

		if ok, _ = regexp.MatchString("/edit(.*)", ctx.Path()); ok {
			setFormWithReturnErrMessage(ctx, errMsg, "edit")
			return
		}
		if ok, _ = regexp.MatchString("/new(.*)", ctx.Path()); ok {
			setFormWithReturnErrMessage(ctx, errMsg, "new")
			return
		}

		alert := aAlert().
			SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").
			SetContent(template2.HTML(errMsg)).
			GetContent()

		user := auth.Auth(ctx)

		tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
		buf := template.Execute(tmpl, tmplName, user, types.Panel{
			Content:     alert,
			Description: "error",
			Title:       "error",
		}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))
		ctx.HTML(http.StatusOK, buf.String())
		return
	}
}

func setFormWithReturnErrMessage(ctx *context.Context, errMsg string, kind string) {

	alert := aAlert().
		SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
		SetTheme("warning").
		SetContent(template2.HTML(errMsg)).
		GetContent()

	var (
		formData           []types.FormField
		groupFormData      [][]types.FormField
		groupHeaders       []string
		title, description string
		prefix             = ctx.Query("__prefix")
		panel              = table.Get(prefix)
	)

	if kind == "edit" {
		id := ctx.Query("id")
		if id == "" {
			id = ctx.Request.MultipartForm.Value[panel.GetPrimaryKey().Name][0]
		}
		formData, groupFormData, groupHeaders, title, description, _ = table.Get(prefix).GetDataFromDatabaseWithId(id)
	} else {
		formData, groupFormData, groupHeaders = table.GetNewFormList(panel.GetForm().TabHeaders, panel.GetForm().TabGroups,
			panel.GetForm().FieldList)
		title = panel.GetForm().Title
		description = panel.GetForm().Description
	}

	queryParam := parameter.GetParam(ctx.Request.URL.Query(), panel.GetInfo().DefaultPageSize,
		panel.GetPrimaryKey().Name, panel.GetInfo().GetSort()).GetRouteParamStr()

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + formContent(aForm().
			SetContent(formData).
			SetTabContents(groupFormData).
			SetTabHeaders(groupHeaders).
			SetTitle(template2.HTML(strings.Title(kind))).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetPrefix(config.PrefixFixSlash()).
			SetUrl(config.Url("/"+kind+"/"+prefix)).
			SetToken(authSrv().AddToken()).
			SetOperationFooter(formFooter()).
			SetHeader(panel.GetForm().HeaderHtml).
			SetFooter(panel.GetForm().FooterHtml).
			SetInfoUrl(config.Url("/info/"+prefix+queryParam))),
		Description: description,
		Title:       title,
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))
	ctx.HTML(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/info/"+prefix+"/"+kind+queryParam))
}
