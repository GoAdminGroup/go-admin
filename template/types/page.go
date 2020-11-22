// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package types

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	textTmpl "text/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/modules/system"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
)

// Attribute is the component interface of template. Every component of
// template should implement it.
type Attribute struct {
	TemplateList map[string]string
	Separation   bool
}

// Page used in the template as a top variable.
type Page struct {
	// User is the login user.
	User models.UserModel

	// Menu is the left side menu of the template.
	Menu menu.Menu

	// Panel is the main content of template.
	Panel Panel

	// System contains some system info.
	System SystemInfo

	// UrlPrefix is the prefix of url.
	UrlPrefix string

	// Title is the title of the web page.
	Title string

	// Logo is the logo of the template.
	Logo template.HTML

	// MiniLogo is the downsizing logo of the template.
	MiniLogo template.HTML

	// ColorScheme is the color scheme of the template.
	ColorScheme string

	// IndexUrl is the home page url of the site.
	IndexUrl string

	// AssetUrl is the cdn link of assets
	CdnUrl string

	// Custom html in the tag head.
	CustomHeadHtml template.HTML

	// Custom html after body.
	CustomFootHtml template.HTML

	TmplHeadHTML template.HTML
	TmplFootJS   template.HTML

	// Components assets
	AssetsList template.HTML

	// Footer info
	FooterInfo template.HTML

	// Load as Iframe or not
	Iframe bool

	// Whether update menu or not
	UpdateMenu bool

	// Top Nav Buttons
	navButtons     Buttons
	NavButtonsHTML template.HTML
}

type NewPageParam struct {
	User           models.UserModel
	Menu           *menu.Menu
	UpdateMenu     bool
	Panel          Panel
	Logo           template.HTML
	Assets         template.HTML
	Buttons        Buttons
	Iframe         bool
	TmplHeadHTML   template.HTML
	TmplFootJS     template.HTML
	NavButtonsHTML template.HTML
	NavButtonsJS   template.HTML
}

func (param *NewPageParam) NavButtonsAndJS() (template.HTML, template.HTML) {
	navBtnFooter := template.HTML("")
	navBtn := template.HTML("")
	btnJS := template.JS("")

	for _, btn := range param.Buttons {
		if btn.IsType(ButtonTypeNavDropDown) {
			content, js := btn.Content()
			navBtn += content
			btnJS += js
			for _, item := range btn.(*NavDropDownButton).Items {
				navBtnFooter += item.GetAction().FooterContent()
				_, js := item.Content()
				btnJS += js
			}
		} else {
			navBtnFooter += btn.GetAction().FooterContent()
			content, js := btn.Content()
			navBtn += content
			btnJS += js
		}
	}

	return template.HTML(ParseTableDataTmpl(navBtn)),
		navBtnFooter + template.HTML(ParseTableDataTmpl(`<script>`+btnJS+`</script>`))
}

func NewPage(param *NewPageParam) *Page {

	if param.NavButtonsHTML == template.HTML("") {
		param.NavButtonsHTML, param.NavButtonsJS = param.NavButtonsAndJS()
	}

	logo := param.Logo
	if logo == template.HTML("") {
		logo = config.GetLogo()
	}

	return &Page{
		User:       param.User,
		Menu:       *param.Menu,
		Panel:      param.Panel,
		UpdateMenu: param.UpdateMenu,
		System: SystemInfo{
			Version: system.Version(),
			Theme:   config.GetTheme(),
		},
		UrlPrefix:      config.AssertPrefix(),
		Title:          config.GetTitle(),
		Logo:           logo,
		MiniLogo:       config.GetMiniLogo(),
		ColorScheme:    config.GetColorScheme(),
		IndexUrl:       config.GetIndexURL(),
		CdnUrl:         config.GetAssetUrl(),
		CustomHeadHtml: config.GetCustomHeadHtml(),
		CustomFootHtml: config.GetCustomFootHtml() + param.NavButtonsJS,
		FooterInfo:     config.GetFooterInfo(),
		AssetsList:     param.Assets,
		navButtons:     param.Buttons,
		Iframe:         param.Iframe,
		NavButtonsHTML: param.NavButtonsHTML,
		TmplHeadHTML:   param.TmplHeadHTML,
		TmplFootJS:     param.TmplFootJS,
	}
}

func (page *Page) AddButton(title template.HTML, icon string, action Action) *Page {
	page.navButtons = append(page.navButtons, GetNavButton(title, icon, action))
	page.CustomFootHtml += action.FooterContent()
	return page
}

