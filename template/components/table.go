package components

import (
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
)

type TableAttribute struct {
	Name       string
	Thead      types.Thead
	InfoList   []map[string]types.InfoItem
	Type       string
	PrimaryKey string
	NoAction   bool
	Action     template.HTML
	EditUrl    string
	MinWidth   int
	DeleteUrl  string
	DetailUrl  string
	UpdateUrl  string
	Layout     string
	IsTab      bool
	ExportUrl  string
	types.Attribute
}

func (compo *TableAttribute) SetThead(value types.Thead) types.TableAttribute {
	compo.Thead = value
	return compo
}

func (compo *TableAttribute) SetInfoList(value []map[string]types.InfoItem) types.TableAttribute {
	compo.InfoList = value
	return compo
}

func (compo *TableAttribute) SetType(value string) types.TableAttribute {
	compo.Type = value
	return compo
}

func (compo *TableAttribute) SetMinWidth(value int) types.TableAttribute {
	compo.MinWidth = value
	return compo
}

func (compo *TableAttribute) SetLayout(value string) types.TableAttribute {
	compo.Layout = value
	return compo
}

func (compo *TableAttribute) GetContent() template.HTML {
	if compo.MinWidth == 0 {
		compo.MinWidth = 1000
	}
	return ComposeHtml(compo.TemplateList, *compo, "table")
}

type DataTableAttribute struct {
	TableAttribute
	EditUrl           string
	NewUrl            string
	UpdateUrl         string
	DetailUrl         string
	DeleteUrl         string
	PrimaryKey        string
	IsTab             bool
	ExportUrl         string
	InfoUrl           string
	Buttons           template.HTML
	ActionJs          template.JS
	IsHideFilterArea  bool
	IsHideRowSelector bool
	NoAction          bool
	HasFilter         bool
	Action            template.HTML
	types.Attribute
}

func (compo *DataTableAttribute) GetDataTableHeader() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "table/box-header")
}

func (compo *DataTableAttribute) SetThead(value types.Thead) types.DataTableAttribute {
	compo.Thead = value
	return compo
}

func (compo *DataTableAttribute) SetLayout(value string) types.DataTableAttribute {
	compo.Layout = value
	return compo
}

func (compo *DataTableAttribute) SetIsTab(value bool) types.DataTableAttribute {
	compo.IsTab = value
	return compo
}

func (compo *DataTableAttribute) SetButtons(btns template.HTML) types.DataTableAttribute {
	compo.Buttons = btns
	return compo
}

func (compo *DataTableAttribute) SetHideFilterArea(value bool) types.DataTableAttribute {
	compo.IsHideFilterArea = value
	return compo
}

func (compo *DataTableAttribute) SetActionJs(aj template.JS) types.DataTableAttribute {
	compo.ActionJs = aj
	return compo
}

func (compo *DataTableAttribute) SetHasFilter(hasFilter bool) types.DataTableAttribute {
	compo.HasFilter = hasFilter
	return compo
}

func (compo *DataTableAttribute) SetInfoUrl(value string) types.DataTableAttribute {
	compo.InfoUrl = value
	return compo
}

func (compo *DataTableAttribute) SetAction(action template.HTML) types.DataTableAttribute {
	compo.Action = action
	return compo
}

func (compo *DataTableAttribute) SetExportUrl(value string) types.DataTableAttribute {
	compo.ExportUrl = value
	return compo
}

func (compo *DataTableAttribute) SetHideRowSelector(value bool) types.DataTableAttribute {
	compo.IsHideRowSelector = value
	return compo
}

func (compo *DataTableAttribute) SetUpdateUrl(value string) types.DataTableAttribute {
	compo.UpdateUrl = value
	return compo
}

func (compo *DataTableAttribute) SetDetailUrl(value string) types.DataTableAttribute {
	compo.DetailUrl = value
	return compo
}

func (compo *DataTableAttribute) SetPrimaryKey(value string) types.DataTableAttribute {
	compo.PrimaryKey = value
	return compo
}

func (compo *DataTableAttribute) SetInfoList(value []map[string]types.InfoItem) types.DataTableAttribute {
	compo.InfoList = value
	return compo
}

func (compo *DataTableAttribute) SetEditUrl(value string) types.DataTableAttribute {
	compo.EditUrl = value
	return compo
}

func (compo *DataTableAttribute) SetDeleteUrl(value string) types.DataTableAttribute {
	compo.DeleteUrl = value
	return compo
}

func (compo *DataTableAttribute) SetNewUrl(value string) types.DataTableAttribute {
	compo.NewUrl = value
	return compo
}

func (compo *DataTableAttribute) SetNoAction() types.DataTableAttribute {
	compo.NoAction = true
	return compo
}

func (compo *DataTableAttribute) GetContent() template.HTML {
	if compo.MinWidth == 0 {
		compo.MinWidth = 1000
	}
	if !compo.NoAction && compo.EditUrl == "" && compo.DeleteUrl == "" && compo.DetailUrl == "" && compo.Action == "" {
		compo.NoAction = true
	}
	return ComposeHtml(compo.TemplateList, *compo, "table")
}
