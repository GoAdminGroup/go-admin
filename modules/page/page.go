// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package page

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(ctx *context.Context, user models.UserModel, c func(ctx interface{}) (types.Panel, error), conn db.Connection) {

	panel, err := c(ctx)

	globalConfig := config.Get()
	errMsg := language.Get("error")

	if err != nil {
		logger.Error("SetPageContent", err)
		alert := template.Get(globalConfig.Theme).
			Alert().
			SetTitle(icon.Icon(icon.Warning, 1) + template.HTML(errMsg) + `!`).
			SetTheme("warning").SetContent(template2.HTML(err.Error())).GetContent()
		panel = types.Panel{
			Content:     alert,
			Description: errMsg,
			Title:       errMsg,
		}
	}

	tmpl, tmplName := template.Get(globalConfig.Theme).GetTemplate(ctx.Headers(constant.PjaxHeader) == "true")

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user,
		*(menu.GetGlobalMenu(user, conn).SetActiveClass(globalConfig.URLRemovePrefix(ctx.Path()))),
		panel.GetContent(globalConfig.IsProductionEnvironment()), globalConfig, template.GetComponentAssetListsHTML()))
	if err != nil {
		logger.Error("SetPageContent", err)
	}
	ctx.WriteString(buf.String())
}
