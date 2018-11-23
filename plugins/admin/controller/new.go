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

	menu.GlobalMenu.SetActiveClass(ctx.Path())
	prefix := ctx.Query("prefix")
	params := models.GetParam(ctx.Request.URL.Query())

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(tmpl, tmplName, auth.Auth(ctx), types.Panel{
		Content: template.Get(Config.THEME).Form().
			SetPrefix(Config.PREFIX).
			SetContent(models.GetNewFormList(models.TableList[prefix].Form.FormList)).
			SetUrl(Config.PREFIX + "/new/" + prefix).
			SetToken(auth.TokenHelper.AddToken()).
			SetTitle("New").
			SetInfoUrl(Config.PREFIX + "/info/" + prefix + params.GetRouteParamStr()).
			GetContent(),
		Description: models.TableList[prefix].Form.Description,
		Title:       models.TableList[prefix].Form.Title,
	}, Config)
	ctx.Html(http.StatusOK, buf.String())
}

// 新建数据
func NewForm(ctx *context.Context) {

	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "新增失败",
		})
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
		models.TableList[prefix].InsertDataFromDatabase((*form).Value)
	}

	models.RefreshTableList()

	previous := ctx.FormValue("_previous_")
	prevUrlArr := strings.Split(previous, "?")
	params := models.GetParamFromUrl(previous)

	panelInfo := models.TableList[prefix].GetDataFromDatabase(prevUrlArr[0], params)

	menu.GlobalMenu.SetActiveClass(previous)

	editUrl := Config.PREFIX + "/info/" + prefix + "/edit" + params.GetRouteParamStr()
	newUrl := Config.PREFIX + "/info/" + prefix + "/new" + params.GetRouteParamStr()
	deleteUrl := Config.PREFIX + "/delete/" + prefix

	dataTable := template.Get(Config.THEME).
		DataTable().
		SetInfoList(panelInfo.InfoList).
		SetThead(panelInfo.Thead).
		SetEditUrl(editUrl).
		SetNewUrl(newUrl).
		SetDeleteUrl(deleteUrl)

	table := dataTable.GetContent()

	box := template.Get(Config.THEME).Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(panelInfo.Paginator.GetContent()).
		GetContent()

	user := ctx.UserValue["user"].(auth.User)

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(true)
	buffer := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, Config)

	ctx.Html(http.StatusOK, buffer.String())
	ctx.AddHeader("X-PJAX-URL", previous)
}
