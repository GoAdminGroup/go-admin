package guard

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"net/http"
)

type UpdateParam struct {
	Panel  table.Table
	Prefix string
	Value  form.Values
}

func Update(ctx *context.Context) {
	prefix := ctx.Query("__prefix")
	panel := table.Get(prefix)

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
	f.Add("__go_admin_single_update", "1")
	f.Add(pname, id)
	f.Add(ctx.FormValue("name"), ctx.FormValue("value"))

	ctx.SetUserValue("update_param", &UpdateParam{
		Panel:  panel,
		Prefix: prefix,
		Value:  f,
	})
	ctx.Next()
}

func GetUpdateParam(ctx *context.Context) *UpdateParam {
	return ctx.UserValue["update_param"].(*UpdateParam)
}
