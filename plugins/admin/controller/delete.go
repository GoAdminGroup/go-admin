package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/context"
	"github.com/valyala/fasthttp"
)

func DeleteData(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	token := string(ctx.Request.FormValue("_t"))

	if !auth.TokenHelper.CheckToken(token) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
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
