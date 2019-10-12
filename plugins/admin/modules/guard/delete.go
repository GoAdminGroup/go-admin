package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

type DeleteParam struct {
	Panel  table.Table
	Id     string
	Prefix string
}

func Delete(ctx *context.Context) {

	prefix := ctx.Query("__prefix")
	panel := table.List[prefix]
	if !panel.GetDeletable() {
		alert(ctx, panel, "operation not allow")
		ctx.Abort()
		return
	}

	id := ctx.FormValue("id")
	if id == "" {
		alert(ctx, panel, "wrong id")
		ctx.Abort()
		return
	}

	ctx.SetUserValue("delete_param", &DeleteParam{
		Panel:  panel,
		Id:     id,
		Prefix: prefix,
	})
	ctx.Next()
}

func GetDeleteParam(ctx *context.Context) *DeleteParam {
	return ctx.UserValue["delete_param"].(*DeleteParam)
}
