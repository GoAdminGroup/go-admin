// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/chenhg5/go-admin/modules/menu"
	"html/template"
)

type FormAttribute interface {
	SetContent(value []Form) FormAttribute
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
	SetTheme(value string) BoxAttribute
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
	SetDeleteUrl(value string) DataTableAttribute
	SetNewUrl(value string) DataTableAttribute
	SetFilterUrl(value string) DataTableAttribute
	SetInfoUrl(value string) DataTableAttribute
	SetFilters(value []map[string]string) DataTableAttribute
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

type PaginatorAttribute interface {
	SetCurPageStartIndex(value string) PaginatorAttribute
	SetCurPageEndIndex(value string) PaginatorAttribute
	SetTotal(value string) PaginatorAttribute
	SetPreviousClass(value string) PaginatorAttribute
	SetPreviousUrl(value string) PaginatorAttribute
	SetPages(value []map[string]string) PaginatorAttribute
	SetNextClass(value string) PaginatorAttribute
	SetNextUrl(value string) PaginatorAttribute
	SetOption(value map[string]template.HTML) PaginatorAttribute
	SetUrl(value string) PaginatorAttribute
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

type LineChartAttribute interface {
	SetID(value string) LineChartAttribute
	SetTitle(value string) LineChartAttribute
	SetHeight(value int) LineChartAttribute
	SetData(value string) LineChartAttribute
	GetContent() template.HTML
}

type BarChartAttribute interface {
	SetID(value string) BarChartAttribute
	SetTitle(value string) BarChartAttribute
	SetWidth(value int) BarChartAttribute
	SetData(value string) BarChartAttribute
	GetContent() template.HTML
}

type PieChartAttribute interface {
	SetID(value string) PieChartAttribute
	SetData(value string) PieChartAttribute
	SetTitle(value string) PieChartAttribute
	SetHeight(value int) PieChartAttribute
	GetContent() template.HTML
}

type ChartLegendAttribute interface {
	SetData(value []map[string]string) ChartLegendAttribute
	GetContent() template.HTML
}

type DescriptionAttribute interface {
	SetNumber(value string) DescriptionAttribute
	SetTitle(value string) DescriptionAttribute
	SetArrow(value string) DescriptionAttribute
	SetPercent(value string) DescriptionAttribute
	SetBorder(value string) DescriptionAttribute
	SetColor(value string) DescriptionAttribute
	GetContent() template.HTML
}

type AreaChartAttribute interface {
	SetTitle(value string) AreaChartAttribute
	SetID(value string) AreaChartAttribute
	SetData(value string) AreaChartAttribute
	SetHeight(value int) AreaChartAttribute
	GetContent() template.HTML
}

type ProductListAttribute interface {
	SetData(value []map[string]string) ProductListAttribute
	GetContent() template.HTML
}

type TabsAttribute interface {
	SetData(value []map[string]template.HTML) TabsAttribute
	GetContent() template.HTML
}

type AlertAttribute interface {
	SetTheme(value string) AlertAttribute
	SetTitle(value template.HTML) AlertAttribute
	SetContent(value template.HTML) AlertAttribute
	GetContent() template.HTML
}

type PopupAttribute interface {
	SetID(value string) PopupAttribute
	SetTitle(value string) PopupAttribute
	SetFooter(value string) PopupAttribute
	SetBody(value template.HTML) PopupAttribute
	GetContent() template.HTML
}
