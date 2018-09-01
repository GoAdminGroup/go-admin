package page

import (
	"github.com/chenhg5/go-admin/template/adminlte/components"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/context"
	"bytes"
	"net/http"
	"github.com/chenhg5/go-admin/modules/menu"
)

// 设置页面内容
func SetPageContent(AssertRootUrl string, ctx *context.Context, c func() components.Panel) {
	user := ctx.UserValue["user"].(auth.User)

	panel := c()

	tmpl := components.GetTemplate(string(ctx.Request.Header.Get("X-PJAX")) == "true")

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, "layout", components.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: components.SystemInfo{
			"0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: AssertRootUrl,
	})
	ctx.Write(http.StatusOK, map[string]string{}, buf.String())

}
