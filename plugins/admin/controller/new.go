package controller

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/plugins/admin/modules/file"
	"strings"
	"github.com/chenhg5/go-admin/template/adminlte/components"
	"github.com/chenhg5/go-admin/context"
	"net/http"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template"
)

// 显示新建表单
func ShowNewForm(ctx *context.Context) {
	defer GlobalDeferHandler(ctx)

	user := ctx.UserValue["user"].(auth.User)

	prefix := ctx.Request.URL.Query().Get("prefix")

	tmpl, tmplName := components.GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

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
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel: types.Panel{
			Content:     template.Get(Config.THEME).Form().
				SetPrefix(Config.ADMIN_PREFIX).
				SetContent(models.GetNewFormList(models.GlobalTableList[prefix].Form.FormList)).
				SetUrl(Config.ADMIN_PREFIX + "/new/" + prefix).
				SetToken(auth.TokenHelper.AddToken()).
				SetInfoUrl(Config.ADMIN_PREFIX + "/info/" + prefix + "?page=" + string(page) + "&pageSize=" + string(pageSize) + "&sort=" + string(sortField) + "&sort_type=" + string(sortType)).
				GetContent(),
			Description: models.GlobalTableList[prefix].Form.Description,
			Title:       models.GlobalTableList[prefix].Form.Title,
		},
		AssertRootUrl: Config.ADMIN_PREFIX,
	})
	ctx.Write(http.StatusOK, map[string]string{}, buf.String())
}

// 新建数据
func NewForm(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	token := string(ctx.Request.FormValue("_t"))

	if !auth.TokenHelper.CheckToken(token) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString(`{"code":400, "msg":"新增失败"}`)
		return
	}

	prefix := ctx.Request.URL.Query().Get("prefix")

	form := ctx.Request.MultipartForm

	// 处理上传文件，目前仅仅支持传本地
	if len((*form).File) > 0 {
		file.GetFileEngine("local").Upload(form)
	}

	if prefix == "manager" { // 管理员管理新建
		NewManager((*form).Value)
	} else if prefix == "roles" { // 管理员角色管理新建
		NewRole((*form).Value)
	} else {
		models.GlobalTableList[prefix].InsertDataFromDatabase(prefix, (*form).Value)
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

	buffer := new(bytes.Buffer)

	editUrl := Config.ADMIN_PREFIX + "/info/" + prefix + "/edit?page=" + string(page) + "&pageSize=" + string(pageSize)

	tmpl, tmplName := components.GetTemplate(true)

	user := ctx.UserValue["user"].(auth.User)

	tmpl.ExecuteTemplate(buffer, tmplName, types.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel: types.Panel{
			Content:     template.Get(Config.THEME).DataTable().SetInfoList(infoList).SetThead(thead).SetEditUrl(editUrl).GetContent(),
			Description: description,
			Title:       title,
		},
		AssertRootUrl: Config.ADMIN_PREFIX,
	})

	ctx.WriteString(buffer.String())
	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
	ctx.Response.Header.Add("X-PJAX-URL", previous)
}
