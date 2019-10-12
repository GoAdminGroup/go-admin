// Copyright 2019 GoAdmin.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	"fmt"
	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/login"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"sync"
)

// Template is the interface which contains methods of ui components.
// It will be used in the plugins for custom the ui.
type Template interface {
	// Components
	Form() types.FormAttribute
	Box() types.BoxAttribute
	Col() types.ColAttribute
	Image() types.ImgAttribute
	SmallBox() types.SmallBoxAttribute
	Label() types.LabelAttribute
	Row() types.RowAttribute
	Table() types.TableAttribute
	DataTable() types.DataTableAttribute
	Tree() types.TreeAttribute
	InfoBox() types.InfoBoxAttribute
	Paginator() types.PaginatorAttribute
	AreaChart() types.AreaChartAttribute
	ProgressGroup() types.ProgressGroupAttribute
	LineChart() types.LineChartAttribute
	BarChart() types.BarChartAttribute
	ProductList() types.ProductListAttribute
	Description() types.DescriptionAttribute
	Alert() types.AlertAttribute
	PieChart() types.PieChartAttribute
	ChartLegend() types.ChartLegendAttribute
	Tabs() types.TabsAttribute
	Popup() types.PopupAttribute

	// Builder methods
	GetTmplList() map[string]string
	GetAssetList() []string
	GetAsset(string) ([]byte, error)
	GetTemplate(bool) (*template.Template, string)
}

// The templateMap contains templates registered.
var templateMap = make(map[string]Template, 0)

// Get the template interface by theme name. If the
// name is not found, it panics.
func Get(theme string) Template {
	if temp, ok := templateMap[theme]; ok {
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
