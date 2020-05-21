package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

type DeleteParam struct {
	Panel  table.Table
	Id     string
	Prefix string
}

func (g *Guard) Delete(ctx *context.Context) {
	panel, prefix := g.table(ctx)
	if !panel.GetDeletable() {
		alert(ctx, panel, errors.OperationNotAllow, g.conn, g.navBtns)
		ctx.Abort()
		return
	}

	id := ctx.FormValue("id")
	if id == "" {
		alert(ctx, panel, errors.WrongID, g.conn, g.navBtns)
		ctx.Abort()
		return
	}

	ctx.SetUserValue(deleteParamKey, &DeleteParam{
		Panel:  panel,
		Id:     id,
		Prefix: prefix,
	})
	ctx.Next()
}

func GetDeleteParam(ctx *context.Context) *DeleteParam {
	return ctx.UserValue[deleteParamKey].(*DeleteParam)
}
