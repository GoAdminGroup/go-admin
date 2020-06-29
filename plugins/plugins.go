// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package plugins

import (
	"bytes"
	"errors"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/modules/ui"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
	"plugin"
	"time"
)

// Plugin as one of the key components of goAdmin has three
// methods. GetRequest return all the path registered in the
// plugin. GetHandler according the url and method return the
// corresponding handler. InitPlugin init the plugin which do
// something like init the database and set the config and register
// the routes. The Plugin must implement the three methods.
type Plugin interface {
	GetHandler() context.HandlerMap
	InitPlugin(services service.List)
	Name() string
	Prefix() string
	GetInfo() Info
	GetInstallationPage() (skip bool, gen table.Generator)
	IsInstalled() bool
	CheckUpdate() (update bool, version string)
	Translate(word string) string
	Uninstall() error
	Upgrade() error
}

type Info struct {
	Title       string    `json:"title" yaml:"title" ini:"title"`
	Description string    `json:"description" yaml:"description" ini:"description"`
	Version     string    `json:"version" yaml:"version" ini:"version"`
	Author      string    `json:"author" yaml:"author" ini:"author"`
	Banners     []string  `json:"banners" yaml:"banners" ini:"banners"`
	Url         string    `json:"url" yaml:"url" ini:"url"`
	Cover       string    `json:"cover" yaml:"cover" ini:"cover"`
	Website     string    `json:"website" yaml:"website" ini:"website"`
	Agreement   string    `json:"agreement" yaml:"agreement" ini:"agreement"`
	CreatedAt   time.Time `json:"created_at" yaml:"created_at" ini:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" yaml:"updated_at" ini:"updated_at"`
}

type Base struct {
	App       *context.App
	Services  service.List
	Conn      db.Connection
	UI        *ui.Service
	PlugName  string
	URLPrefix string
}

func (b *Base) GetHandler() context.HandlerMap                        { return b.App.Handlers }
func (b *Base) Name() string                                          { return b.PlugName }
func (b *Base) GetInfo() Info                                         { return Info{} }
func (b *Base) Translate(word string) string                          { return word }
func (b *Base) Prefix() string                                        { return b.URLPrefix }
func (b *Base) IsInstalled() bool                                     { return false }
func (b *Base) Uninstall() error                                      { return nil }
func (b *Base) Upgrade() error                                        { return nil }
func (b *Base) CheckUpdate() (update bool, version string)            { return false, "" }
func (b *Base) GetInstallationPage() (skip bool, gen table.Generator) { return true, nil }

func (b *Base) InitBase(srv service.List) {
	b.Services = srv
	b.Conn = db.GetConnection(b.Services)
	b.UI = ui.GetService(b.Services)
}

func (b *Base) ExecuteTmpl(ctx *context.Context, panel types.Panel, animation ...bool) *bytes.Buffer {
	return Execute(ctx, b.Conn, *b.UI.NavButtons, auth.Auth(ctx), panel, animation...)
}

func (b *Base) HTML(ctx *context.Context, panel types.Panel, animation ...bool) {
	buf := b.ExecuteTmpl(ctx, panel, animation...)
	ctx.HTMLByte(http.StatusOK, buf.Bytes())
}

func (b *Base) HTMLFile(ctx *context.Context, path string, data map[string]interface{}, animation ...bool) {

	buf := new(bytes.Buffer)
	var panel types.Panel

	t, err := template2.ParseFiles(path)
	if err != nil {
		panel = template.WarningPanel(err.Error()).GetContent(config.IsProductionEnvironment())
	} else {
		if err := t.Execute(buf, data); err != nil {
			panel = template.WarningPanel(err.Error()).GetContent(config.IsProductionEnvironment())
		} else {
			panel = types.Panel{
				Content: template.HTML(buf.String()),
			}
		}
	}

	b.HTML(ctx, panel, animation...)
}

func (b *Base) HTMLFiles(ctx *context.Context, data map[string]interface{}, files []string, animation ...bool) {
	buf := new(bytes.Buffer)
	var panel types.Panel

	t, err := template2.ParseFiles(files...)
	if err != nil {
		panel = template.WarningPanel(err.Error()).GetContent(config.IsProductionEnvironment())
	} else {
		if err := t.Execute(buf, data); err != nil {
			panel = template.WarningPanel(err.Error()).GetContent(config.IsProductionEnvironment())
		} else {
			panel = types.Panel{
				Content: template.HTML(buf.String()),
			}
		}
	}

	b.HTML(ctx, panel, animation...)
}

func LoadFromPlugin(mod string) Plugin {

	plug, err := plugin.Open(mod)
	if err != nil {
		logger.Error("LoadFromPlugin err", err)
		panic(err)
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		logger.Error("LoadFromPlugin err", err)
		panic(err)
	}

	var p Plugin
	p, ok := symPlugin.(Plugin)
	if !ok {
		logger.Error("LoadFromPlugin err: unexpected type from module symbol")
		panic(errors.New("LoadFromPlugin err: unexpected type from module symbol"))
	}

	return p
}

// GetHandler is a help method for Plugin GetHandler.
func GetHandler(app *context.App) context.HandlerMap { return app.Handlers }

func Execute(ctx *context.Context, conn db.Connection, navButtons types.Buttons, user models.UserModel,
	panel types.Panel, animation ...bool) *bytes.Buffer {
	tmpl, tmplName := template.Get(config.GetTheme()).GetTemplate(ctx.IsPjax())

	return template.Execute(template.ExecuteParam{
		User:       user,
		TmplName:   tmplName,
		Tmpl:       tmpl,
		Panel:      panel,
		Config:     *config.Get(),
		Menu:       menu.GetGlobalMenu(user, conn).SetActiveClass(config.URLRemovePrefix(ctx.Path())),
		Animation:  len(animation) > 0 && animation[0] || len(animation) == 0,
		Buttons:    navButtons.CheckPermission(user),
		NoCompress: len(animation) > 1 && animation[1],
	})
}

type Plugins []Plugin

func (pp Plugins) Add(p Plugin) Plugins {
	if !pp.Exist(p) {
		pp = append(pp, p)
	}
	return pp
}

func (pp Plugins) Exist(p Plugin) bool {
	for _, v := range pp {
		if v.Name() == p.Name() {
			return true
		}
	}
	return false
}

var pluginList = make(Plugins, 0)

func Add(p Plugin) {
	pluginList = pluginList.Add(p)
}

func Get() Plugins {
	return pluginList
}
