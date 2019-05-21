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

// 显示表单
func ShowForm(ctx *context.Context) {

	prefix := ctx.Query("prefix")
	if !models.TableList[prefix].Editable {
		ctx.Html(http.StatusNotFound, "page not found")
		return
	}

	formData, title, description := models.TableList[prefix].
		GetDataFromDatabaseWithId(ctx.Query("id"))

	params := models.GetParam(ctx.Request.URL.Query())

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content: template.Get(Config.THEME).Form().
			SetContent(formData).
			SetPrefix(Config.PREFIX).
			SetUrl(Config.PREFIX + "/edit/" + prefix).
			SetToken(auth.TokenHelper.AddToken()).
			SetInfoUrl(Config.PREFIX + "/info/" + prefix + params.GetRouteParamStr()).
			GetContent(),
		Description: description,
		Title:       title,
	}, Config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1)))
	ctx.Html(http.StatusOK, buf.String())
}

// 编辑数据
func EditForm(ctx *context.Context) {
	prefix := ctx.Query("prefix")
	if !models.TableList[prefix].Editable {
		ctx.Html(http.StatusNotFound, "page not found")
		return
	}
	token := ctx.FormValue("_t")

	if !auth.TokenHelper.CheckToken(token) {
		ctx.Json(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "编辑失败",
		})
		return
	}

	form := ctx.Request.MultipartForm

	menu.GlobalMenu.SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1))

	// 处理上传文件，目前仅仅支持传本地
	if len((*form).File) > 0 {
		_, _ = file.GetFileEngine("local").Upload(form)
	}

	if prefix == "manager" { // 管理员管理编辑
		EditManager((*form).Value)
	} else if prefix == "roles" { // 管理员角色管理编辑
		EditRole((*form).Value)
	} else {
		val := (*form).Value
		for _, f := range models.TableList[prefix].Form.FormList {
			if f.Editable {
				continue
			}
			if len(val[f.Field]) > 0 && f.Field != "id"{
				ctx.Json(http.StatusBadRequest, map[string]interface{}{
					"code": 400,
					"msg":  "字段[" + f.Field + "]不可编辑",
				})
				return
			}
		}
		models.TableList[prefix].UpdateDataFromDatabase((*form).Value)
	}

	models.RefreshTableList()

	previous := ctx.FormValue("_previous_")
	prevUrlArr := strings.Split(previous, "?")
	params := models.GetParamFromUrl(previous)

	previous = Config.PREFIX + "/info/" + prefix + params.GetRouteParamStr()
	editUrl := Config.PREFIX + "/info/" + prefix + "/edit" + params.GetRouteParamStr()
	newUrl := Config.PREFIX + "/info/" + prefix + "/new" + params.GetRouteParamStr()
	deleteUrl := Config.PREFIX + "/delete/" + prefix

	panelInfo := models.TableList[prefix].GetDataFromDatabase(prevUrlArr[0], params)

	dataTable := template.Get(Config.THEME).
		DataTable().
		SetInfoList(panelInfo.InfoList).
		SetThead(panelInfo.Thead).
		SetNewUrl(newUrl)

	if panelInfo.Editable {
		dataTable.SetEditUrl(editUrl)
	}
	if panelInfo.Deletable {
		dataTable.SetDeleteUrl(deleteUrl)
	}

	table := dataTable.GetContent()

	box := template.Get(Config.THEME).Box().
		SetBody(table).
		SetHeader(dataTable.GetDataTableHeader()).
		WithHeadBorder(false).
		SetFooter(panelInfo.Paginator.GetContent()).
		GetContent()

	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(true)
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     box,
		Description: panelInfo.Description,
		Title:       panelInfo.Title,
	}, Config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(previous, Config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buf.String())
	ctx.AddHeader("X-PJAX-URL", previous)
}
