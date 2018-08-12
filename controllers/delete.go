package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/models"
)

func DeleteData(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	prefix := ctx.UserValue("prefix").(string)

	id := string(ctx.FormValue("id"))

	models.GlobalTableList[prefix].DeleteDataFromDatabase(prefix, id)

	ctx.WriteString(`{"code":200, "msg":"删除成功"}`)
	return
}
