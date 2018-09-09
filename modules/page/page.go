package page

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/context"
	"bytes"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(theme string, AssertRootUrl string, ctx *context.Context, c func() types.Panel) {
	user := ctx.UserValue["user"].(auth.User)

	panel := c()

	tmpl, tmplName := template.Get(theme).GetTemplate(ctx.Request.Header.Get("X-PJAX") == "true")

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *menu.GlobalMenu,
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: AssertRootUrl,
	})
	ctx.WriteString(buf.String())

}
