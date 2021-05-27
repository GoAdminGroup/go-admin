// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package plugins

import (
	"bytes"
	"encoding/json"
	"errors"
	template2 "html/template"
	"net/http"
	"plugin"
	"time"

	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types/action"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/remote_server"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/modules/ui"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
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
	GetGenerators() table.GeneratorList
	Name() string
	Prefix() string
	GetInfo() Info
	GetIndexURL() string
	GetSettingPage() table.Generator
	IsInstalled() bool
	Uninstall() error
	Upgrade() error
}

type Info struct {
	Title            string    `json:"title" yaml:"title" ini:"title"`
	Description      string    `json:"description" yaml:"description" ini:"description"`
	OldVersion       string    `json:"old_version" yaml:"old_version" ini:"old_version"`
	Version          string    `json:"version" yaml:"version" ini:"version"`
	Author           string    `json:"author" yaml:"author" ini:"author"`
	Banners          []string  `json:"banners" yaml:"banners" ini:"banners"`
	Url              string    `json:"url" yaml:"url" ini:"url"`
	Cover            string    `json:"cover" yaml:"cover" ini:"cover"`
	MiniCover        string    `json:"mini_cover" yaml:"mini_cover" ini:"mini_cover"`
	Website          string    `json:"website" yaml:"website" ini:"website"`
	Agreement        string    `json:"agreement" yaml:"agreement" ini:"agreement"`
	CreateDate       time.Time `json:"create_date" yaml:"create_date" ini:"create_date"`
	UpdateDate       time.Time `json:"update_date" yaml:"update_date" ini:"update_date"`
	ModulePath       string    `json:"module_path" yaml:"module_path" ini:"module_path"`
	Name             string    `json:"name" yaml:"name" ini:"name"`
	Uuid             string    `json:"uuid" yaml:"uuid" ini:"uuid"`
	Downloaded       bool      `json:"downloaded" yaml:"downloaded" ini:"downloaded"`
	ExtraDownloadUrl string    `json:"extra_download_url" yaml:"extra_download_url" ini:"extra_download_url"`
	Price            []string  `json:"price" yaml:"price" ini:"price"`
	GoodUUIDs        []string  `json:"good_uuids" yaml:"good_uuids" ini:"good_uuids"`
	GoodNum          int64     `json:"good_num" yaml:"good_num" ini:"good_num"`
	CommentNum       int64     `json:"comment_num" yaml:"comment_num" ini:"comment_num"`
	Order            int64     `json:"order" yaml:"order" ini:"order"`
	Features         string    `json:"features" yaml:"features" ini:"features"`
	Questions        []string  `json:"questions" yaml:"questions" ini:"questions"`
	HasBought        bool      `json:"has_bought" yaml:"has_bought" ini:"has_bought"`
	CanUpdate        bool      `json:"can_update" yaml:"can_update" ini:"can_update"`
	Legal            bool      `json:"legal" yaml:"legal" ini:"legal"`
	SkipInstallation bool      `json:"skip_installation" yaml:"skip_installation" ini:"skip_installation"`
}

func (i Info) IsFree() bool {
	return len(i.Price) == 0
}

type Base struct {
	App       *context.App
	Services  service.List
	Conn      db.Connection
	UI        *ui.Service
	PlugName  string
	URLPrefix string
	Info      Info
}

func (b *Base) InitPlugin(services service.List)   {}
func (b *Base) GetGenerators() table.GeneratorList { return make(table.GeneratorList) }
func (b *Base) GetHandler() context.HandlerMap     { return b.App.Handlers }
func (b *Base) Name() string                       { return b.PlugName }
func (b *Base) GetInfo() Info                      { return b.Info }
func (b *Base) Prefix() string                     { return b.URLPrefix }
func (b *Base) IsInstalled() bool                  { return false }
func (b *Base) Uninstall() error                   { return nil }
func (b *Base) Upgrade() error                     { return nil }
func (b *Base) GetIndexURL() string                { return "" }
func (b *Base) GetSettingPage() table.Generator    { return nil }

func (b *Base) InitBase(srv service.List, prefix string) {
	b.Services = srv
	b.Conn = db.GetConnection(b.Services)
	b.UI = ui.GetService(b.Services)
	b.URLPrefix = prefix
}

func (b *Base) SetInfo(info Info) {
	b.Info = info
}

func (b *Base) Title() string {
	return language.GetWithScope(b.Info.Title, b.Name())
}

func (b *Base) ExecuteTmpl(ctx *context.Context, panel types.Panel, options template.ExecuteOptions) *bytes.Buffer {
	return Execute(ctx, b.Conn, *b.UI.NavButtons, auth.Auth(ctx), panel, options)
}

func (b *Base) ExecuteTmplWithNavButtons(ctx *context.Context, panel types.Panel, btns types.Buttons,
	options template.ExecuteOptions) *bytes.Buffer {
	return Execute(ctx, b.Conn, btns, auth.Auth(ctx), panel, options)
}

