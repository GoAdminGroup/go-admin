package controller

import (
	"bytes"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/mgutz/ansi"
	"github.com/valyala/fasthttp"
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
)

// 全局错误处理
func GlobalDeferHandler(ctx *context.Context) {

	log.Println("[GoAdmin]",
		ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
		ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
		string(ctx.Path()))

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

		if ok, _ = regexp.Match("/edit(.*)", []byte(ctx.Path())); ok {
			prefix := ctx.Request.URL.Query().Get("prefix")

			form := ctx.Request.MultipartForm

			id := (*form).Value["id"][0]

			//previous := string(ctx.Request.FormValue("_previous_"))

			formData, title, description := models.GlobalTableList[prefix].GetDataFromDatabaseWithId(prefix, id)

			//url := "/edit/" + prefix + "?id=" + id

			//token := auth.TokenHelper.AddToken()

			tmpl, tmplName := template.Get("adminlte").GetTemplate(true)

			if err != nil {
				fmt.Println(err)
			}

			user := ctx.UserValue["user"].(auth.User)

			buf := new(bytes.Buffer)
			tmpl.ExecuteTemplate(buf, tmplName, types.Page{
				User: user,
				Menu: *menu.GlobalMenu,
				System: types.SystemInfo{
					"0.0.1",
				},
				Panel: types.Panel{
					Content:     template.Get(Config.THEME).Form().SetPrefix(Config.ADMIN_PREFIX).SetContent(formData).GetContent(),
					Description: description,
					Title:       title,
				},
				AssertRootUrl: Config.ADMIN_PREFIX,
			})

			ctx.WriteString(buf.String())
			ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
			ctx.Response.Header.Add("X-PJAX-URL", "/info/user/edit?id="+id)
			return
		}

		if ok, _ = regexp.Match("/new(.*)", []byte(ctx.Path())); ok {
			//prefix := ctx.Request.URL.Query().Get("prefix")

			buffer := new(bytes.Buffer)

			//form, _ := ctx.MultipartForm()

			//id := (*form).Value["id"][0]

			//previous := string(ctx.Request.FormValue("_previous_"))
			//
			//url := "/edit/" + prefix + "?id=" + id
			//
			//token := auth.TokenHelper.AddToken()

			//template.NewPanelPjax(models.GlobalTableList[prefix].Form.FormList, url,
			//	previous, id, models.GlobalTableList[prefix].Form.Title,
			//	models.GlobalTableList[prefix].Form.Description, models.ErrStruct{"hidden", errMsg}, token, buffer)

			ctx.WriteString(buffer.String())
			ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
			ctx.Response.Header.Add("X-PJAX-URL", "/info/user/new")
			return
		}

		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetContentType("application/json")
		ctx.WriteString(`{"code":500, "msg":"` + errMsg + `"}`)
		return
	}
}
