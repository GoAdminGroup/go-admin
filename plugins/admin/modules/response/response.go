package response

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
)

func Ok(ctx *context.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
	})
}

func OkWithData(ctx *context.Context, data map[string]interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

func BadRequest(ctx *context.Context, msg string) {
	var user = models.UserModel{}
	var ok bool
	if user, ok = ctx.UserValue["user"].(models.UserModel);!ok {
		user.Id = -1
	}
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"code": 400,
		"msg":  language.GetUser(msg, user.Id),
	})
}

func Alert(ctx *context.Context, config config.Config, desc, title, msg string, conn db.Connection) {
	user := auth.Auth(ctx)

	alert := template.Get(config.Theme).Alert().
		SetTitle(constant.DefaultErrorMsg).
		SetTheme("warning").
		SetContent(template2.HTML(msg)).
		GetContent()

	tmpl, tmplName := template.Get(config.Theme).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     alert,
		Description: desc,
		Title:       title,
	}, config, menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())))
	ctx.HTML(http.StatusOK, buf.String())
}

func Error(ctx *context.Context, msg string) {
	var user = models.UserModel{}
	var ok bool
	if user, ok = ctx.UserValue["user"].(models.UserModel);!ok {
		user.Id = -1
	}
	ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		"code": 500,
		"msg":  language.GetUser(msg, user.Id),
	})
}
