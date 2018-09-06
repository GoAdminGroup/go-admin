package types

import (
	"html/template"
	"goAdmin/modules/menu"
)

type FormAttribute interface {
	SetContent(value []FormStruct) FormAttribute
	SetPrefix(value string) FormAttribute
	SetUrl(value string) FormAttribute
	SetInfoUrl(value string) FormAttribute
	SetMethod(value string) FormAttribute
	SetTitle(value string) FormAttribute
	SetToken(value string) FormAttribute
	GetContent() template.HTML
}

type BoxAttribute interface {
	SetHeader(value template.HTML) BoxAttribute
	SetBody(value template.HTML) BoxAttribute
	SetFooter(value template.HTML) BoxAttribute
	SetTitle(value template.HTML) BoxAttribute
	WithHeadBorder(has bool) BoxAttribute
	GetContent() template.HTML
}

type ColAttribute interface {
	SetSize(value map[string]string) ColAttribute
	SetContent(value template.HTML) ColAttribute
	GetContent() template.HTML
}

type ImgAttribute interface {
	SetWidth(value string) ImgAttribute
	SetHeight(value string) ImgAttribute
	SetSrc(value string) ImgAttribute
	GetContent() template.HTML
}

type SmallBoxAttribute interface {
	SetTitle(value string) SmallBoxAttribute
	SetValue(value string) SmallBoxAttribute
	SetUrl(value string) SmallBoxAttribute
	GetContent() template.HTML
}

type LabelAttribute interface {
	SetContent(value string) LabelAttribute
	GetContent() template.HTML
}

type RowAttribute interface {
	SetContent(value template.HTML) RowAttribute
	GetContent() template.HTML
}

type TableAttribute interface {
	SetThead(value []map[string]string) TableAttribute
	SetInfoList(value []map[string]template.HTML) TableAttribute
	SetType(value string) TableAttribute
	GetContent() template.HTML
}

type DataTableAttribute interface {
	GetDataTableHeader() template.HTML
	SetThead(value []map[string]string) DataTableAttribute
	SetInfoList(value []map[string]template.HTML) DataTableAttribute
	SetEditUrl(value string) DataTableAttribute
	SetNewUrl(value string) DataTableAttribute
	GetContent() template.HTML
}

type TreeAttribute interface {
	SetTree(value []menu.MenuItem) TreeAttribute
	SetEditUrl(value string) TreeAttribute
	SetOrderUrl(value string) TreeAttribute
	SetDeleteUrl(value string) TreeAttribute
	GetContent() template.HTML
	GetTreeHeader() template.HTML
}

type PaninatorAttribute interface {
	SetCurPageStartIndex(value string) PaninatorAttribute
	SetCurPageEndIndex(value string) PaninatorAttribute
	SetTotal(value string) PaninatorAttribute
	SetPreviousClass(value string) PaninatorAttribute
	SetPreviousUrl(value string) PaninatorAttribute
	SetPages(value []map[string]string) PaninatorAttribute
	SetNextClass(value string) PaninatorAttribute
	SetNextUrl(value string) PaninatorAttribute
	SetOption(value map[string]template.HTML) PaninatorAttribute
	SetUrl(value string) PaninatorAttribute
	GetContent() template.HTML
}

type InfoBoxAttribute interface {
	SetIcon(value string) InfoBoxAttribute
	SetText(value string) InfoBoxAttribute
	SetNumber(value template.HTML) InfoBoxAttribute
	SetContent(value string) InfoBoxAttribute
	SetColor(value string) InfoBoxAttribute
	GetContent() template.HTML
}

type ProgressGroupAttribute interface {
	SetTitle(value string) ProgressGroupAttribute
	SetColor(value string) ProgressGroupAttribute
	SetPercent(value int) ProgressGroupAttribute
	SetDenominator(value int) ProgressGroupAttribute
	SetMolecular(value int) ProgressGroupAttribute
	GetContent() template.HTML
}

type ProgressAttribute interface{}

type DescriptionAttribute interface {
	SetNumber(value string) DescriptionAttribute
	SetTitle(value string) DescriptionAttribute
	SetArrow(value string) DescriptionAttribute
	SetPercent(value string) DescriptionAttribute
	SetBorder(value string) DescriptionAttribute
	SetColor(value string) DescriptionAttribute
	GetContent() template.HTML
}

type LineChartAttribute interface {
	SetTitle(value string) LineChartAttribute
	SetPrefix(value string) LineChartAttribute
	SetID(value string) LineChartAttribute
	SetData(value string) LineChartAttribute
	SetHeight(value int) LineChartAttribute
	GetContent() template.HTML
}

type ProductListAttribute interface{}
