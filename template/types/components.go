// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"html/template"
)

type FormAttribute interface {
	SetHeader(value template.HTML) FormAttribute
	SetContent(value FormFields) FormAttribute
	SetTabContents(value []FormFields) FormAttribute
	SetTabHeaders(value []string) FormAttribute
	SetFooter(value template.HTML) FormAttribute
	SetPrefix(value string) FormAttribute
	SetUrl(value string) FormAttribute
	SetPrimaryKey(value string) FormAttribute
	SetHiddenFields(fields map[string]string) FormAttribute
	SetMethod(value string) FormAttribute
	SetHeadWidth(width int) FormAttribute
	SetInputWidth(width int) FormAttribute
	SetTitle(value template.HTML) FormAttribute
	SetLayout(layout form.Layout) FormAttribute
	SetOperationFooter(value template.HTML) FormAttribute
	GetDefaultBoxHeader() template.HTML
	GetDetailBoxHeader(editUrl, deleteUrl string) template.HTML
	GetBoxHeaderNoButton() template.HTML
	GetContent() template.HTML
}

type BoxAttribute interface {
	SetHeader(value template.HTML) BoxAttribute
	SetBody(value template.HTML) BoxAttribute
	SetNoPadding() BoxAttribute
	SetFooter(value template.HTML) BoxAttribute
	SetTitle(value template.HTML) BoxAttribute
	WithHeadBorder() BoxAttribute
	SetStyle(value template.HTMLAttr) BoxAttribute
	SetHeadColor(value string) BoxAttribute
	SetTheme(value string) BoxAttribute
	SetSecondHeader(value template.HTML) BoxAttribute
	SetSecondHeadColor(value string) BoxAttribute
	WithSecondHeadBorder() BoxAttribute
	SetSecondHeaderClass(value string) BoxAttribute
	GetContent() template.HTML
}

type ColAttribute interface {
	SetSize(value S) ColAttribute
	SetContent(value template.HTML) ColAttribute
	AddContent(value template.HTML) ColAttribute
	GetContent() template.HTML
}

type ImgAttribute interface {
	SetWidth(value string) ImgAttribute
	SetHeight(value string) ImgAttribute
	WithModal() ImgAttribute
	SetSrc(value template.HTML) ImgAttribute
	GetContent() template.HTML
}

type LabelAttribute interface {
	SetContent(value template.HTML) LabelAttribute
	SetColor(value template.HTML) LabelAttribute
	SetType(value string) LabelAttribute
	GetContent() template.HTML
}

type RowAttribute interface {
	SetContent(value template.HTML) RowAttribute
	AddContent(value template.HTML) RowAttribute
	GetContent() template.HTML
}

type ButtonAttribute interface {
	SetContent(value template.HTML) ButtonAttribute
	SetOrientationRight() ButtonAttribute
	SetOrientationLeft() ButtonAttribute
	SetMarginLeft(int) ButtonAttribute
	SetMarginRight(int) ButtonAttribute
	SetThemePrimary() ButtonAttribute
	SetSmallSize() ButtonAttribute
	SetMiddleSize() ButtonAttribute
	SetHref(string) ButtonAttribute
	SetThemeWarning() ButtonAttribute
	SetTheme(value string) ButtonAttribute
	SetLoadingText(value template.HTML) ButtonAttribute
	SetThemeDefault() ButtonAttribute
	SetType(string) ButtonAttribute
	GetContent() template.HTML
}

type TableAttribute interface {
	SetThead(value Thead) TableAttribute
	SetInfoList(value []map[string]InfoItem) TableAttribute
	SetType(value string) TableAttribute
	SetMinWidth(value int) TableAttribute
	SetLayout(value string) TableAttribute
	GetContent() template.HTML
}

type DataTableAttribute interface {
	GetDataTableHeader() template.HTML
	SetThead(value Thead) DataTableAttribute
	SetInfoList(value []map[string]InfoItem) DataTableAttribute
	SetEditUrl(value string) DataTableAttribute
	SetDeleteUrl(value string) DataTableAttribute
	SetNewUrl(value string) DataTableAttribute
	SetPrimaryKey(value string) DataTableAttribute
	SetAction(action template.HTML) DataTableAttribute
	SetIsTab(value bool) DataTableAttribute
	SetLayout(value string) DataTableAttribute
	SetButtons(btns template.HTML) DataTableAttribute
	SetHideFilterArea(value bool) DataTableAttribute
	SetHideRowSelector(value bool) DataTableAttribute
	SetActionJs(aj template.JS) DataTableAttribute
	SetNoAction() DataTableAttribute
	SetInfoUrl(value string) DataTableAttribute
	SetDetailUrl(value string) DataTableAttribute
	SetHasFilter(hasFilter bool) DataTableAttribute
	SetExportUrl(value string) DataTableAttribute
	SetUpdateUrl(value string) DataTableAttribute
	GetContent() template.HTML
}

type TreeAttribute interface {
	SetTree(value []menu.Item) TreeAttribute
	SetEditUrl(value string) TreeAttribute
	SetOrderUrl(value string) TreeAttribute
	SetUrlPrefix(value string) TreeAttribute
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
	SetPageSizeList(value []string) PaginatorAttribute
	SetNextClass(value string) PaginatorAttribute
	SetNextUrl(value string) PaginatorAttribute
	SetOption(value map[string]template.HTML) PaginatorAttribute
	SetUrl(value string) PaginatorAttribute
	SetExtraInfo(value template.HTML) PaginatorAttribute
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
	Warning(msg string) template.HTML
	GetContent() template.HTML
}

type LinkAttribute interface {
	OpenInNewTab() LinkAttribute
	SetURL(value string) LinkAttribute
	SetAttributes(attr template.HTMLAttr) LinkAttribute
	SetClass(class template.HTML) LinkAttribute
	NoPjax() LinkAttribute
	SetTabTitle(value template.HTML) LinkAttribute
	SetContent(value template.HTML) LinkAttribute
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

type Thead []TheadItem

type TheadItem struct {
	Head       string       `json:"head"`
	Sortable   bool         `json:"sortable"`
	Field      string       `json:"field"`
	Hide       bool         `json:"hide"`
	Editable   bool         `json:"editable"`
	EditType   string       `json:"edit_type"`
	EditOption FieldOptions `json:"edit_option"`
	Width      int          `json:"width"`
}

func (t Thead) GroupBy(group [][]string) []Thead {
	var res = make([]Thead, len(group))

	for key, value := range group {
		var newThead = make(Thead, len(t))

		for index, info := range t {
			if modules.InArray(value, info.Field) {
				newThead[index] = info
			}
		}

		res[key] = newThead
	}

	return res
}
