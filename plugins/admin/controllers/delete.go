package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/modules/auth"
	"goAdmin/plugins/admin/models"
)

func DeleteData(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	token := string(ctx.FormValue("_t"))

	if !auth.TokenHelper.CheckToken(token) {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.WriteString(`{"code":400, "msg":"删除失败"}`)
		return
	}

	prefix := ctx.UserValue("prefix").(string)

	id := string(ctx.FormValue("id"))

	models.GlobalTableList[prefix].DeleteDataFromDatabase(prefix, id)

	newToken := auth.TokenHelper.AddToken()

	ctx.WriteString(`{"code":200, "msg":"删除成功", "data":"` + newToken + `"}`)
	return
}
