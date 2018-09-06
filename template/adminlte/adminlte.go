package adminlte

import (
	"html/template"
	"goAdmin/template/types"
	"goAdmin/template/adminlte/components"
	"goAdmin/modules/menu"
	"goAdmin/template/adminlte/tmpl"
	"fmt"
	"goAdmin/template/adminlte/resource"
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
	var (
		err error
	)

	if !isPjax {
		name = "layout"
		tmpler, err = template.New("layout").Parse(tmpl.List["layout"] +
			tmpl.List["head"] + tmpl.List["header"] + tmpl.List["sidebar"] +
			tmpl.List["footer"] + tmpl.List["js"] + tmpl.List["menu"] +
			tmpl.List["admin_panel"] + tmpl.List["content"])
	} else {
		name = "content"
		tmpler, err = template.New("content").Parse(tmpl.List["admin_panel"] + tmpl.List["content"])
	}

	if err != nil {
		fmt.Println(err)
	}

	return
}

func (*Theme) GetAsset(path string) ([]byte, error) {
	return resource.Asset(path)
}

func (*Theme) GetAssetList() []string {
	return asserts
}

func (*Theme) Box() types.BoxAttribute {
	return &components.BoxAttribute{
		Name:       "box",
		Header:     template.HTML(""),
		Body:       template.HTML(""),
		Footer:     template.HTML(""),
		Title:      "",
		HeadBorder: "",
	}
}

func (*Theme) Col() types.ColAttribute {
	return &components.ColAttribute{
		Name:    "col",
		Size:    "col-md-2",
		Content: "",
	}
}

func (*Theme) Form() types.FormAttribute {
	return &components.FormAttribute{
		Name:    "form",
		Content: []types.FormStruct{},
		Url:     "/",
		Method:  "post",
		InfoUrl: "",
		Title:   "edit",
	}
}

func (*Theme) Image() types.ImgAttribute {
	return &components.ImgAttribute{
		Name:   "image",
		Witdh:  "50",
		Height: "50",
		Src:    "",
	}
}

func (*Theme) SmallBox() types.SmallBoxAttribute {
	return &components.SmallBoxAttribute{
		Name:  "smallbox",
		Title: "标题",
		Value: "值",
		Url:   "/",
		Color: "aqua",
	}
}

func (*Theme) InfoBox() types.InfoBoxAttribute {
	return &components.InfoBoxAttribute{
		Name:   "infobox",
		Text:   "标题",
		Icon:   "ion-ios-cart-outline",
		Number: "90",
		Color:  "red",
	}
}

func (*Theme) LineChart() types.LineChartAttribute {
	return &components.LineChartAttribute{
		Name: "line-chart",
	}
}

func (*Theme) ProgressGroup() types.ProgressGroupAttribute {
	return &components.ProgressGroupAttribute{
		Name: "progress-group",
	}
}

func (*Theme) Description() types.DescriptionAttribute {
	return &components.DescriptionAttribute{
		Name: "description",
	}
}

func (*Theme) Label() types.LabelAttribute {
	return &components.LabelAttribute{
		Name:    "label",
		Color:   "success",
		Content: "",
	}
}

func (*Theme) Paninator() types.PaninatorAttribute {
	return &components.PaninatorAttribute{
		Name: "paninator",
	}
}

func (*Theme) Row() types.RowAttribute {
	return &components.RowAttribute{
		Name:    "row",
		Content: "",
	}
}

func (*Theme) Table() types.TableAttribute {
	return &components.TableAttribute{
		Name:     "table",
		Thead:    []map[string]string{},
		InfoList: []map[string]template.HTML{},
		Type:     "normal",
	}
}

func (theme *Theme) DataTable() types.DataTableAttribute {
	return &components.DataTableAttribute{
		TableAttribute: *(theme.Table().SetType("data-table").(*components.TableAttribute)),
		EditUrl:        "",
		NewUrl:         "",
	}
}

func (*Theme) Tree() types.TreeAttribute {
	return &components.TreeAttribute{
		Name: "tree",
		Tree: []menu.MenuItem{},
	}
}
