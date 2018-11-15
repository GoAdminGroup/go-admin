package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/plugins/admin/modules/file"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"net/http"
	"strings"
)

// 显示新建表单
func ShowNewForm(ctx *context.Context) {

	user := ctx.UserValue["user"].(auth.User)

	prefix := ctx.Query("prefix")

	path := ctx.Path()
	menu.GlobalMenu.SetActiveClass(path)

	page := ctx.QueryDefault("page", "1")
	pageSize := ctx.QueryDefault("pageSize", "10")
	sortField := ctx.QueryDefault("sort", "id")
	sortType := ctx.QueryDefault("sort_type", "desc")

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content: template.Get(Config.THEME).Form().
			SetPrefix(Config.PREFIX).
			SetContent(models.GetNewFormList(models.TableList[prefix].Form.FormList)).
			SetUrl(Config.PREFIX + "/new/" + prefix).
			SetToken(auth.TokenHelper.AddToken()).
			SetTitle("New").
			SetInfoUrl(Config.PREFIX + "/info/" + prefix + GetRouteParameterString(page, pageSize, sortType, sortField)).
			GetContent(),
		Description: models.TableList[prefix].Form.Description,
		Title:       models.TableList[prefix].Form.Title,
	}, Config)
	ctx.WriteString(buf.String())
}

// 新建数据
func NewForm(ctx *context.Context) {

	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteString(`{"code":400, "msg":"新增失败"}`)
		return
	}

	prefix := ctx.Query("prefix")

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
		models.TableList[prefix].InsertDataFromDatabase(prefix, (*form).Value)
	}

	models.RefreshTableList()

	previous := ctx.FormValue("_previous_")

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

	thead, infoList, paginator, title, description := models.TableList[prefix].GetDataFromDatabase(map[string]string{
		"page":      page,
		"path":      prevUrlArr[0],
		"sortField": sort,
		"sortType":  sortType,
		"prefix":    prefix,
		"pageSize":  pageSize,
	})

	menu.GlobalMenu.SetActiveClass(previous)

	editUrl := Config.PREFIX + "/info/" + prefix + "/edit" + GetRouteParameterString(page, pageSize, sortType, sort)
	newUrl := Config.PREFIX + "/info/" + prefix + "/new" + GetRouteParameterString(page, pageSize, sortType, sort)
	deleteUrl := Config.PREFIX + "/delete/" + prefix

	dataTable := template.Get(Config.THEME).
		DataTable().
		SetInfoList(infoList).
		SetThead(thead).
		SetEditUrl(editUrl).
		SetNewUrl(newUrl).
		SetDeleteUrl(deleteUrl)

	table := dataTable.GetContent()

	box := template.Get(Config.THEME).Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(paginator.GetContent()).
		GetContent()

	user := ctx.UserValue["user"].(auth.User)

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(true)
	buffer := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: description,
		Title:       title,
	}, Config)

	ctx.WriteString(buffer.String())
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader("X-PJAX-URL", previous)
}
