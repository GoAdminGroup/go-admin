// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"html/template"
)

type FormAttribute interface {
	SetHeader(value template.HTML) FormAttribute
	SetContent(value []FormField) FormAttribute
	SetTabContents(value [][]FormField) FormAttribute
	SetTabHeaders(value []string) FormAttribute
	SetFooter(value template.HTML) FormAttribute
	SetPrefix(value string) FormAttribute
	SetUrl(value string) FormAttribute
	SetPrimaryKey(value string) FormAttribute
	SetInfoUrl(value string) FormAttribute
	SetMethod(value string) FormAttribute
	SetTitle(value template.HTML) FormAttribute
	SetToken(value string) FormAttribute
	GetContent() template.HTML
}

type BoxAttribute interface {
	SetHeader(value template.HTML) BoxAttribute
	SetBody(value template.HTML) BoxAttribute
	SetFooter(value template.HTML) BoxAttribute
	SetTitle(value template.HTML) BoxAttribute
	WithHeadBorder(has bool) BoxAttribute
	SetHeadColor(value string) BoxAttribute
	SetTheme(value string) BoxAttribute
	GetContent() template.HTML
}

type ColAttribute interface {
	SetSize(value map[string]string) ColAttribute
	SetContent(value template.HTML) ColAttribute
	AddContent(value template.HTML) ColAttribute
	GetContent() template.HTML
}

type ImgAttribute interface {
	SetWidth(value string) ImgAttribute
	SetHeight(value string) ImgAttribute
	SetSrc(value string) ImgAttribute
	GetContent() template.HTML
}

type SmallBoxAttribute interface {
	SetTitle(value template.HTML) SmallBoxAttribute
	SetValue(value template.HTML) SmallBoxAttribute
	SetColor(value template.HTML) SmallBoxAttribute
	SetIcon(value template.HTML) SmallBoxAttribute
	SetUrl(value string) SmallBoxAttribute
	GetContent() template.HTML
}

type LabelAttribute interface {
	SetContent(value template.HTML) LabelAttribute
	GetContent() template.HTML
}

type RowAttribute interface {
	SetContent(value template.HTML) RowAttribute
	AddContent(value template.HTML) RowAttribute
	GetContent() template.HTML
}

type TableAttribute interface {
	SetThead(value []map[string]string) TableAttribute
	SetInfoList(value []map[string]template.HTML) TableAttribute
	SetType(value string) TableAttribute
	SetMinWidth(value int) TableAttribute
	GetContent() template.HTML
}

type DataTableAttribute interface {
	GetDataTableHeader() template.HTML
	SetThead(value []map[string]string) DataTableAttribute
	SetInfoList(value []map[string]template.HTML) DataTableAttribute
	SetEditUrl(value string) DataTableAttribute
	SetDeleteUrl(value string) DataTableAttribute
	SetNewUrl(value string) DataTableAttribute
	SetPrimaryKey(value string) DataTableAttribute
	SetAction(action template.HTML) DataTableAttribute
	SetIsTab(value bool) DataTableAttribute
	SetFilterUrl(value string) DataTableAttribute
	SetInfoUrl(value string) DataTableAttribute
	SetExportUrl(value string) DataTableAttribute
	SetFilters(value []map[string]string) DataTableAttribute
	GetContent() template.HTML
}

type TreeAttribute interface {
	SetTree(value []menu.Item) TreeAttribute
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
	SetIcon(value template.HTML) InfoBoxAttribute
	SetText(value template.HTML) InfoBoxAttribute
	SetNumber(value template.HTML) InfoBoxAttribute
	SetContent(value template.HTML) InfoBoxAttribute
	SetColor(value template.HTML) InfoBoxAttribute
	GetContent() template.HTML
}

type ProgressGroupAttribute interface {
	SetTitle(value template.HTML) ProgressGroupAttribute
	SetColor(value template.HTML) ProgressGroupAttribute
	SetPercent(value int) ProgressGroupAttribute
	SetDenominator(value int) ProgressGroupAttribute
	SetMolecular(value int) ProgressGroupAttribute
	GetContent() template.HTML
}

type ProgressAttribute interface{}

type LineChartAttribute interface {
	SetID(value string) LineChartAttribute
	SetTitle(value template.HTML) LineChartAttribute
	SetHeight(value int) LineChartAttribute
	SetData(value string) LineChartAttribute
	GetContent() template.HTML
}

type BarChartAttribute interface {
	SetID(value string) BarChartAttribute
	SetTitle(value template.HTML) BarChartAttribute
	SetWidth(value int) BarChartAttribute
	SetData(value string) BarChartAttribute
	GetContent() template.HTML
}

type PieChartAttribute interface {
	SetID(value string) PieChartAttribute
	SetData(value string) PieChartAttribute
	SetTitle(value template.HTML) PieChartAttribute
	SetHeight(value int) PieChartAttribute
	GetContent() template.HTML
}

type ChartLegendAttribute interface {
	SetData(value []map[string]string) ChartLegendAttribute
	GetContent() template.HTML
}

type DescriptionAttribute interface {
	SetNumber(value template.HTML) DescriptionAttribute
	SetTitle(value template.HTML) DescriptionAttribute
	SetArrow(value string) DescriptionAttribute
	SetPercent(value template.HTML) DescriptionAttribute
	SetBorder(value string) DescriptionAttribute
	SetColor(value template.HTML) DescriptionAttribute
	GetContent() template.HTML
}

type AreaChartAttribute interface {
	SetTitle(value template.HTML) AreaChartAttribute
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
	SetTitle(value template.HTML) PopupAttribute
	SetFooter(value template.HTML) PopupAttribute
	SetBody(value template.HTML) PopupAttribute
	SetSize(value string) PopupAttribute
	GetContent() template.HTML
}
