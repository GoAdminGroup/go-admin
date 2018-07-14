package controller

import (
	"github.com/valyala/fasthttp"
	"goAdmin/transform"
)

func DeleteData(ctx *fasthttp.RequestCtx) {

	defer GlobalDeferHandler(ctx)

	prefix := ctx.UserValue("prefix").(string)

	id := string(ctx.FormValue("id")[:])

	transform.DeleteDataFromDatabase(prefix, id)

	// TODO: 增加反馈

	ctx.WriteString(`{"code":200, "msg":"删除成功"`)
	return
}
