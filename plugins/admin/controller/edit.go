package controller

import (
	"bytes"
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

// 显示表单
func ShowForm(ctx *context.Context) {

	user := ctx.UserValue["user"].(auth.User)

	prefix := ctx.Query("prefix")

	id := ctx.Query("id")

	formData, title, description := models.TableList[prefix].GetDataFromDatabaseWithId(prefix, id)

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")

	path := ctx.Path()
	menu.GlobalMenu.SetActiveClass(path)

	page := ctx.QueryDefault("page", "1")
	pageSize := ctx.QueryDefault("pageSize", "10")
	sortField := ctx.QueryDefault("sort", "id")
	sortType := ctx.QueryDefault("sort_type", "desc")

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: menu.GetGlobalMenu(user),
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel: types.Panel{
			Content: template.Get(Config.THEME).Form().
				SetContent(formData).
				SetPrefix(Config.PREFIX).
				SetUrl(Config.PREFIX + "/edit/" + prefix).
				SetToken(auth.TokenHelper.AddToken()).
				SetInfoUrl(Config.PREFIX + "/info/" + prefix + "?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType).
				GetContent(),
			Description: description,
			Title:       title,
		},
		AssertRootUrl: Config.PREFIX,
		Title:         Config.TITLE,
		Logo:          Config.LOGO,
		MiniLogo:      Config.MINILOGO,
	})
	ctx.WriteString(buf.String())
}

// 编辑数据
func EditForm(ctx *context.Context) {

	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "编辑失败",
		})
		return
	}

	prefix := ctx.Query("prefix")
	user := ctx.UserValue["user"].(auth.User)

	form := ctx.Request.MultipartForm

	path := ctx.Path()
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
		models.TableList[prefix].UpdateDataFromDatabase(prefix, (*form).Value)
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

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(true)

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

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: menu.GetGlobalMenu(user),
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel: types.Panel{
			Content:     box,
			Description: description,
			Title:       title,
		},
		AssertRootUrl: Config.PREFIX,
		Title:         Config.TITLE,
		Logo:          Config.LOGO,
		MiniLogo:      Config.MINILOGO,
	})

	ctx.WriteString(buf.String())
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader("X-PJAX-URL", previous)
}
