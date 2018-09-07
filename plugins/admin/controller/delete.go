package controller

import (
	"goAdmin/modules/auth"
	"goAdmin/plugins/admin/models"
	"goAdmin/context"
	"net/http"
)

func DeleteData(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	token := string(ctx.Request.FormValue("_t"))

	if !auth.TokenHelper.CheckToken(token) {
		ctx.SetStatusCode(http.StatusBadRequest)
		ctx.WriteString(`{"code":400, "msg":"删除失败"}`)
		return
	}

	prefix := ctx.Request.URL.Query().Get("prefix")

	id := string(ctx.Request.FormValue("id"))

	models.GlobalTableList[prefix].DeleteDataFromDatabase(prefix, id)

	newToken := auth.TokenHelper.AddToken()

	ctx.WriteString(`{"code":200, "msg":"删除成功", "data":"` + newToken + `"}`)
	return
}
