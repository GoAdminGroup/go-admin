package template

import (
	"goAdmin/template/types"
	"goAdmin/template/adminlte"
	"html/template"
	"goAdmin/template/login"
)

type Template interface {
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
	Paninator() types.PaninatorAttribute
	AreaChart() types.AreaChartAttribute
	ProgressGroup() types.ProgressGroupAttribute
	ProductList() types.ProductListAttribute
	Description() types.DescriptionAttribute
	PieChart() types.PieChartAttribute
	ChartLegend() types.ChartLegendAttribute
	GetTmplList() map[string]string
	GetAssetList() []string
	GetAsset(string) ([]byte, error)
	GetTemplate(bool) (*template.Template, string)
}

func Get(theme string) Template {
	switch theme {
	case "adminlte":
		return adminlte.GetAdminlte()
	default:
		panic("wrong theme name")
	}
}

type Component interface {
	GetTemplate() (*template.Template, string)
	GetAssetList() []string
	GetAsset(string) ([]byte, error)
}

func GetComp(name string) Component {
	switch name {
	case "login":
		return login.GetLoginComponent()
	default:
		panic("wrong component name")
	}
}