// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package page

import (
	"bytes"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"strings"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(ctx *context.Context, c func() types.Panel) {
	user := auth.Auth(ctx)

	panel := c()

	globalConfig := config.Get()

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(ctx.Headers("X-PJAX") == "true")

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *(menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), "/"+config.Get().PREFIX, "", 1))),
		System: types.SystemInfo{
			"0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: "/" + globalConfig.PREFIX,
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
	})
	ctx.WriteString(buf.String())

}
