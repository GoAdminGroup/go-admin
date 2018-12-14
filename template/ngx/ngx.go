package ngx

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type Theme struct {
	Name string
}

func (*Theme) Form() types.FormAttribute {
	panic("implement me")
}

func (*Theme) Box() types.BoxAttribute {
	panic("implement me")
}

func (*Theme) Col() types.ColAttribute {
	panic("implement me")
}

func (*Theme) Image() types.ImgAttribute {
	panic("implement me")
}

func (*Theme) SmallBox() types.SmallBoxAttribute {
	panic("implement me")
}

func (*Theme) Label() types.LabelAttribute {
	panic("implement me")
}

func (*Theme) Row() types.RowAttribute {
	panic("implement me")
}

func (*Theme) Table() types.TableAttribute {
	panic("implement me")
}

func (*Theme) DataTable() types.DataTableAttribute {
	panic("implement me")
}

func (*Theme) Tree() types.TreeAttribute {
	panic("implement me")
}

func (*Theme) InfoBox() types.InfoBoxAttribute {
	panic("implement me")
}

func (*Theme) Paginator() types.PaginatorAttribute {
	panic("implement me")
}

func (*Theme) AreaChart() types.AreaChartAttribute {
	panic("implement me")
}

func (*Theme) ProgressGroup() types.ProgressGroupAttribute {
	panic("implement me")
}

func (*Theme) LineChart() types.LineChartAttribute {
	panic("implement me")
}

func (*Theme) BarChart() types.BarChartAttribute {
	panic("implement me")
}

func (*Theme) ProductList() types.ProductListAttribute {
	panic("implement me")
}

func (*Theme) Description() types.DescriptionAttribute {
	panic("implement me")
}

func (*Theme) Alert() types.AlertAttribute {
	panic("implement me")
}

func (*Theme) PieChart() types.PieChartAttribute {
	panic("implement me")
}

func (*Theme) ChartLegend() types.ChartLegendAttribute {
	panic("implement me")
}

func (*Theme) Tabs() types.TabsAttribute {
	panic("implement me")
}

func (*Theme) Popup() types.PopupAttribute {
	panic("implement me")
}

func (*Theme) GetTmplList() map[string]string {
	panic("implement me")
}

func (*Theme) GetAssetList() []string {
	panic("implement me")
}

func (*Theme) GetAsset(string) ([]byte, error) {
	panic("implement me")
}

func (*Theme) GetTemplate(bool) (*template.Template, string) {
	panic("implement me")
}

var Ngx = Theme{
	Name: "ngx",
}
