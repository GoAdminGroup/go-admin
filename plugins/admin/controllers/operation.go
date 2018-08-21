package controller

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"goAdmin/modules/auth"
	"goAdmin/modules/connections/mysql"
)

func RecordOperationLog(ctx *fasthttp.RequestCtx) {
	if user, ok := ctx.UserValue("user").(auth.User); ok {
		var input []byte
		if form, err := ctx.MultipartForm(); err == nil {
			input, _ = json.Marshal((*form).Value)
		} else {
			input = []byte("[]")
		}

		mysql.Exec("insert into goadmin_operation_log (user_id, path, method, ip, input) values (?, ?, ?, ?, ?)", user.ID, ctx.Path(),
			ctx.Method(), ctx.LocalIP().String(), string(input))
	}
}
