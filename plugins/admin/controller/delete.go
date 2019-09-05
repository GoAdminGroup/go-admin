package controller

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/modules/response"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
)

func Delete(ctx *context.Context) {
	prefix := ctx.Query("prefix")
	if !table.List[prefix].GetDeletable() {
		response.PageNotFound(ctx)
		return
	}

	//token := ctx.FormValue("_t")
	//
	//if !auth.TokenHelper.CheckToken(token) {
	//	ctx.SetStatusCode(http.StatusBadRequest)
	//	ctx.WriteString(`{"code":400, "msg":"delete fail"}`)
	//	return
	//}

	table.List[ctx.Query("prefix")].
		DeleteDataFromDatabase(ctx.FormValue("id"))

	newToken := auth.TokenHelper.AddToken()

	response.OkWithData(ctx, map[string]interface{}{
		"token": newToken,
	})
	return
}
