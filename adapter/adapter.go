// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package adapter

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
)

// WebFrameWork is a interface which is used as an adapter of
// framework and goAdmin. It must implement two methods. Use registers
// the routes and the corresponding handlers. Content writes the
// response to the corresponding context of framework.
type WebFrameWork interface {
	Use(interface{}, []plugins.Plugin) error
	Content(interface{}, types.GetPanelFn)
}

type BaseAdapter struct{}

func (BaseAdapter) HTMLContentType() string {
	return "text/html; charset=utf-8"
}

func (BaseAdapter) CookieKey() string {
	return auth.DefaultCookieKey
}

func (BaseAdapter) GetContent(cookie, path, method, pjax string, c types.GetPanelFn,
	ctx interface{}) (body []byte, authSuccess bool, hasError error) {

	var (
		user  models.UserModel
		panel types.Panel
		err   error
	)

	user, authSuccess = auth.GetCurUser(cookie)

	if !authSuccess {
		return
	}

	if !auth.CheckPermissions(user, path, method) {
		alert := getErrorAlert("no permission")
		errTitle := language.Get("error")

		panel = types.Panel{
			Content:     alert,
			Description: errTitle,
			Title:       errTitle,
		}
	} else {
		panel, err = c(ctx)
		if err != nil {
			alert := getErrorAlert(err.Error())
			errTitle := language.Get("error")

			panel = types.Panel{
				Content:     alert,
				Description: errTitle,
				Title:       errTitle,
			}
		}
	}

	tmpl, tmplName := template.Default().GetTemplate(pjax == "true")

	buf := new(bytes.Buffer)
	hasError = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user,
		*(menu.GetGlobalMenu(user).SetActiveClass(config.Get().URLRemovePrefix(path))),
		panel, config.Get(), template.GetComponentAssetListsHTML()))
	body = buf.Bytes()
	return
}

func getErrorAlert(msg string) template2.HTML {
	return template.Default().Alert().
		SetTitle(template2.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
		SetTheme("warning").
		SetContent(template2.HTML(msg)).
		GetContent()
}
