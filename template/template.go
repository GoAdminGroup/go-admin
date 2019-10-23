// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	"errors"
	"fmt"
	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/login"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"plugin"
	"strings"
	"sync"
)

// Template is the interface which contains methods of ui components.
// It will be used in the plugins for custom the ui.
type Template interface {
	// Components must
	Col() types.ColAttribute
	Row() types.RowAttribute
	Form() types.FormAttribute
	Table() types.TableAttribute
	DataTable() types.DataTableAttribute
	Tree() types.TreeAttribute
	Label() types.LabelAttribute
	Tabs() types.TabsAttribute
	Alert() types.AlertAttribute

	// Components
	Box() types.BoxAttribute
	Image() types.ImgAttribute
	SmallBox() types.SmallBoxAttribute
	InfoBox() types.InfoBoxAttribute
	Paginator() types.PaginatorAttribute
	AreaChart() types.AreaChartAttribute
	ProgressGroup() types.ProgressGroupAttribute
	LineChart() types.LineChartAttribute
	BarChart() types.BarChartAttribute
	ProductList() types.ProductListAttribute
	Description() types.DescriptionAttribute
	PieChart() types.PieChartAttribute
	ChartLegend() types.ChartLegendAttribute
	Popup() types.PopupAttribute

	// Builder methods
	GetTmplList() map[string]string
	GetAssetList() []string
	GetAsset(string) ([]byte, error)
	GetTemplate(bool) (*template.Template, string)
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

func GetAssetLists() []string {
	assets := make([]string, 0)
	for _, comp := range compMap {
		assets = append(assets, comp.GetAssetList()...)
	}
	return assets
}

// AddComp makes a component available by the provided name.
// If Add is called twice with the same name or if component is nil,
// it panics.
func AddComp(name string, comp Component) {
	compMu.Lock()
	defer compMu.Unlock()
	if comp == nil {
		panic("component is nil")
	}
	if _, dup := compMap[name]; dup {
		panic("add component twice " + name)
	}
	compMap[name] = comp
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

	buf := new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(user, *globalMenu, panel, config))
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
	}
}

type BaseComponent struct {
}

func (b BaseComponent) GetAssetList() []string {
	return make([]string, 0)
}

func (b BaseComponent) GetAsset(name string) ([]byte, error) {
	return nil, nil
}
