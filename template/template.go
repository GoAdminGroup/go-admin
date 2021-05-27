// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	"errors"
	"html/template"
	"path"
	"plugin"
	"strconv"
	"strings"
	"sync"

	c "github.com/GoAdminGroup/go-admin/modules/config"
	errors2 "github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/template/login"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Template is the interface which contains methods of ui components.
// It will be used in the plugins for custom the ui.
type Template interface {
	Name() string

	// Components

	// layout
	Col() types.ColAttribute
	Row() types.RowAttribute

	// form and table
	Form() types.FormAttribute
	Table() types.TableAttribute
	DataTable() types.DataTableAttribute

	TreeView() types.TreeViewAttribute
	Tree() types.TreeAttribute
	Tabs() types.TabsAttribute
	Alert() types.AlertAttribute
	Link() types.LinkAttribute

	Paginator() types.PaginatorAttribute
	Popup() types.PopupAttribute
	Box() types.BoxAttribute

	Label() types.LabelAttribute
	Image() types.ImgAttribute

	Button() types.ButtonAttribute

	// Builder methods
	GetTmplList() map[string]string
	GetAssetList() []string
	GetAssetImportHTML(exceptComponents ...string) template.HTML
	GetAsset(string) ([]byte, error)
	GetTemplate(bool) (*template.Template, string)
	GetVersion() string
	GetRequirements() []string
	GetHeadHTML() template.HTML
	GetFootJS() template.HTML
	Get404HTML() template.HTML
	Get500HTML() template.HTML
	Get403HTML() template.HTML
}

type PageType uint8

const (
	NormalPage PageType = iota
	Missing404Page
	Error500Page
	NoPermission403Page
)

func GetPageTypeFromPageError(err errors2.PageError) PageType {
	if err == nil {
		return NormalPage
	} else if err == errors2.PageError403 {
		return NoPermission403Page
	} else if err == errors2.PageError404 {
		return Missing404Page
	} else {
		return Error500Page
	}
}

const (
	CompCol       = "col"
	CompRow       = "row"
	CompForm      = "form"
	CompTable     = "table"
	CompDataTable = "datatable"
	CompTree      = "tree"
	CompTreeView  = "treeview"
	CompTabs      = "tabs"
	CompAlert     = "alert"
	CompLink      = "link"
	CompPaginator = "paginator"
	CompPopup     = "popup"
	CompBox       = "box"
	CompLabel     = "label"
	CompImage     = "image"
	CompButton    = "button"
)

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

