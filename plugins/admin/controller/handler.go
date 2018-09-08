package controller

import (
	"bytes"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/mgutz/ansi"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"log"
	"regexp"
	"runtime/debug"
	"strconv"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template"
	"net/http"
	template2 "html/template"
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
				errMsg = "系统错误"
			}
		}

		alert := template.Get(Config.THEME).Alert().SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
			SetTheme("warning").SetContent(template2.HTML(errMsg)).GetContent()

		if ok, _ = regexp.Match("/edit(.*)", []byte(ctx.Path())); ok {

			user := ctx.UserValue["user"].(auth.User)

			prefix := ctx.Request.URL.Query().Get("prefix")

			id := ctx.Request.URL.Query().Get("id")

			formData, title, description := models.GlobalTableList[prefix].GetDataFromDatabaseWithId(prefix, id)

			tmpl, tmplName := template.Get("adminlte").GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

			path := ctx.Path()
			menu.GlobalMenu.SetActiveClass(path)

			page := ctx.Request.URL.Query().Get("page")
			if page == "" {
				page = "1"
			}
			pageSize := ctx.Request.URL.Query().Get("pageSize")
			if pageSize == "" {
				pageSize = "10"
			}

			sortField := ctx.Request.URL.Query().Get("sort")
			if sortField == "" {
				sortField = "id"
			}
			sortType := ctx.Request.URL.Query().Get("sort_type")
			if sortType == "" {
				sortType = "desc"
			}

			ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

			buf := new(bytes.Buffer)
			tmpl.ExecuteTemplate(buf, tmplName, types.Page{
				User: user,
				Menu: *menu.GlobalMenu,
				System: types.SystemInfo{
					"0.0.1",
				},
				Panel: types.Panel{
					Content: alert + template.Get(Config.THEME).Form().
						SetContent(formData).
						SetPrefix(Config.ADMIN_PREFIX).
						SetUrl(Config.ADMIN_PREFIX + "/edit/" + prefix).
						SetToken(auth.TokenHelper.AddToken()).
						SetInfoUrl(Config.ADMIN_PREFIX + "/info/" + prefix + "?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType).
						GetContent(),
					Description: description,
					Title:       title,
				},
				AssertRootUrl: Config.ADMIN_PREFIX,
			})
			ctx.WriteString(buf.String())
		}

		if ok, _ = regexp.Match("/new(.*)", []byte(ctx.Path())); ok {
			prefix := ctx.Request.URL.Query().Get("prefix")

			user := ctx.UserValue["user"].(auth.User)

			tmpl, tmplName := template.Get("adminlte").GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

			path := ctx.Path()
			menu.GlobalMenu.SetActiveClass(path)

			page := ctx.Request.URL.Query().Get("page")
			if page == "" {
				page = "1"
			}
			pageSize := ctx.Request.URL.Query().Get("pageSize")
			if pageSize == "" {
				pageSize = "10"
			}

			sortField := ctx.Request.URL.Query().Get("sort")
			if sortField == "" {
				sortField = "id"
			}
			sortType := ctx.Request.URL.Query().Get("sort_type")
			if sortType == "" {
				sortType = "desc"
			}

			ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

			buf := new(bytes.Buffer)
			tmpl.ExecuteTemplate(buf, tmplName, types.Page{
				User: user,
				Menu: *menu.GlobalMenu,
				System: types.SystemInfo{
					"0.0.1",
				},
				Panel: types.Panel{
					Content: alert + template.Get(Config.THEME).Form().
						SetPrefix(Config.ADMIN_PREFIX).
						SetContent(models.GetNewFormList(models.GlobalTableList[prefix].Form.FormList)).
						SetUrl(Config.ADMIN_PREFIX + "/new/" + prefix).
						SetToken(auth.TokenHelper.AddToken()).
						SetInfoUrl(Config.ADMIN_PREFIX + "/info/" + prefix + "?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType).
						GetContent(),
					Description: models.GlobalTableList[prefix].Form.Description,
					Title:       models.GlobalTableList[prefix].Form.Title,
				},
				AssertRootUrl: Config.ADMIN_PREFIX,
			})
			ctx.WriteString(buf.String())
			return
		}

		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetContentType("application/json")
		ctx.WriteString(`{"code":500, "msg":"` + errMsg + `"}`)
		return
	}
}
