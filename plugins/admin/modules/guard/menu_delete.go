package guard

import (
	"github.com/digroad/go-admin/context"
	"github.com/digroad/go-admin/modules/errors"
)

type MenuDeleteParam struct {
	Id string
}

func (g *Guard) MenuDelete(ctx *context.Context) {

	id := ctx.Query("id")

	if id == "" {
		alertWithTitleAndDesc(ctx, "Menu", "menu", errors.WrongID, g.conn, g.navBtns)
		ctx.Abort()
		return
	}

	// TODO: check the user permission

	ctx.SetUserValue(deleteMenuParamKey, &MenuDeleteParam{
		Id: id,
	})
	ctx.Next()
}

func GetMenuDeleteParam(ctx *context.Context) *MenuDeleteParam {
	return ctx.UserValue[deleteMenuParamKey].(*MenuDeleteParam)
}
