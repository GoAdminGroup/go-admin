package controller

import (
	"fmt"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/go-sql-driver/mysql"
	"github.com/mgutz/ansi"
	template2 "html/template"
	"log"
	"net/http"
	"regexp"
	"runtime/debug"
	"strconv"
)

// 全局错误处理
func GlobalDeferHandler(ctx *context.Context) {

	log.Println("[GoAdmin]",
		ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
		ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
		ctx.Path())

	RecordOperationLog(ctx)

	if err := recover(); err != nil {
		fmt.Println(err)
		fmt.Println(string(debug.Stack()[:]))

		var (
			errMsg     string
			mysqlError *mysql.MySQLError
			ok         bool
		)
		if errMsg, ok = err.(string); !ok {
			if mysqlError, ok = err.(*mysql.MySQLError); ok {
				errMsg = mysqlError.Error()
			} else {
				errMsg = fmt.Sprint("%v", err)
			}
		}

		alert := template.Get(Config.THEME).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML(errMsg)).GetContent()

		if ok, _ = regexp.Match("/edit(.*)", []byte(ctx.Path())); ok {

			user := ctx.UserValue["user"].(auth.User)

			prefix := ctx.Query("prefix")

			id := ctx.Query("id")

			formData, title, description := models.TableList[prefix].GetDataFromDatabaseWithId(prefix, id)

			path := ctx.Path()
			menu.GlobalMenu.SetActiveClass(path)

			ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

			queryParam := models.GetParam(ctx.Request.URL.Query()).GetRouteParamStr()

			tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
			buf := template.Excecute(tmpl, tmplName, user, types.Panel{
				Content: alert + template.Get(Config.THEME).Form().
					SetContent(formData).
					SetPrefix(Config.PREFIX).
					SetUrl(Config.PREFIX+"/edit/"+prefix).
					SetToken(auth.TokenHelper.AddToken()).
					SetInfoUrl(Config.PREFIX+"/info/"+prefix+queryParam).
					GetContent(),
				Description: description,
				Title:       title,
			}, Config)
			ctx.WriteString(buf.String())
			ctx.AddHeader("X-PJAX-URL", Config.PREFIX+"/info/"+prefix+"/new"+queryParam)
			return
		}

		if ok, _ = regexp.Match("/new(.*)", []byte(ctx.Path())); ok {
			prefix := ctx.Query("prefix")

			user := ctx.UserValue["user"].(auth.User)

			path := ctx.Path()
			menu.GlobalMenu.SetActiveClass(path)

			ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

			queryParam := models.GetParam(ctx.Request.URL.Query()).GetRouteParamStr()

			tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
			buf := template.Excecute(tmpl, tmplName, user, types.Panel{
				Content: alert + template.Get(Config.THEME).Form().
					SetPrefix(Config.PREFIX).
					SetContent(models.GetNewFormList(models.TableList[prefix].Form.FormList)).
					SetUrl(Config.PREFIX+"/new/"+prefix).
					SetToken(auth.TokenHelper.AddToken()).
					SetInfoUrl(Config.PREFIX+"/info/"+prefix+queryParam).
					GetContent(),
				Description: models.TableList[prefix].Form.Description,
				Title:       models.TableList[prefix].Form.Title,
			}, Config)
			ctx.WriteString(buf.String())
			ctx.AddHeader("X-PJAX-URL", Config.PREFIX+"/info/"+prefix+"/new"+queryParam)
			return
		}

		ctx.Json(http.StatusInternalServerError, map[string]interface{}{
			"code": 500,
			"msg":  errMsg,
		})
		return
	}
}
