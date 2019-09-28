package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/logger"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/modules/constant"
	"github.com/chenhg5/go-admin/plugins/admin/modules/parameter"
	"github.com/chenhg5/go-admin/plugins/admin/modules/response"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/go-sql-driver/mysql"
	template2 "html/template"
	"net/http"
	"regexp"
	"runtime/debug"
	"strings"
)

func GlobalDeferHandler(ctx *context.Context) {

	logger.Access(ctx)

	RecordOperationLog(ctx)

	if err := recover(); err != nil {
		logger.Error(err)
		logger.Error(string(debug.Stack()[:]))

		var (
			errMsg     string
			mysqlError *mysql.MySQLError
			ok         bool
			e          error
		)

		if errMsg, ok = err.(string); !ok {
			if mysqlError, ok = err.(*mysql.MySQLError); ok {
				errMsg = mysqlError.Error()
			} else if e, ok = err.(error); ok {
				errMsg = e.Error()
			}
		}

		if ok, _ = regexp.MatchString("/edit(.*)", ctx.Path()); ok {
			setFormWithReturnErrMessage(ctx, errMsg, "edit")
			return
		}
		if ok, _ = regexp.MatchString("/new(.*)", ctx.Path()); ok {
			setFormWithReturnErrMessage(ctx, errMsg, "new")
			return
		}

		response.Error(ctx, errMsg)
		return
	}
}

func setFormWithReturnErrMessage(ctx *context.Context, errMsg string, kind string) {

	alert := aAlert().
		SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
		SetTheme("warning").
		SetContent(template2.HTML(errMsg)).
		GetContent()

	var (
		formData           []types.Form
		title, description string
		prefix             = ctx.Query("prefix")
		panel              = table.List[prefix]
	)

	if kind == "edit" {
		formData, title, description = table.List[prefix].GetDataFromDatabaseWithId(ctx.Query("id"))
	} else {
		formData = table.GetNewFormList(panel.GetForm().FormList, panel.GetPrimaryKey().Name)
		title = panel.GetForm().Title
		description = panel.GetForm().Description
	}

	queryParam := parameter.GetParam(ctx.Request.URL.Query()).GetRouteParamStr()

	user := auth.Auth(ctx)

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + aForm().
			SetContent(formData).
			SetTitle(template2.HTML(strings.Title(kind))).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetPrefix(config.PrefixFixSlash()).
			SetUrl(config.Url("/"+kind+"/"+prefix)).
			SetToken(auth.TokenHelper.AddToken()).
			SetInfoUrl(config.Url("/info/"+prefix+queryParam)).
			GetContent(),
		Description: description,
		Title:       title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))
	ctx.Html(http.StatusOK, buf.String())
	ctx.AddHeader(constant.PjaxUrlHeader, config.Url("/info/"+prefix+"/"+kind+queryParam))
}
