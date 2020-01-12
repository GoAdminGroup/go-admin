// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path"
	"plugin"
	"strings"
	"sync"

	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/login"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Template is the interface which contains methods of ui components.
// It will be used in the plugins for custom the ui.
type Template interface {
	// Components

	// layout
	Col() types.ColAttribute
	Row() types.RowAttribute

	// form and table
	Form() types.FormAttribute
	Table() types.TableAttribute
	DataTable() types.DataTableAttribute

	Tree() types.TreeAttribute
	Tabs() types.TabsAttribute
	Alert() types.AlertAttribute

	Paginator() types.PaginatorAttribute
	Popup() types.PopupAttribute
	Box() types.BoxAttribute

	Label() types.LabelAttribute
	Image() types.ImgAttribute

	Button() types.ButtonAttribute

	// Builder methods
	GetTmplList() map[string]string
	GetAssetList() []string
	GetAsset(string) ([]byte, error)
	GetTemplate(bool) (*template.Template, string)
}

func HTML(s string) template.HTML {
	return template.HTML(s)
}

func CSS(s string) template.CSS {
	return template.CSS(s)
}

func JS(s string) template.JS {
	return template.JS(s)
}

// The templateMap contains templates registered.
var templateMap = make(map[string]Template)

// Get the template interface by theme name. If the
// name is not found, it panics.
func Get(theme string) Template {
	if temp, ok := templateMap[theme]; ok {
		return temp
	}
	panic("wrong theme name")
}

// Get the default template with the theme name set with the global config.
// If the name is not found, it panics.
func Default() Template {
	if temp, ok := templateMap[c.Get().Theme]; ok {
		return temp
	}
	panic("wrong theme name")
}

var (
	templateMu sync.Mutex
	compMu     sync.Mutex
)

// Add makes a template available by the provided theme name.
// If Add is called twice with the same name or if template is nil,
// it panics.
func Add(name string, temp Template) {
	templateMu.Lock()
	defer templateMu.Unlock()
	if temp == nil {
		panic("template is nil")
	}
	if _, dup := templateMap[name]; dup {
		panic("add template twice " + name)
	}
	templateMap[name] = temp
}

func AddFromPlugin(name string, mod string) {

	plug, err := plugin.Open(mod)
	if err != nil {
		logger.Error("AddFromPlugin err", err)
		panic(err)
	}

	tempPlugin, err := plug.Lookup(strings.Title(name))
	if err != nil {
		logger.Error("AddFromPlugin err", err)
		panic(err)
	}

	var temp Template
	temp, ok := tempPlugin.(Template)
	if !ok {
		logger.Error("AddFromPlugin err: unexpected type from module symbol")
		panic(errors.New("AddFromPlugin err: unexpected type from module symbol"))
	}

	Add(name, temp)
}

// Component is the interface which stand for a ui component.
type Component interface {
	// GetTemplate return a *template.Template and a given key.
	GetTemplate() (*template.Template, string)

	// GetAssetList return the assets url suffix used in the component.
	// example:
	//
	// {{.UrlPrefix}}/assets/login/css/bootstrap.min.css => login/css/bootstrap.min.css
	//
	// See:
	// https://github.com/GoAdminGroup/go-admin/blob/master/template/login/theme1.tmpl#L32
	// https://github.com/GoAdminGroup/go-admin/blob/master/template/login/list.go
	GetAssetList() []string

	// GetAsset return the asset content according to the corresponding url suffix.
	// Asset content is recommended to use the tool go-bindata to generate.
	//
	// See: http://github.com/jteeuwen/go-bindata
	GetAsset(string) ([]byte, error)

	GetContent() template.HTML

	IsAPage() bool

	GetName() string
}

var compMap = map[string]Component{
	"login": login.GetLoginComponent(),
}

// GetComp gets the component by registered name. If the
// name is not found, it panics.
func GetComp(name string) Component {
	if comp, ok := compMap[name]; ok {
		return comp
	}
	panic("wrong component name")
}

