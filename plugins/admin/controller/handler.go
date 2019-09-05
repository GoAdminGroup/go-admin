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

// 全局错误处理
func GlobalDeferHandler(ctx *context.Context) {

	logger.Access(ctx)

	// TODO: sqlite will cause a panic. database is locked.
	if config.DATABASE.GetDefault().DRIVER != "sqlite" {
		RecordOperationLog(ctx)
	}

	if err := recover(); err != nil {
		logger.Error(err)
		logger.Error(string(debug.Stack()[:]))

		var (
			errMsg     string
			mysqlError *mysql.MySQLError
			ok         bool
			aerr       error
		)
		if errMsg, ok = err.(string); !ok {
			if mysqlError, ok = err.(*mysql.MySQLError); ok {
				errMsg = mysqlError.Error()
			} else if aerr, ok = err.(error); ok {
				errMsg = aerr.Error()
			}
		}

		alert := template.Get(config.THEME).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML(errMsg)).GetContent()

		if ok, _ = regexp.Match("/edit(.*)", []byte(ctx.Path())); ok {

			prefix := ctx.Query("prefix")

			formData, title, description := table.List[prefix].GetDataFromDatabaseWithId(ctx.Query("id"))

			queryParam := parameter.GetParam(ctx.Request.URL.Query()).GetRouteParamStr()

			user := auth.Auth(ctx)

			tmpl, tmplName := template.Get(config.THEME).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
			buf := template.Excecute(tmpl, tmplName, user, types.Panel{
				Content: alert + template.Get(config.THEME).Form().
					SetContent(formData).
					SetPrefix(config.PREFIX).
					SetUrl(config.PREFIX + "/edit/" + prefix).
					SetToken(auth.TokenHelper.AddToken()).
					SetInfoUrl(config.PREFIX + "/info/" + prefix + queryParam).
					GetContent(),
				Description: description,
				Title:       title,
			}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1)))
			ctx.Html(http.StatusOK, buf.String())
			ctx.AddHeader(constant.PjaxUrlHeader, config.PREFIX+"/info/"+prefix+"/new"+queryParam)
			return
		}

		if ok, _ = regexp.Match("/new(.*)", []byte(ctx.Path())); ok {
			prefix := ctx.Query("prefix")

			queryParam := parameter.GetParam(ctx.Request.URL.Query()).GetRouteParamStr()

			user := auth.Auth(ctx)

			tmpl, tmplName := template.Get(config.THEME).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
			buf := template.Excecute(tmpl, tmplName, user, types.Panel{
				Content: alert + template.Get(config.THEME).Form().
					SetPrefix(config.PREFIX).
					SetContent(table.GetNewFormList(table.List[prefix].GetForm().FormList)).
					SetUrl(config.PREFIX + "/new/" + prefix).
					SetToken(auth.TokenHelper.AddToken()).
					SetInfoUrl(config.PREFIX + "/info/" + prefix + queryParam).
					GetContent(),
				Description: table.List[prefix].GetForm().Description,
				Title:       table.List[prefix].GetForm().Title,
			}, config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), config.PREFIX, "", 1)))
			ctx.Html(http.StatusOK, buf.String())
			ctx.AddHeader(constant.PjaxUrlHeader, config.PREFIX+"/info/"+prefix+"/new"+queryParam)
			return
		}

		response.Error(ctx, errMsg)
		return
	}
}
