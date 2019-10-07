package adminlte

import (
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template/adminlte/resource"
	"github.com/chenhg5/go-admin/template/adminlte/tmpl"
	"github.com/chenhg5/go-admin/template/components"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

const (
	COLORSCHEME_SKIN_BLACK        = "skin-black"
	COLORSCHEME_SKIN_BLACK_LIGHT  = "skin-black-light"
	COLORSCHEME_SKIN_BLUE         = "skin-blue"
	COLORSCHEME_SKIN_BLUE_LIGHT   = "skin-blue-light"
	COLORSCHEME_SKIN_GREEN        = "skin-green"
	COLORSCHEME_SKIN_GREEN_LIGHT  = "skin-green-light"
	COLORSCHEME_SKIN_PURPLE       = "skin-purple"
	COLORSCHEME_SKIN_PURPLE_LIGHT = "skin-purple-light"
	COLORSCHEME_SKIN_RED          = "skin-red"
	COLORSCHEME_SKIN_RED_LIGHT    = "skin-red-light"
	COLORSCHEME_SKIN_YELLOW       = "skin-yellow"
	COLORSCHEME_SKIN_YELLOW_LIGHT = "skin-yellow-light"
)

type Theme struct {
	Name string
	components.Base
}

var Adminlte = Theme{
	Name: "adminlte",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: tmpl.List,
		},
	},
}

func Get() *Theme {
	return &Adminlte
}

func (*Theme) GetTmplList() map[string]string {
	return tmpl.List
}

func (*Theme) GetTemplate(isPjax bool) (tmpler *template.Template, name string) {
	var err error

	if !isPjax {
		name = "layout"
		tmpler, err = template.New("layout").Funcs(template.FuncMap{
			"lang":     language.Get,
			"langHtml": language.GetFromHtml,
			"isLinkUrl": func(s string) bool {
				return (len(s) > 7 && s[:7] == "http://") || (len(s) > 8 && s[:8] == "https://")
			},
		}).Parse(tmpl.List["layout"] +
			tmpl.List["head"] + tmpl.List["header"] + tmpl.List["sidebar"] +
			tmpl.List["footer"] + tmpl.List["js"] + tmpl.List["menu"] +
			tmpl.List["admin_panel"] + tmpl.List["content"])
	} else {
		name = "content"
		tmpler, err = template.New("content").Funcs(template.FuncMap{
			"lang":     language.Get,
			"langHtml": language.GetFromHtml,
		}).Parse(tmpl.List["admin_panel"] + tmpl.List["content"])
	}

	if err != nil {
		panic(err)
	}

	return
}

func (*Theme) GetAsset(path string) ([]byte, error) {
	path = "template/adminlte/resource" + path
	return resource.Asset(path)
}

func (*Theme) GetAssetList() []string {
	return resource.AssetsList
}