func GetComponentAssetLists() []string {
	assets := make([]string, 0)
	for _, comp := range compMap {
		assets = append(assets, comp.GetAssetList()...)
	}
	return assets
}

func GetComponentAssetListsWithinPage() []string {
	assets := make([]string, 0)
	for _, comp := range compMap {
		if !comp.IsAPage() {
			assets = append(assets, comp.GetAssetList()...)
		}
	}
	return assets
}

func GetComponentAssetListsHTML() (res template.HTML) {
	assets := GetComponentAssetListsWithinPage()
	for i := 0; i < len(assets); i++ {
		res += getHTMLFromAssetUrl(assets[i])
	}
	return
}

func getHTMLFromAssetUrl(s string) template.HTML {
	fileSuffix := path.Ext(s)
	fileSuffix = strings.Replace(fileSuffix, ".", "", -1)

	if fileSuffix == "css" {
		return template.HTML(`<link rel="stylesheet" href="` + c.Get().AssetUrl + c.Get().Url("/assets"+s) + `">`)
	}
	if fileSuffix == "js" {
		return template.HTML(`<script src="` + c.Get().AssetUrl + c.Get().Url("/assets"+s) + `"></script>`)
	}
	return ""
}

func GetAsset(path string) ([]byte, error) {
	for _, comp := range compMap {
		res, err := comp.GetAsset(path)
		if err == nil {
			return res, err
		}
	}
	return nil, errors.New(path + " not found")
}

// AddComp makes a component available by the provided name.
// If Add is called twice with the same name or if component is nil,
// it panics.
func AddComp(comp Component) {
	compMu.Lock()
	defer compMu.Unlock()
	if comp == nil {
		panic("component is nil")
	}
	if _, dup := compMap[comp.GetName()]; dup {
		panic("add component twice " + comp.GetName())
	}
	compMap[comp.GetName()] = comp
}

// AddLoginComp add the specified login component.
func AddLoginComp(comp Component) {
	compMu.Lock()
	defer compMu.Unlock()
	compMap["login"] = comp
}

// SetComp makes a component available by the provided name.
// If the value corresponding to the key is empty or if component is nil,
// it panics.
func SetComp(name string, comp Component) {
	compMu.Lock()
	defer compMu.Unlock()
	if comp == nil {
		panic("component is nil")
	}
	if _, dup := compMap[name]; dup {
		compMap[name] = comp
	}
}

func Execute(tmpl *template.Template,
	tmplName string,
	user models.UserModel,
	panel types.Panel,
	config c.Config,
	globalMenu *menu.Menu) *bytes.Buffer {

	if !config.Debug {
		utils.CompressedContent(&panel.Content)
	}
	buf := new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user, *globalMenu, panel, config, GetComponentAssetListsHTML()))
	if err != nil {
		fmt.Println("Execute err", err)
	}
	return buf
}

func DefaultFuncMap() template.FuncMap {
	return template.FuncMap{
		"lang":     language.Get,
		"langHtml": language.GetFromHtml,
		"link": func(cdnUrl, prefixUrl, assetsUrl string) string {
			if cdnUrl == "" {
				return prefixUrl + assetsUrl
			}
			return cdnUrl + assetsUrl
		},
		"isLinkUrl": func(s string) bool {
			return (len(s) > 7 && s[:7] == "http://") || (len(s) > 8 && s[:8] == "https://")
		},
		"render": func(s, old, repl template.HTML) template.HTML {
			return template.HTML(strings.Replace(string(s), string(old), string(repl), -1))
		},
		"renderJS": func(s template.JS, old, repl template.HTML) template.JS {
			return template.JS(strings.Replace(string(s), string(old), string(repl), -1))
		},
	}
}

type BaseComponent struct{}

func (b BaseComponent) GetAssetList() []string               { return make([]string, 0) }
func (b BaseComponent) GetAsset(name string) ([]byte, error) { return nil, nil }
