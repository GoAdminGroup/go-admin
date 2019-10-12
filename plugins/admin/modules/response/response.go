package response

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
)

func Ok(ctx *context.Context) {
	ctx.Json(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
	})
}

func OkWithData(ctx *context.Context, data map[string]interface{}) {
	ctx.Json(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

func BadRequest(ctx *context.Context, msg string) {
	ctx.Json(http.StatusBadRequest, map[string]interface{}{
		"code": 400,
		"msg":  language.Get(msg),
	})
}

func Alert(ctx *context.Context, config config.Config, desc, title, msg string) {
	user := auth.Auth(ctx)

	alert := template.Get(config.Theme).Alert().
		SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
		SetTheme("warning").
		SetContent(template2.HTML(msg)).
		GetContent()

	tmpl, tmplName := template.Get(config.Theme).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     alert,
		Description: desc,
		Title:       title,
	}, config, menu.GetGlobalMenu(user).SetActiveClass(config.UrlRemovePrefix(ctx.Path())))
	ctx.Html(http.StatusOK, buf.String())
}

func Error(ctx *context.Context, msg string) {
	ctx.Json(http.StatusInternalServerError, map[string]interface{}{
		"code": 500,
		"msg":  language.Get(msg),
	})
}
