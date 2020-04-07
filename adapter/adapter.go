// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package adapter

import (
	"bytes"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"net/url"
)

// WebFrameWork is an interface which is used as an adapter of
// framework and goAdmin. It must implement two methods. Use registers
// the routes and the corresponding handlers. Content writes the
// response to the corresponding context of framework.
type WebFrameWork interface {
	Use(interface{}, []plugins.Plugin) error
	Content(interface{}, types.GetPanelFn, ...types.Button)
	SetConnection(db.Connection)
	GetConnection() db.Connection
	SetContext(ctx interface{}) WebFrameWork
	GetCookie() (string, error)
	Path() string
	Method() string
	FormParam() url.Values
	IsPjax() bool
	Redirect()
	SetContentType()
	Write(body []byte)
	CookieKey() string
	HTMLContentType() string
	Name() string
	User(ci interface{}) (models.UserModel, bool)
	SetApp(app interface{}) error
	AddHandler(method, path string, handlers context.Handlers)
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
		for path, handlers := range plug.GetHandler() {
			wf.AddHandler(path.Method, path.URL, handlers)
		}
	}

	return nil
}

func (base *BaseAdapter) GetContent(ctx interface{}, getPanelFn types.GetPanelFn, wf WebFrameWork, btns types.Buttons) {

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

	if !auth.CheckPermissions(user, newBase.Path(), newBase.Method(), newBase.FormParam()) {
		panel = template.WarningPanel(errors.NoPermission)
	} else {
		panel, err = getPanelFn(ctx)
		if err != nil {
			panel = template.WarningPanel(err.Error())
		}
	}

	tmpl, tmplName := template.Default().GetTemplate(newBase.IsPjax())

	buf := new(bytes.Buffer)
	hasError = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(types.NewPageParam{
		User:    user,
		Menu:    menu.GetGlobalMenu(user, wf.GetConnection()).SetActiveClass(config.URLRemovePrefix(newBase.Path())),
		Panel:   panel.GetContent(config.IsProductionEnvironment()),
		Assets:  template.GetComponentAssetListsHTML(),
		Buttons: btns.CheckPermission(user),
	}))

	if hasError != nil {
		logger.Error(fmt.Sprintf("error: %s adapter content, ", newBase.Name()), hasError)
	}

	newBase.SetContentType()
	newBase.Write(buf.Bytes())
}
