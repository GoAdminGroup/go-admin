package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"goAdmin/modules/auth"
	"goAdmin/plugins/admin/models"
	"goAdmin/plugins/admin/modules/file"
	"strings"
	"goAdmin/template/adminlte/components"
	"goAdmin/context"
	"net/http"
	"goAdmin/modules/menu"
)

// 显示表单
func ShowForm(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue["user"].(auth.User)

	prefix := ctx.Request.URL.Query().Get("prefix")

	id := ctx.Request.URL.Query().Get("id")

	formData, title, description := models.GlobalTableList[prefix].GetDataFromDatabaseWithId(prefix, id)

	tmpl := components.GetTemplate(string(ctx.Request.Header.Get("X-PJAX")) == "true")

	path := string(ctx.Path())
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
	tmpl.ExecuteTemplate(buf, "layout", components.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: components.SystemInfo{
			"0.0.1",
		},
		Panel: components.Panel{
			Content:     components.Form().
				SetContent(formData).
				SetPrefix(AssertRootUrl).
				SetUrl(AssertRootUrl + "/edit/" + prefix).
				SetToken(auth.TokenHelper.AddToken()).
				SetInfoUrl(AssertRootUrl + "/info/" + prefix + "?page=" + string(page) + "&pageSize=" + string(pageSize) + "&sort=" + string(sortField) + "&sort_type=" + string(sortType)).
				GetContent(),
			Description: description,
			Title:       title,
		},
		AssertRootUrl: AssertRootUrl,
	})
	ctx.Write(http.StatusOK, map[string]string{}, buf.String())
}

// 编辑数据
func EditForm(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	token := string(ctx.Request.FormValue("_t"))

	if !auth.TokenHelper.CheckToken(token) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString(`{"code":400, "msg":"编辑失败"}`)
		return
	}

	prefix := ctx.Request.URL.Query().Get("prefix")
	user := ctx.UserValue["user"].(auth.User)

	form := ctx.Request.MultipartForm

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

	previous := string(ctx.Request.FormValue("_previous_"))
	prevUrlArr := strings.Split(previous, "?")
	paramArr := strings.Split(prevUrlArr[1], "&")
	page := "1"
	pageSize := "10"
	sort := "id"
	sortType := "desc"

	for i := 0; i < len(paramArr); i++ {
		if strings.Index(paramArr[i], "pageSize") >= 0 {
			pageSize = strings.Split(paramArr[i], "=")[1]
		} else {
			if strings.Index(paramArr[i], "page") >= 0 {
				page = strings.Split(paramArr[i], "=")[1]
			} else if strings.Index(paramArr[i], "sort") >= 0 {
				sort = strings.Split(paramArr[i], "=")[1]
			} else {
				sortType = strings.Split(paramArr[i], "=")[1]
			}
		}
	}

	thead, infoList, _, title, description := models.GlobalTableList[prefix].GetDataFromDatabase(map[string]string{
		"page":      page,
		"path":      prevUrlArr[0],
		"sortField": sort,
		"sortType":  sortType,
		"prefix":    prefix,
		"pageSize":  pageSize,
	})

	menu.GlobalMenu.SetActiveClass(previous)

	editUrl := AssertRootUrl + "/info/" + prefix + "/edit?page=" + string(page) + "&pageSize=" + string(pageSize)

	tmpl := components.GetTemplate(true)

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, "layout", components.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: components.SystemInfo{
			"0.0.1",
		},
		Panel: components.Panel{
			Content:     components.DataTable().SetInfoList(infoList).SetThead(thead).SetEditUrl(editUrl).GetContent(),
			Description: description,
			Title:       title,
		},
		AssertRootUrl: AssertRootUrl,
	})

	ctx.WriteString(buf.String())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
	ctx.Response.Header.Add("X-PJAX-URL", previous)
}