func (b *Base) ExecuteTmplWithMenu(ctx *context.Context, panel types.Panel, options template.ExecuteOptions) *bytes.Buffer {
	return ExecuteWithMenu(ctx, b.Conn, *b.UI.NavButtons, auth.Auth(ctx), panel, b.Name(), b.Title(), options)
}

func (b *Base) ExecuteTmplWithCustomMenu(ctx *context.Context, panel types.Panel, menu *menu.Menu, options template.ExecuteOptions) *bytes.Buffer {
	return ExecuteWithCustomMenu(ctx, *b.UI.NavButtons, auth.Auth(ctx), panel, menu, b.Title(), options)
}

func (b *Base) ExecuteTmplWithMenuAndNavButtons(ctx *context.Context, panel types.Panel, menu *menu.Menu,
	btns types.Buttons, options template.ExecuteOptions) *bytes.Buffer {
	return ExecuteWithMenu(ctx, b.Conn, btns, auth.Auth(ctx), panel, b.Name(), b.Title(), options)
}

func (b *Base) NewMenu(data menu.NewMenuData) (int64, error) {
	return menu.NewMenu(b.Conn, data)
}

func (b *Base) HTML(ctx *context.Context, panel types.Panel, options ...template.ExecuteOptions) {
	buf := b.ExecuteTmpl(ctx, panel, template.GetExecuteOptions(options))
	ctx.HTMLByte(http.StatusOK, buf.Bytes())
}

func (b *Base) HTMLCustomMenu(ctx *context.Context, panel types.Panel, menu *menu.Menu, options ...template.ExecuteOptions) {
	buf := b.ExecuteTmplWithCustomMenu(ctx, panel, menu, template.GetExecuteOptions(options))
	ctx.HTMLByte(http.StatusOK, buf.Bytes())
}

func (b *Base) HTMLMenu(ctx *context.Context, panel types.Panel, options ...template.ExecuteOptions) {
	buf := b.ExecuteTmplWithMenu(ctx, panel, template.GetExecuteOptions(options))
	ctx.HTMLByte(http.StatusOK, buf.Bytes())
}

func (b *Base) HTMLBtns(ctx *context.Context, panel types.Panel, btns types.Buttons, options ...template.ExecuteOptions) {
	buf := b.ExecuteTmplWithNavButtons(ctx, panel, btns, template.GetExecuteOptions(options))
	ctx.HTMLByte(http.StatusOK, buf.Bytes())
}

func (b *Base) HTMLMenuWithBtns(ctx *context.Context, panel types.Panel, menu *menu.Menu, btns types.Buttons, options ...template.ExecuteOptions) {
	buf := b.ExecuteTmplWithMenuAndNavButtons(ctx, panel, menu, btns, template.GetExecuteOptions(options))
	ctx.HTMLByte(http.StatusOK, buf.Bytes())
}

func (b *Base) HTMLFile(ctx *context.Context, path string, data map[string]interface{}, options ...template.ExecuteOptions) {

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

	b.HTML(ctx, panel, options...)
}

func (b *Base) HTMLFiles(ctx *context.Context, data map[string]interface{}, files []string, options ...template.ExecuteOptions) {
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

	b.HTML(ctx, panel, options...)
}

type BasePlugin struct {
	Base
	Info      Info
	IndexURL  string
	Installed bool
}

func (b *BasePlugin) GetInfo() Info       { return b.Info }
func (b *BasePlugin) Name() string        { return b.Info.Name }
func (b *BasePlugin) GetIndexURL() string { return b.IndexURL }
func (b *BasePlugin) IsInstalled() bool   { return b.Installed }

func NewBasePluginWithInfo(info Info) Plugin {
	return &BasePlugin{Info: info}
}

func NewBasePluginWithInfoAndIndexURL(info Info, u string, installed bool) Plugin {
	return &BasePlugin{Info: info, IndexURL: u, Installed: installed}
}

func GetPluginsWithInfos(info []Info) Plugins {
	p := make(Plugins, len(info))
	for k, i := range info {
		p[k] = NewBasePluginWithInfo(i)
	}
	return p
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
	panel types.Panel, options template.ExecuteOptions) *bytes.Buffer {
	tmpl, tmplName := template.Get(config.GetTheme()).GetTemplate(ctx.IsPjax())

	return template.Execute(&template.ExecuteParam{
		User:       user,
		TmplName:   tmplName,
		Tmpl:       tmpl,
		Panel:      panel,
		Config:     config.Get(),
		Menu:       menu.GetGlobalMenu(user, conn, ctx.Lang()).SetActiveClass(config.URLRemovePrefix(ctx.Path())),
		Animation:  options.Animation,
		Buttons:    navButtons.CheckPermission(user),
		NoCompress: options.NoCompress,
		IsPjax:     ctx.IsPjax(),
		Iframe:     ctx.IsIframe(),
	})
}

