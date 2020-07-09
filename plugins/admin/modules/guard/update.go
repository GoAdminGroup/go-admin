package guard

import (
	"net/http"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

type UpdateParam struct {
	Panel  table.Table
	Prefix string
	Value  form.Values
}

func (g *Guard) Update(ctx *context.Context) {
	panel, prefix := g.table(ctx)

	pname := panel.GetPrimaryKey().Name

	id := ctx.FormValue("pk")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "wrong " + pname,
		})
		ctx.Abort()
		return
	}

	var f = make(form.Values)
	f.Add(form.PostIsSingleUpdateKey, "1")
	f.Add(pname, id)
	f.Add(ctx.FormValue("name"), ctx.FormValue("value"))

	ctx.SetUserValue(updateParamKey, &UpdateParam{
		Panel:  panel,
		Prefix: prefix,
		Value:  f,
	})
	ctx.Next()
}

func GetUpdateParam(ctx *context.Context) *UpdateParam {
	return ctx.UserValue[updateParamKey].(*UpdateParam)
}
