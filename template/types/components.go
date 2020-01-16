// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/template/types/form"
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
	SetLayout(layout form.Layout) FormAttribute
	SetToken(value string) FormAttribute
	SetOperationFooter(value template.HTML) FormAttribute
	GetBoxHeader() template.HTML
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
	SetSize(value map[string]string) ColAttribute
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
	SetButtons(btns template.HTML) DataTableAttribute
	SetHideFilterArea(value bool) DataTableAttribute
	SetHideRowSelector(value bool) DataTableAttribute
	SetActionJs(aj template.JS) DataTableAttribute
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