func ExecuteWithCustomMenu(ctx *context.Context,
	navButtons types.Buttons,
	user models.UserModel,
	panel types.Panel,
	menu *menu.Menu, logo string, options template.ExecuteOptions) *bytes.Buffer {

	tmpl, tmplName := template.Get(config.GetTheme()).GetTemplate(ctx.IsPjax())

	return template.Execute(&template.ExecuteParam{
		User:       user,
		TmplName:   tmplName,
		Tmpl:       tmpl,
		Panel:      panel,
		Config:     config.Get(),
		Menu:       menu,
		Animation:  options.Animation,
		Buttons:    navButtons.CheckPermission(user),
		NoCompress: options.NoCompress,
		Logo:       template2.HTML(logo),
		IsPjax:     ctx.IsPjax(),
		Iframe:     ctx.IsIframe(),
	})
}

func ExecuteWithMenu(ctx *context.Context,
	conn db.Connection,
	navButtons types.Buttons,
	user models.UserModel,
	panel types.Panel,
	name, logo string, options template.ExecuteOptions) *bytes.Buffer {

	tmpl, tmplName := template.Get(config.GetTheme()).GetTemplate(ctx.IsPjax())

	btns := options.NavDropDownButton
	if btns == nil {
		btns = []*types.NavDropDownItemButton{
			types.GetDropDownItemButton(language.GetFromHtml("plugin setting"),
				action.Jump(config.Url("/info/plugin_"+name+"/edit"))),
			types.GetDropDownItemButton(language.GetFromHtml("menus manage"),
				action.Jump(config.Url("/menu?__plugin_name="+name))),
		}
	} else {
		btns = append(btns, []*types.NavDropDownItemButton{
			types.GetDropDownItemButton(language.GetFromHtml("plugin setting"),
				action.Jump(config.Url("/info/plugin_"+name+"/edit"))),
			types.GetDropDownItemButton(language.GetFromHtml("menus manage"),
				action.Jump(config.Url("/menu?__plugin_name="+name))),
		}...)
	}

	return template.Execute(&template.ExecuteParam{
		User:      user,
		TmplName:  tmplName,
		Tmpl:      tmpl,
		Panel:     panel,
		Config:    config.Get(),
		Menu:      menu.GetGlobalMenu(user, conn, ctx.Lang(), name).SetActiveClass(config.URLRemovePrefix(ctx.Path())),
		Animation: options.Animation,
		Buttons: navButtons.Copy().
			RemoveInfoNavButton().
			RemoveSiteNavButton().
			RemoveToolNavButton().
			Add(types.GetDropDownButton("", icon.Gear, btns)).CheckPermission(user),
		NoCompress: options.NoCompress,
		Logo:       template2.HTML(logo),
		IsPjax:     ctx.IsPjax(),
		Iframe:     ctx.IsIframe(),
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

func FindByName(name string) (Plugin, bool) {
	for _, v := range pluginList {
		if v.Name() == name {
			return v, true
		}
	}
	return nil, false
}

func FindByNameAll(name string) (Plugin, bool) {
	for _, v := range allPluginList {
		if v.Name() == name {
			return v, true
		}
	}
	return nil, false
}

var (
	pluginList    = make(Plugins, 0)
	allPluginList = make(Plugins, 0)
)

func Exist(p Plugin) bool {
	return pluginList.Exist(p)
}

func Add(p Plugin) {
	// TODO: 验证插件合法性
	pluginList = pluginList.Add(p)
}

func GetAll(req remote_server.GetOnlineReq, token string) (Plugins, Page) {

	plugs := make(Plugins, 0)
	page := Page{}

	res, err := remote_server.GetOnline(req, token)

	if err != nil {
		return plugs, page
	}

	var data GetOnlineRes
	err = json.Unmarshal(res, &data)
	if err != nil {
		return plugs, page
	}

	if data.Code != 0 {
		return plugs, page
	}

	plugs = GetPluginsWithInfos(data.Data.List)
	page = data.Data.Page

	for index, p := range plugs {
		for key, value := range pluginList {
			if value.Name() == p.Name() {
				info := pluginList[key].GetInfo()
				info.CanUpdate = utils.CompareVersion(info.Version, plugs[index].GetInfo().Version)
				info.OldVersion = info.Version
				info.Downloaded = true
				info.Description = language.GetWithScope(info.Description, info.Name)
				info.Title = language.GetWithScope(info.Title, info.Name)
				info.Version = plugs[index].GetInfo().Version
				plugs[index] = NewBasePluginWithInfoAndIndexURL(info, value.GetIndexURL(), value.IsInstalled())
				break
			}
		}
	}

	for _, p := range plugs {
		exist := false
		for _, pp := range allPluginList {
			if pp.Name() == p.Name() {
				exist = true
				break
			}
		}
		if !exist {
			allPluginList = append(allPluginList, p)
		}
	}

	return plugs, page
}

func Get() Plugins {
	var plugs = make(Plugins, len(pluginList))
	copy(plugs, pluginList)
	return plugs
}

type GetOnlineRes struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data GetOnlineResData `json:"data"`
}

type GetOnlineResData struct {
	List    []Info `json:"list"`
	Count   int    `json:"count"`
	HasMore bool   `json:"has_more"`
	Page    Page   `json:"page"`
}

type Page struct {
	CSS  string `json:"css"`
	HTML string `json:"html"`
	JS   string `json:"js"`
}
