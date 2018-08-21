package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/modules/auth"
	"goAdmin/components/menu"
	"goAdmin/plugins/admin/models"
	"goAdmin/modules/file"
	"strings"
	"fmt"
	"goAdmin/components"
	"goAdmin/app"
	"goAdmin/components/adminlte"
)

// 显示表单
func ShowForm(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue("user").(auth.User)

	prefix := "user"

	formData, title, description := models.GlobalTableList[prefix].GetDataFromDatabaseWithId(prefix, "2")

	tmpl := adminlte.GetTemplate(string(ctx.Request.Header.Peek("X-PJAX")) == "true")

	path := string(ctx.Path())
	menu.GlobalMenu.SetActiveClass(path)

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

	tmpl.ExecuteTemplate(ctx.Response.BodyWriter(), "layout", components.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: components.SystemInfo{
			"0.0.1",
		},
		Panel: components.Panel{
			Content:     app.GetComponents().Form().SetContent(formData).GetContent(),
			Description: description,
			Title:       title,
		},
	})
}

// 编辑数据
func EditForm(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	token := string(ctx.FormValue("_t"))

	if !auth.TokenHelper.CheckToken(token) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString(`{"code":400, "msg":"编辑失败"}`)
		return
	}

	prefix := ctx.UserValue("prefix").(string)
	user := ctx.UserValue("user").(auth.User)

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.SetStatusCode(500)
		return
	}

	path := string(ctx.Path())
	menu.GlobalMenu.SetActiveClass(path)

	// 处理上传文件，目前仅仅支持传本地
	if len((*form).File) > 0 {
		file.GetFileEngine("local").Upload(form)
	}

	if prefix == "manager" { // 管理员管理编辑
		EditManager((*form).Value)
	} else if prefix == "roles" { // 管理员角色管理编辑
		EditRole((*form).Value)
	} else {
		models.GlobalTableList[prefix].UpdateDataFromDatabase(prefix, (*form).Value)
	}

	models.RefreshGlobalTableList()

	previous := string(ctx.FormValue("_previous_"))
	prevUrlArr := strings.Split(previous, "?")
	paramArr := strings.Split(prevUrlArr[1], "&")
	page := "1"
	pageSize := "10"
	sort := "id"
	sort_type := "desc"

	for i := 0; i < len(paramArr); i++ {
		if strings.Index(paramArr[i], "pageSize") >= 0 {
			pageSize = strings.Split(paramArr[i], "=")[1]
		} else {
			if strings.Index(paramArr[i], "page") >= 0 {
				page = strings.Split(paramArr[i], "=")[1]
			} else if strings.Index(paramArr[i], "sort") >= 0 {
				sort = strings.Split(paramArr[i], "=")[1]
			} else {
				sort_type = strings.Split(paramArr[i], "=")[1]
			}
		}
	}

	thead, infoList, paginator, title, description := models.GlobalTableList[prefix].GetDataFromDatabase(map[string]string{
		"page":      page,
		"path":      prevUrlArr[0],
		"sortField": sort,
		"sortType":  sort_type,
		"prefix":    prefix,
		"pageSize":  pageSize,
	})

	menu.GlobalMenu.SetActiveClass(previous)
	paginator["success"] = true

	buffer := new(bytes.Buffer)

	editUrl := "/info/" + prefix + "/edit?page=" + string(page) + "&pageSize=" + string(pageSize)

	tmpl := adminlte.GetTemplate(true)

	if err != nil {
		fmt.Println(err)
	}
	tmpl.ExecuteTemplate(ctx.Response.BodyWriter(), "layout", components.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: components.SystemInfo{
			"0.0.1",
		},
		Panel: components.Panel{
			Content:     app.GetComponents().Table().SetInfoList(infoList).SetThead(thead).SetUrl(editUrl).GetContent(),
			Description: description,
			Title:       title,
		},
	})

	ctx.Response.AppendBody(buffer.Bytes())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
	ctx.Response.Header.Add("X-PJAX-URL", previous)
}
