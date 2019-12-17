package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
)

type MenuDeleteParam struct {
	Id string
}

func MenuDelete(conn db.Connection) context.Handler {
	return func(ctx *context.Context) {

		id := ctx.Query("id")

		if id == "" {
			alertWithTitleAndDesc(ctx, "Menu", "menu", "wrong id", conn)
			ctx.Abort()
			return
		}

		// TODO: check the user permission

		ctx.SetUserValue("delete_menu_param", &MenuDeleteParam{
			Id: id,
		})
		ctx.Next()
	}
}

func GetMenuDeleteParam(ctx *context.Context) *MenuDeleteParam {
	return ctx.UserValue["delete_menu_param"].(*MenuDeleteParam)
}
