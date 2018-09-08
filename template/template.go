package template

import (
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/adminlte"
	"html/template"
	"github.com/chenhg5/go-admin/template/login"
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
	LineChart() types.LineChartAttribute
	BarChart() types.BarChartAttribute
	ProductList() types.ProductListAttribute
	Description() types.DescriptionAttribute
	Alert() types.AlertAttribute
	PieChart() types.PieChartAttribute
	ChartLegend() types.ChartLegendAttribute
	Tabs() types.TabsAttribute
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