func NewPagePanel(panel Panel) *Page {
	return &Page{
		Panel: panel,
		System: SystemInfo{
			Version: system.Version(),
		},
	}
}

// SystemInfo contains basic info of system.
type SystemInfo struct {
	Version string
	Theme   string
}

type TableRowData struct {
	Id    template.HTML
	Ids   template.HTML
	Value map[string]InfoItem
}

func ParseTableDataTmpl(content interface{}) string {
	var (
		c  string
		ok bool
	)
	if c, ok = content.(string); !ok {
		if cc, ok := content.(template.HTML); ok {
			c = string(cc)
		} else {
			c = string(content.(template.JS))
		}
	}
	t := template.New("row_data_tmpl")
	t, _ = t.Parse(c)
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, TableRowData{Ids: `typeof(selectedRows)==="function" ? selectedRows().join() : ""`})
	return buf.String()
}

func ParseTableDataTmplWithID(id template.HTML, content string, value ...map[string]InfoItem) string {
	t := textTmpl.New("row_data_tmpl")
	t, _ = t.Parse(content)
	buf := new(bytes.Buffer)
	v := make(map[string]InfoItem)
	if len(value) > 0 {
		v = value[0]
	}
	_ = t.Execute(buf, TableRowData{
		Id:    id,
		Ids:   `typeof(selectedRows)==="function" ? selectedRows().join() : ""`,
		Value: v,
	})
	return buf.String()
}

// Panel contains the main content of the template which used as pjax.
type Panel struct {
	Title       template.HTML
	Description template.HTML
	Content     template.HTML

	CSS template.CSS
	JS  template.JS
	Url string

	// Whether to toggle the sidebar
	MiniSidebar bool

	// Auto refresh page switch.
	AutoRefresh bool
	// Refresh page intervals, the unit is second.
	RefreshInterval []int

	Callbacks Callbacks
}

type Component interface {
	GetContent() template.HTML
	GetJS() template.JS
	GetCSS() template.CSS
	GetCallbacks() Callbacks
}

func (p Panel) AddComponent(comp Component) Panel {
	p.JS += comp.GetJS()
	p.CSS += comp.GetCSS()
	p.Content += comp.GetContent()
	p.Callbacks = append(p.Callbacks, comp.GetCallbacks()...)
	return p
}

func (p Panel) AddJS(js template.JS) Panel {
	p.JS += js
	return p
}

func (p Panel) GetContent(params ...bool) Panel {

	compress := false

	if len(params) > 0 {
		compress = params[0]
	}

	var (
		animation, style, remove = template.HTML(""), template.HTML(""), template.HTML("")
		ani                      = config.GetAnimation()
	)

	if ani.Type != "" && (len(params) < 2 || params[1]) {
		animation = template.HTML(` class='pjax-container-content animated ` + ani.Type + `'`)
		if ani.Delay != 0 {
			style = template.HTML(fmt.Sprintf(`animation-delay: %fs;-webkit-animation-delay: %fs;`, ani.Delay, ani.Delay))
		}
		if ani.Duration != 0 {
			style = template.HTML(fmt.Sprintf(`animation-duration: %fs;-webkit-animation-duration: %fs;`, ani.Duration, ani.Duration))
		}
		if style != "" {
			style = ` style="` + style + `"`
		}
		remove = template.HTML(`<script>
		$('.pjax-container-content .modal.fade').on('show.bs.modal', function (event) {
            // Fix Animate.css
			$('.pjax-container-content').removeClass('` + ani.Type + `');
        });
		</script>`)
	}

	p.Content = `<div` + animation + style + ">" + p.Content + "</div>" + remove
	if p.MiniSidebar {
		p.Content += `<script>$("body").addClass("sidebar-collapse")</script>`
	}
	if p.AutoRefresh {
		refreshTime := 60
		if len(p.RefreshInterval) > 0 {
			refreshTime = p.RefreshInterval[0]
		}

		p.Content += `<script>
window.setTimeout(function(){
	$.pjax.reload('#pjax-container');	
}, ` + template.HTML(strconv.Itoa(refreshTime*1000)) + `);
</script>`
	}

	if compress {
		utils.CompressedContent(&p.Content)
	}

	return p
}

type GetPanelFn func(ctx interface{}) (Panel, error)

type GetPanelInfoFn func(ctx *context.Context) (Panel, error)
