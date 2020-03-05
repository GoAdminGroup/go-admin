package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
)

type MenuDeleteParam struct {
	Id string
}

func (g *Guard) MenuDelete(ctx *context.Context) {

	id := ctx.Query("id")

	if id == "" {
		alertWithTitleAndDesc(ctx, "Menu", "menu", "wrong id", g.conn)
		ctx.Abort()
		return
	}

	// TODO: check the user permission

	ctx.SetUserValue("delete_menu_param", &MenuDeleteParam{
		Id: id,
	})
	ctx.Next()
}

func GetMenuDeleteParam(ctx *context.Context) *MenuDeleteParam {
	return ctx.UserValue["delete_menu_param"].(*MenuDeleteParam)
}