// Default get the default template with the theme name set with the global config.
// If the name is not found, it panics.
func Default() Template {
	if temp, ok := templateMap[c.GetTheme()]; ok {
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

// CheckRequirements check the theme and GoAdmin interdependence limit.
// The first return parameter means that whether GoAdmin version meets the requirement of the theme used or not.
// The second return parameter means that whether the version of theme used meets the requirement of GoAdmin or not.
func CheckRequirements() (bool, bool) {
	if !CheckThemeRequirements() {
		return false, true
	}
	// The theme which is not in the default official themes will be ignored.
	if !utils.InArray(DefaultThemeNames, Default().Name()) {
		return true, true
	}
	return true, VersionCompare(Default().GetVersion(), system.RequireThemeVersion()[Default().Name()])
}

func CheckThemeRequirements() bool {
	return VersionCompare(system.Version(), Default().GetRequirements())
}

func VersionCompare(toCompare string, versions []string) bool {
	for _, v := range versions {
		if v == toCompare || utils.CompareVersion(v, toCompare) {
			return true
		}
	}
	return false
}

func GetPageContentFromPageType(title, desc, msg string, pt PageType) (template.HTML, template.HTML, template.HTML) {
	if c.GetDebug() {
		return template.HTML(title), template.HTML(desc), Default().Alert().SetTitle(errors2.MsgWithIcon).Warning(msg)
	}

	if pt == Missing404Page {
		if c.GetCustom404HTML() != template.HTML("") {
			return "", "", c.GetCustom404HTML()
		} else {
			return "", "", Default().Get404HTML()
		}
	} else if pt == NoPermission403Page {
		if c.GetCustom404HTML() != template.HTML("") {
			return "", "", c.GetCustom403HTML()
		} else {
			return "", "", Default().Get403HTML()
		}
	} else {
		if c.GetCustom500HTML() != template.HTML("") {
			return "", "", c.GetCustom500HTML()
		} else {
			return "", "", Default().Get500HTML()
		}
	}
}

var DefaultThemeNames = []string{"sword", "adminlte"}

func Themes() []string {
	names := make([]string, len(templateMap))
	i := 0
	for k := range templateMap {
		names[i] = k
		i++
	}
	return names
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

	GetJS() template.JS
	GetCSS() template.CSS
	GetCallbacks() types.Callbacks
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

func GetComponentAsset() []string {
	assets := make([]string, 0)
	for _, comp := range compMap {
		assets = append(assets, comp.GetAssetList()...)
	}
	return assets
}

func GetComponentAssetWithinPage() []string {
	assets := make([]string, 0)
	for _, comp := range compMap {
		if !comp.IsAPage() {
			assets = append(assets, comp.GetAssetList()...)
		}
	}
	return assets
}

func GetComponentAssetImportHTML() (res template.HTML) {
	res = Default().GetAssetImportHTML(c.GetExcludeThemeComponents()...)
	assets := GetComponentAssetWithinPage()
	for i := 0; i < len(assets); i++ {
		res += getHTMLFromAssetUrl(assets[i])
	}
	return
}

func getHTMLFromAssetUrl(s string) template.HTML {
	switch path.Ext(s) {
	case ".css":
		return template.HTML(`<link rel="stylesheet" href="` + c.GetAssetUrl() + c.Url("/assets"+s) + `">`)
	case ".js":
		return template.HTML(`<script src="` + c.GetAssetUrl() + c.Url("/assets"+s) + `"></script>`)
	default:
		return ""
	}
}

func GetAsset(path string) ([]byte, error) {
	for _, comp := range compMap {
		res, err := comp.GetAsset(path)
		if err == nil {
			return res, nil
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

type ExecuteParam struct {
	User       models.UserModel
	Tmpl       *template.Template
	TmplName   string
	IsPjax     bool
	Panel      types.Panel
	Logo       template.HTML
	Config     *c.Config
	Menu       *menu.Menu
	Animation  bool
	Buttons    types.Buttons
	NoCompress bool
	Iframe     bool
}

func updateNavAndLogoJS(logo template.HTML) template.JS {
	if logo == template.HTML("") {
		return ""
	}
	return `$(function () {
	$(".logo-lg").html("` + template.JS(logo) + `");
});`
}

func updateNavJS(isPjax bool) template.JS {
	if !isPjax {
		return ""
	}
	return `$(function () {
	let lis = $(".navbar-custom-menu .nav.navbar-nav li");
	for (var i = lis.length - 8; i > -1; i--) {
		$(lis[i]).remove();
	}
	$(".navbar-custom-menu .nav.navbar-nav").prepend($("#navbar-nav-custom").html());
});`
}

type ExecuteOptions struct {
	Animation         bool
	NoCompress        bool
	HideSideBar       bool
	HideHeader        bool
	UpdateMenu        bool
	NavDropDownButton []*types.NavDropDownItemButton
}

func GetExecuteOptions(options []ExecuteOptions) ExecuteOptions {
	if len(options) == 0 {
		return ExecuteOptions{Animation: true}
	}
	return options[0]
}

func Execute(param *ExecuteParam) *bytes.Buffer {

	buf := new(bytes.Buffer)
	err := param.Tmpl.ExecuteTemplate(buf, param.TmplName,
		types.NewPage(&types.NewPageParam{
			User:       param.User,
			Menu:       param.Menu,
			Assets:     GetComponentAssetImportHTML(),
			Buttons:    param.Buttons,
			Iframe:     param.Iframe,
			UpdateMenu: param.IsPjax,
			Panel: param.Panel.
				GetContent(append([]bool{param.Config.IsProductionEnvironment() && !param.NoCompress},
					param.Animation)...).AddJS(param.Menu.GetUpdateJS(param.IsPjax)).
				AddJS(updateNavAndLogoJS(param.Logo)).AddJS(updateNavJS(param.IsPjax)),
			TmplHeadHTML: Default().GetHeadHTML(),
			TmplFootJS:   Default().GetFootJS(),
			Logo:         param.Logo,
		}))
	if err != nil {
		logger.Error("template execute error", err)
	}
	return buf
}

func WarningPanel(msg string, pts ...PageType) types.Panel {
	pt := Error500Page
	if len(pts) > 0 {
		pt = pts[0]
	}
	pageTitle, description, content := GetPageContentFromPageType(msg, msg, msg, pt)
	return types.Panel{
		Content:     content,
		Description: description,
		Title:       pageTitle,
	}
}

func WarningPanelWithDescAndTitle(msg, desc, title string, pts ...PageType) types.Panel {
	pt := Error500Page
	if len(pts) > 0 {
		pt = pts[0]
	}
	pageTitle, description, content := GetPageContentFromPageType(msg, desc, title, pt)
	return types.Panel{
		Content:     content,
		Description: description,
		Title:       pageTitle,
	}
}

var DefaultFuncMap = template.FuncMap{
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
		return template.HTML(strings.ReplaceAll(string(s), string(old), string(repl)))
	},
	"renderJS": func(s template.JS, old, repl template.HTML) template.JS {
		return template.JS(strings.ReplaceAll(string(s), string(old), string(repl)))
	},
	"divide": func(a, b int) int {
		return a / b
	},
	"renderRowDataHTML": func(id, content template.HTML, value ...map[string]types.InfoItem) template.HTML {
		return template.HTML(types.ParseTableDataTmplWithID(id, string(content), value...))
	},
	"renderRowDataJS": func(id template.HTML, content template.JS, value ...map[string]types.InfoItem) template.JS {
		return template.JS(types.ParseTableDataTmplWithID(id, string(content), value...))
	},
	"attr": func(s template.HTML) template.HTMLAttr {
		return template.HTMLAttr(s)
	},
	"js": func(s interface{}) template.JS {
		if ss, ok := s.(string); ok {
			return template.JS(ss)
		}
		if ss, ok := s.(template.HTML); ok {
			return template.JS(ss)
		}
		return ""
	},
	"changeValue": func(f types.FormField, index int) types.FormField {
		if len(f.ValueArr) > 0 {
			f.Value = template.HTML(f.ValueArr[index])
		}
		if len(f.OptionsArr) > 0 {
			f.Options = f.OptionsArr[index]
		}
		if f.FormType.IsSelect() {
			f.FieldClass += "_" + strconv.Itoa(index)
		}
		return f
	},
}

type BaseComponent struct {
	Name      string
	HTMLData  string
	CSS       template.CSS
	JS        template.JS
	Callbacks types.Callbacks
}

func (b *BaseComponent) IsAPage() bool                        { return false }
func (b *BaseComponent) GetName() string                      { return b.Name }
func (b *BaseComponent) GetAssetList() []string               { return make([]string, 0) }
func (b *BaseComponent) GetAsset(name string) ([]byte, error) { return nil, nil }
func (b *BaseComponent) GetJS() template.JS                   { return b.JS }
func (b *BaseComponent) GetCSS() template.CSS                 { return b.CSS }
func (b *BaseComponent) GetCallbacks() types.Callbacks        { return b.Callbacks }
func (b *BaseComponent) BindActionTo(action types.Action, id string) {
	action.SetBtnId(id)
	b.JS += action.Js()
	b.HTMLData += string(action.ExtContent())
	b.Callbacks = append(b.Callbacks, action.GetCallbacks())
}
func (b *BaseComponent) GetContentWithData(obj interface{}) template.HTML {
	buffer := new(bytes.Buffer)
	tmpl, defineName := b.GetTemplate()
	err := tmpl.ExecuteTemplate(buffer, defineName, obj)
	if err != nil {
		logger.Error(b.Name+" GetContent error:", err)
	}
	return template.HTML(buffer.String())
}

func (b *BaseComponent) GetTemplate() (*template.Template, string) {
	tmpl, err := template.New(b.Name).
		Funcs(DefaultFuncMap).
		Parse(b.HTMLData)

	if err != nil {
		logger.Error(b.Name+" GetTemplate Error: ", err)
	}

	return tmpl, b.Name
}
