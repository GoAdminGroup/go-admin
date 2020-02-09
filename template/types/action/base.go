package action

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"net/http"
	"strings"
)

type AjaxData map[string]interface{}

func NewAjaxData() AjaxData {
	return AjaxData{"ids": "{%ids}"}
}

func (a AjaxData) Add(m map[string]interface{}) AjaxData {
	for k, v := range m {
		a[k] = v
	}
	return a
}

func (a AjaxData) JSON() string {
	b, _ := json.Marshal(a)
	s := strings.Replace(string(b), `"{%ids}"`, "{%ids}", -1)
	s = strings.Replace(s, `"{%id}"`, "{%id}", -1)
	return s
}

type Handler func(ctx *context.Context) (success bool, data, msg string)

func (h Handler) Wrap() context.Handler {
	return func(ctx *context.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(err)
				ctx.JSON(http.StatusOK, map[string]interface{}{
					"code": 500,
					"data": "",
					"msg":  "error",
				})
			}
		}()

		code := 0
		s, d, m := h(ctx)

		if !s {
			code = 500
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": code,
			"data": d,
			"msg":  m,
		})
	}
}
