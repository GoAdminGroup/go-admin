// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package types

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/template/types/form"
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
	SetId(id string) FormAttribute
	SetAjax(successJS, errorJS template.JS) FormAttribute
	SetHiddenFields(fields map[string]string) FormAttribute
	SetFieldsHTML(html template.HTML) FormAttribute
	SetMethod(value string) FormAttribute
	SetHeadWidth(width int) FormAttribute
	SetInputWidth(width int) FormAttribute
	SetTitle(value template.HTML) FormAttribute
	SetLayout(layout form.Layout) FormAttribute
	SetOperationFooter(value template.HTML) FormAttribute
	GetDefaultBoxHeader(hideBack bool) template.HTML
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
	SetIframeStyle(iframe bool) BoxAttribute
	SetAttr(attr template.HTMLAttr) BoxAttribute
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
	AddClass(class string) ButtonAttribute
	SetID(id string) ButtonAttribute
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
	SetName(name string) TableAttribute
	SetMinWidth(value string) TableAttribute
	SetHideThead() TableAttribute
	SetLayout(value string) TableAttribute
	SetStyle(style string) TableAttribute
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
	SetStyle(style string) DataTableAttribute
	SetAction(action template.HTML) DataTableAttribute
	SetIsTab(value bool) DataTableAttribute
	SetActionFold(fold bool) DataTableAttribute
	SetHideThead() DataTableAttribute
	SetLayout(value string) DataTableAttribute
	SetButtons(btns template.HTML) DataTableAttribute
	SetHideFilterArea(value bool) DataTableAttribute
	SetHideRowSelector(value bool) DataTableAttribute
	SetActionJs(aj template.JS) DataTableAttribute
	SetNoAction() DataTableAttribute
	SetInfoUrl(value string) DataTableAttribute
	SetDetailUrl(value string) DataTableAttribute
	SetHasFilter(hasFilter bool) DataTableAttribute
	SetSortUrl(value string) DataTableAttribute
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

type TreeViewAttribute interface {
	SetTree(value TreeViewData) TreeViewAttribute
	SetUrlPrefix(value string) TreeViewAttribute
	SetID(id string) TreeViewAttribute
	GetContent() template.HTML
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
	SetEntriesInfo(value template.HTML) PaginatorAttribute
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
	SetDraggable() PopupAttribute
	SetHideFooter() PopupAttribute
	SetWidth(width string) PopupAttribute
	SetHeight(height string) PopupAttribute
	SetFooter(value template.HTML) PopupAttribute
	SetFooterHTML(value template.HTML) PopupAttribute
	SetBody(value template.HTML) PopupAttribute
	SetSize(value string) PopupAttribute
	GetContent() template.HTML
}

type PanelInfo struct {
	Thead    Thead    `json:"thead"`
	InfoList InfoList `json:"info_list"`
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
	Width      string       `json:"width"`
}

func (t Thead) GroupBy(group [][]string) []Thead {
	var res = make([]Thead, len(group))

	for key, value := range group {
		var newThead = make(Thead, 0)

		for _, info := range t {
			if modules.InArray(value, info.Field) {
				newThead = append(newThead, info)
			}
		}

		res[key] = newThead
	}

	return res
}

type TreeViewData struct {
	Data              TreeViewItems `json:"data,omitempty"`
	Levels            int           `json:"levels,omitempty"`
	BackColor         string        `json:"backColor,omitempty"`
	BorderColor       string        `json:"borderColor,omitempty"`
	CheckedIcon       string        `json:"checkedIcon,omitempty"`
	CollapseIcon      string        `json:"collapseIcon,omitempty"`
	Color             string        `json:"color,omitempty"`
	EmptyIcon         string        `json:"emptyIcon,omitempty"`
	EnableLinks       bool          `json:"enableLinks,omitempty"`
	ExpandIcon        string        `json:"expandIcon,omitempty"`
	MultiSelect       bool          `json:"multiSelect,omitempty"`
	NodeIcon          string        `json:"nodeIcon,omitempty"`
	OnhoverColor      string        `json:"onhoverColor,omitempty"`
	SelectedIcon      string        `json:"selectedIcon,omitempty"`
	SearchResultColor string        `json:"searchResultColor,omitempty"`
	SelectedBackColor string        `json:"selectedBackColor,omitempty"`
	SelectedColor     string        `json:"selectedColor,omitempty"`
	ShowBorder        bool          `json:"showBorder,omitempty"`
	ShowCheckbox      bool          `json:"showCheckbox,omitempty"`
	ShowIcon          bool          `json:"showIcon,omitempty"`
	ShowTags          bool          `json:"showTags,omitempty"`
	UncheckedIcon     string        `json:"uncheckedIcon,omitempty"`

	SearchResultBackColor  string `json:"searchResultBackColor,omitempty"`
	HighlightSearchResults bool   `json:"highlightSearchResults,omitempty"`
}

type TreeViewItems []TreeViewItem

type TreeViewItemState struct {
	Checked  bool `json:"checked,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Expanded bool `json:"expanded,omitempty"`
	Selected bool `json:"selected,omitempty"`
}

type TreeViewItem struct {
	Text         string            `json:"text,omitempty"`
	Icon         string            `json:"icon,omitempty"`
	SelectedIcon string            `json:"selected_icon,omitempty"`
	Color        string            `json:"color,omitempty"`
	BackColor    string            `json:"backColor,omitempty"`
	Href         string            `json:"href,omitempty"`
	Selectable   bool              `json:"selectable,omitempty"`
	State        TreeViewItemState `json:"state,omitempty"`
	Tags         []string          `json:"tags,omitempty"`
	Nodes        TreeViewItems     `json:"nodes,omitempty"`
}
