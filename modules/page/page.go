// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package page

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(ctx *context.Context, user models.UserModel, c func(ctx interface{}) (types.Panel, error)) {

	panel, err := c(ctx)

	globalConfig := config.Get()

	if err != nil {
		logger.Error("SetPageContent", err)
		alert := template.Get(globalConfig.Theme).
			Alert().
			SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
			SetTheme("warning").SetContent(template2.HTML(err.Error())).GetContent()
		panel = types.Panel{
			Content:     alert,
			Description: language.Get("error"),
			Title:       language.Get("error"),
		}
	}

	tmpl, tmplName := template.Get(globalConfig.Theme).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user,
		*(menu.GetGlobalMenu(user).SetActiveClass(globalConfig.UrlRemovePrefix(ctx.Path()))),
		panel, globalConfig))
	if err != nil {
		logger.Error("SetPageContent", err)
	}
	ctx.WriteString(buf.String())
}
