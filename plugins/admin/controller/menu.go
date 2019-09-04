package controller

import (
	"encoding/json"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/db/dialect"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"strings"
)

// 显示菜单
func ShowMenu(ctx *context.Context) {
	GetMenuInfoPanel(ctx)
	return
}

// 显示编辑菜单
func ShowEditMenu(ctx *context.Context) {

	formData, title, description := models.TableList["menu"].GetDataFromDatabaseWithId(ctx.Query("id"))

	user := auth.Auth(ctx)

	js := `<script>
$('.icon').iconpicker({placement: 'bottomLeft'});
</script>`

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content: template.Get(Config.THEME).Form().
			SetContent(formData).
			SetPrefix(Config.PREFIX).
			SetUrl(Config.PREFIX + "/menu/edit").
			SetToken(auth.TokenHelper.AddToken()).
			SetInfoUrl(Config.PREFIX + "/menu").
			GetContent() + template2.HTML(js),
		Description: description,
		Title:       title,
	}, Config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buf.String())
}

// 删除菜单
func DeleteMenu(ctx *context.Context) {

	_ = db.Table("goadmin_menu").Where("id", "=", ctx.Query("id")).Delete()

	menu.SetGlobalMenu(auth.Auth(ctx))
	ctx.Json(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
	})
}

// 编辑菜单
func EditMenu(ctx *context.Context) {
	id := ctx.FormValue("id")
	title := ctx.FormValue("title")
	parentId := ctx.FormValue("parent_id")
	if parentId == "" {
		parentId = "0"
	}
	icon := ctx.FormValue("icon")
	uri := ctx.FormValue("uri")

	roles := ctx.Request.Form["roles[]"]

	for _, roleId := range roles {
		checkRoleMenu, _ := db.Table("goadmin_role_menu").
			Where("role_id", "=", roleId).
			Where("menu_id", "=", id).
			First()

		if checkRoleMenu == nil {
			_, _ = db.Table("goadmin_role_menu").Insert(dialect.H{
				"menu_id": id,
				"role_id": roleId,
			})
		}
	}

	_, _ = db.Table("goadmin_menu").
		Where("id", "=", id).Update(dialect.H{
		"title":     title,
		"parent_id": parentId,
		"icon":      icon,
		"uri":       uri,
	})

	menu.SetGlobalMenu(auth.Auth(ctx))

	GetMenuInfoPanel(ctx)
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader("X-PJAX-URL", Config.PREFIX+"/menu")
}

// 新建菜单
func NewMenu(ctx *context.Context) {

	title := ctx.FormValue("title")
	parentId := ctx.FormValue("parent_id")
	if parentId == "" {
		parentId = "0"
	}
	icon := ctx.FormValue("icon")
	uri := ctx.FormValue("uri")

	user := auth.Auth(ctx)

	id, _ := db.Table("goadmin_menu").Insert(dialect.H{
		"title":     title,
		"parent_id": parentId,
		"icon":      icon,
		"uri":       uri,
		"order":     (menu.GetGlobalMenu(user)).MaxOrder + 1,
	})

	roles := ctx.Request.Form["roles[]"]

	for _, roleId := range roles {
		_, _ = db.Table("goadmin_role_menu").Insert(dialect.H{
			"menu_id": id,
			"role_id": roleId,
		})
	}

	globalMenu := menu.GetGlobalMenu(user)
	globalMenu.SexMaxOrder(globalMenu.MaxOrder + 1)
	menu.SetGlobalMenu(user)

	GetMenuInfoPanel(ctx)
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader("X-PJAX-URL", Config.PREFIX+"/menu")
}

// 修改菜单顺序
func MenuOrder(ctx *context.Context) {

	var data []map[string]interface{}
	_ = json.Unmarshal([]byte(ctx.FormValue("_order")), &data)

	count := 1
	for _, v := range data {
		if child, ok := v["children"]; ok {
			_, _ = db.Table("goadmin_menu").
				Where("id", "=", v["id"]).Update(dialect.H{
				"order":     count,
				"parent_id": 0,
			})

			for _, v2 := range child.([]interface{}) {
				_, _ = db.Table("goadmin_menu").
					Where("id", "=", v2.(map[string]interface{})["id"]).Update(dialect.H{
					"order":     count,
					"parent_id": v["id"],
				})
				count++
			}
		} else {
			_, _ = db.Table("goadmin_menu").
				Where("id", "=", v["id"]).Update(dialect.H{
				"order":     count,
				"parent_id": 0,
			})
			count++
		}
	}
	menu.SetGlobalMenu(auth.Auth(ctx))

	ctx.Json(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
	})
	return
}

func GetMenuInfoPanel(ctx *context.Context) {
	user := auth.Auth(ctx)

	menu.GlobalMenu.SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1))

	editUrl := Config.PREFIX + "/menu/edit/show"
	deleteUrl := Config.PREFIX + "/menu/delete"
	orderUrl := Config.PREFIX + "/menu/order"

	tree := template.Get(Config.THEME).Tree().
		SetTree((menu.GetGlobalMenu(user)).GlobalMenuList).
		SetEditUrl(editUrl).
		SetDeleteUrl(deleteUrl).
		SetOrderUrl(orderUrl).
		GetContent()

	header := template.Get(Config.THEME).Tree().GetTreeHeader()
	box := template.Get(Config.THEME).Box().SetHeader(header).SetBody(tree).GetContent()
	col1 := template.Get(Config.THEME).Col().SetSize(map[string]string{"md": "6"}).SetContent(box).GetContent()

	newForm := template.Get(Config.THEME).Form().
		SetPrefix(Config.PREFIX).
		SetUrl(Config.PREFIX + "/menu/new").
		SetInfoUrl(Config.PREFIX + "/menu").
		SetTitle("New").
		SetContent(models.GetNewFormList(models.TableList["menu"].GetForm().FormList)).
		GetContent()

	col2 := template.Get(Config.THEME).Col().SetSize(map[string]string{"md": "6"}).SetContent(newForm).GetContent()

	row := template.Get(Config.THEME).Row().SetContent(col1 + col2).GetContent()

	menu.GlobalMenu.SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1))

	tmpl, tmplName := template.Get(Config.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")
	buf := template.Excecute(tmpl, tmplName, user, types.Panel{
		Content:     row,
		Description: "Menus Manage",
		Title:       "Menus Manage",
	}, Config, menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), Config.PREFIX, "", 1)))

	ctx.Html(http.StatusOK, buf.String())
}
