package adminlte

import (
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/template/adminlte/resource"
	"github.com/chenhg5/go-admin/template/adminlte/tmpl"
	"github.com/chenhg5/go-admin/template/components"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type Theme struct {
	Name string
}

var Adminlte = Theme{
	Name: "adminlte",
}

func GetAdminlte() *Theme {
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
	return resource.Asset(path)
}

func (*Theme) GetAssetList() []string {
	return asserts
}

var Attribute = types.Attribute{
	TemplateList: tmpl.List,
}

func (*Theme) Box() types.BoxAttribute {
	return &components.BoxAttribute{
		Name:       "box",
		Header:     template.HTML(""),
		Body:       template.HTML(""),
		Footer:     template.HTML(""),
		Title:      "",
		HeadBorder: "",
		Attribute:  Attribute,
	}
}

func (*Theme) Col() types.ColAttribute {
	return &components.ColAttribute{
		Name:      "col",
		Size:      "col-md-2",
		Content:   "",
		Attribute: Attribute,
	}
}

func (*Theme) Form() types.FormAttribute {
	return &components.FormAttribute{
		Name:      "form",
		Content:   []types.Form{},
		Url:       "/",
		Method:    "post",
		InfoUrl:   "",
		Title:     "edit",
		Attribute: Attribute,
	}
}

func (*Theme) Image() types.ImgAttribute {
	return &components.ImgAttribute{
		Name:      "image",
		Witdh:     "50",
		Height:    "50",
		Src:       "",
		Attribute: Attribute,
	}
}

func (*Theme) SmallBox() types.SmallBoxAttribute {
	return &components.SmallBoxAttribute{
		Name:      "smallbox",
		Title:     "title",
		Value:     "value",
		Url:       "/",
		Color:     "aqua",
		Attribute: Attribute,
	}
}

func (*Theme) InfoBox() types.InfoBoxAttribute {
	return &components.InfoBoxAttribute{
		Name:      "infobox",
		Text:      "title",
		Icon:      "ion-ios-cart-outline",
		Number:    "90",
		Color:     "red",
		Attribute: Attribute,
	}
}

func (*Theme) AreaChart() types.AreaChartAttribute {
	return &components.AreaChartAttribute{
		Name:      "area-chart",
		Attribute: Attribute,
	}
}

func (*Theme) ProgressGroup() types.ProgressGroupAttribute {
	return &components.ProgressGroupAttribute{
		Name:      "progress-group",
		Attribute: Attribute,
	}
}

func (*Theme) Description() types.DescriptionAttribute {
	return &components.DescriptionAttribute{
		Name:      "description",
		Attribute: Attribute,
	}
}

func (*Theme) PieChart() types.PieChartAttribute {
	return &components.PieChartAttribute{
		Name:      "pie-chart",
		Attribute: Attribute,
	}
}

func (*Theme) LineChart() types.LineChartAttribute {
	return &components.LineChartAttribute{
		Name:      "line-chart",
		Attribute: Attribute,
	}
}

func (*Theme) BarChart() types.BarChartAttribute {
	return &components.BarChartAttribute{
		Name:      "bar-chart",
		Attribute: Attribute,
	}
}

func (*Theme) ChartLegend() types.ChartLegendAttribute {
	return &components.ChartLegendAttribute{
		Name:      "chart-legend",
		Attribute: Attribute,
	}
}

func (*Theme) ProductList() types.ProductListAttribute {
	return &components.ProductListAttribute{
		Name:      "productlist",
		Attribute: Attribute,
	}
}

func (*Theme) Tabs() types.TabsAttribute {
	return &components.TabsAttribute{
		Name:      "tabs",
		Attribute: Attribute,
	}
}

func (*Theme) Alert() types.AlertAttribute {
	return &components.AlertAttribute{
		Name:      "alert",
		Attribute: Attribute,
	}
}

func (*Theme) Label() types.LabelAttribute {
	return &components.LabelAttribute{
		Name:      "label",
		Color:     "success",
		Content:   "",
		Attribute: Attribute,
	}
}

func (*Theme) Popup() types.PopupAttribute {
	return &components.PopupAttribute{
		Name:      "popup",
		Attribute: Attribute,
	}
}

func (*Theme) Paginator() types.PaginatorAttribute {
	return &components.PaginatorAttribute{
		Name:      "paginator",
		Attribute: Attribute,
	}
}

func (*Theme) Row() types.RowAttribute {
	return &components.RowAttribute{
		Name:      "row",
		Content:   "",
		Attribute: Attribute,
	}
}

func (*Theme) Table() types.TableAttribute {
	return &components.TableAttribute{
		Name:      "table",
		Thead:     []map[string]string{},
		InfoList:  []map[string]template.HTML{},
		Type:      "normal",
		Attribute: Attribute,
	}
}

func (theme *Theme) DataTable() types.DataTableAttribute {
	return &components.DataTableAttribute{
		TableAttribute: *(theme.Table().SetType("data-table").(*components.TableAttribute)),
		EditUrl:        "",
		NewUrl:         "",
		Attribute:      Attribute,
	}
}

func (*Theme) Tree() types.TreeAttribute {
	return &components.TreeAttribute{
		Name:      "tree",
		Tree:      []menu.MenuItem{},
		Attribute: Attribute,
	}
}
