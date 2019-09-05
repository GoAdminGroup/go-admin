package response

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/language"
	"net/http"
)

func Ok(ctx *context.Context) {
	ctx.Json(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
	})
}

func OkWithData(ctx *context.Context, data map[string]interface{}) {
	ctx.Json(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

func BadRequest(ctx *context.Context, msg string) {
	ctx.Json(http.StatusBadRequest, map[string]interface{}{
		"code": 400,
		"msg":  language.Get(msg),
	})
}

func PageNotFound(ctx *context.Context) {
	ctx.Html(http.StatusNotFound, "page not found")
}

func Error(ctx *context.Context, msg string) {
	ctx.Json(http.StatusInternalServerError, map[string]interface{}{
		"code": 500,
		"msg":  language.Get(msg),
	})
}
