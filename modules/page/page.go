// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package page

import (
	"bytes"
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/logger"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"github.com/chenhg5/go-admin/plugins/admin/modules/constant"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"strings"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(ctx *context.Context, user models.UserModel, c func() types.Panel) {

	panel := c()

	globalConfig := config.Get()

	tmpl, tmplName := template.Get(globalConfig.THEME).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buf, tmplName, types.Page{
		User: user,
		Menu: *(menu.GetGlobalMenu(user).SetActiveClass(strings.Replace(ctx.Path(), globalConfig.Prefix(), "", 1))),
		System: types.SystemInfo{
			Version: "0.0.1",
		},
		Panel:         panel,
		AssertRootUrl: globalConfig.Prefix(),
		Title:         globalConfig.TITLE,
		Logo:          globalConfig.LOGO,
		MiniLogo:      globalConfig.MINILOGO,
		ColorScheme:   globalConfig.COLORSCHEME,
	})
	if err != nil {
		logger.Error("SetPageContent", err)
	}
	ctx.WriteString(buf.String())
}
