package controller

import (
	"bytes"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/mgutz/ansi"
	"github.com/valyala/fasthttp"
	"goAdmin/auth"
	"goAdmin/models"
	"goAdmin/template"
	"log"
	"regexp"
	"runtime/debug"
	"strconv"
)

// 全局错误处理
func GlobalDeferHandler(ctx *fasthttp.RequestCtx) {

	log.Println("[GoAdmin]",
		ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode())+" ", "white:blue"),
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

		if ok, _ = regexp.Match("/edit(.*)", ctx.Path()); ok {
			prefix := ctx.UserValue("prefix").(string)

			buffer := new(bytes.Buffer)

			form, _ := ctx.MultipartForm()

			id := (*form).Value["id"][0]

			previous := string(ctx.FormValue("_previous_"))

			formData, title, description := models.GlobalTableList[prefix].GetDataFromDatabaseWithId(prefix, id)

			url := "/edit/" + prefix + "?id=" + id

			token := auth.TokenHelper.AddToken()

			template.EditPanelPjax(formData, url, previous, id, title, description, models.ErrStruct{"", errMsg}, token, buffer)

			ctx.Response.AppendBody(buffer.Bytes())
			ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
			ctx.Response.Header.Add("X-PJAX-URL", "/info/user/edit?id="+id)
			return
		}

		if ok, _ = regexp.Match("/new(.*)", ctx.Path()); ok {
			prefix := ctx.UserValue("prefix").(string)

			buffer := new(bytes.Buffer)

			form, _ := ctx.MultipartForm()

			id := (*form).Value["id"][0]

			previous := string(ctx.FormValue("_previous_"))

			url := "/edit/" + prefix + "?id=" + id

			token := auth.TokenHelper.AddToken()

			template.NewPanelPjax(models.GlobalTableList[prefix].Form.FormList, url,
				previous, id, models.GlobalTableList[prefix].Form.Title,
				models.GlobalTableList[prefix].Form.Description, models.ErrStruct{"hidden", errMsg}, token, buffer)

			ctx.Response.AppendBody(buffer.Bytes())
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
