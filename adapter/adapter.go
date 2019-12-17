// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package adapter

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
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
	SetConnection(db.Connection)
	GetConnection() db.Connection
	SetContext(ctx interface{}) WebFrameWork
	GetCookie() (string, error)
	Path() string
	Method() string
	PjaxHeader() string
	Redirect()
	SetContentType()
	Write(body []byte)
	CookieKey() string
	HTMLContentType() string
	Name() string
	User(ci interface{}) (models.UserModel, bool)
	SetApp(app interface{}) error
	AddHandler(method, path string, plug plugins.Plugin)
}

type BaseAdapter struct {
	db db.Connection
}

func (base *BaseAdapter) SetConnection(conn db.Connection) {
	base.db = conn
}

func (base *BaseAdapter) GetConnection() db.Connection {
	return base.db
}

func (base *BaseAdapter) HTMLContentType() string {
	return "text/html; charset=utf-8"
}

func (base *BaseAdapter) CookieKey() string {
	return auth.DefaultCookieKey
}

func (base *BaseAdapter) GetUser(ci interface{}, wf WebFrameWork) (models.UserModel, bool) {
	cookie, err := wf.SetContext(ci).GetCookie()

	if err != nil {
		return models.UserModel{}, false
	}

	user, exist := auth.GetCurUser(cookie, wf.GetConnection())
	return user.ReleaseConn(), exist
}

func (base *BaseAdapter) GetUse(router interface{}, plugin []plugins.Plugin, wf WebFrameWork) error {
	if err := wf.SetApp(router); err != nil {
		return err
	}

	for _, plug := range plugin {
		var plugCopy = plug
		for _, req := range plug.GetRequest() {
			wf.AddHandler(req.Method, req.URL, plugCopy)
		}
	}

	return nil
}

func (base *BaseAdapter) GetContent(ctx interface{}, getPanelFn types.GetPanelFn, wf WebFrameWork) {

	newBase := wf.SetContext(ctx)

	cookie, hasError := newBase.GetCookie()

	if hasError != nil || cookie == "" {
		newBase.Redirect()
		return
	}

	user, authSuccess := auth.GetCurUser(cookie, wf.GetConnection())

	if !authSuccess {
		newBase.Redirect()
		return
	}

	var (
		panel types.Panel
		err   error
	)

	if !auth.CheckPermissions(user, newBase.Path(), newBase.Method()) {
		alert := getErrorAlert("no permission")
		errTitle := language.Get("error")

		panel = types.Panel{
			Content:     alert,
			Description: errTitle,
			Title:       errTitle,
		}
	} else {
		panel, err = getPanelFn(ctx)
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

	tmpl, tmplName := template.Default().GetTemplate(newBase.PjaxHeader() == "true")

	buf := new(bytes.Buffer)
	hasError = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user,
		*(menu.GetGlobalMenu(user, wf.GetConnection()).SetActiveClass(config.Get().URLRemovePrefix(newBase.Path()))),
		panel, config.Get(), template.GetComponentAssetListsHTML()))

	if hasError != nil {
		logger.Error(fmt.Sprintf("error: %s adapter content, ", newBase.Name()), err)
	}

	newBase.SetContentType()
	newBase.Write(buf.Bytes())
}

func getErrorAlert(msg string) template2.HTML {
	return template.Default().Alert().
		SetTitle(template.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
		SetTheme("warning").
		SetContent(template.HTML(msg)).
		GetContent()
}
