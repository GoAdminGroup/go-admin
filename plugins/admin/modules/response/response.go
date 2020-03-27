package response

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
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
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"code": 400,
		"msg":  language.Get(msg),
	})
}

func Alert(ctx *context.Context, desc, title, msg string, conn db.Connection) {
	user := auth.Auth(ctx)

	tmpl, tmplName := template.Get(config.Get().Theme).GetTemplate(ctx.IsPjax())
	buf := template.Execute(template.ExecuteParam{
		User:     user,
		TmplName: tmplName,
		Tmpl:     tmpl,
		Panel: types.Panel{
			Content:     template.Get(config.Get().Theme).Alert().Warning(msg),
			Description: desc,
			Title:       title,
		},
		Config:    config.Get(),
		Menu:      menu.GetGlobalMenu(user, conn).SetActiveClass(config.Get().URLRemovePrefix(ctx.Path())),
		Animation: true,
	})
	ctx.HTML(http.StatusOK, buf.String())
}

func Error(ctx *context.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		"code": 500,
		"msg":  language.Get(msg),
	})
}
