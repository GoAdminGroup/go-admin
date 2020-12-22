package components

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type PaginatorAttribute struct {
	Name              string
	CurPageStartIndex string
	CurPageEndIndex   string
	Total             string
	PreviousClass     string
	PreviousUrl       string
	Pages             []map[string]string
	NextClass         string
	NextUrl           string
	PageSizeList      []string
	Option            map[string]template.HTML
	Url               string
	ExtraInfo         template.HTML
	EntriesInfo       template.HTML
	types.Attribute
}

func (compo *PaginatorAttribute) SetCurPageStartIndex(value string) types.PaginatorAttribute {
	compo.CurPageStartIndex = value
	return compo
}

func (compo *PaginatorAttribute) SetCurPageEndIndex(value string) types.PaginatorAttribute {
	compo.CurPageEndIndex = value
	return compo
}

func (compo *PaginatorAttribute) SetTotal(value string) types.PaginatorAttribute {
	compo.Total = value
	return compo
}

func (compo *PaginatorAttribute) SetExtraInfo(value template.HTML) types.PaginatorAttribute {
	compo.ExtraInfo = value
	return compo
}

func (compo *PaginatorAttribute) SetEntriesInfo(value template.HTML) types.PaginatorAttribute {
	compo.EntriesInfo = value
	return compo
}

func (compo *PaginatorAttribute) SetPreviousClass(value string) types.PaginatorAttribute {
	compo.PreviousClass = value
	return compo
}

func (compo *PaginatorAttribute) SetPreviousUrl(value string) types.PaginatorAttribute {
	compo.PreviousUrl = value
	return compo
}

func (compo *PaginatorAttribute) SetPages(value []map[string]string) types.PaginatorAttribute {
	compo.Pages = value
	return compo
}

func (compo *PaginatorAttribute) SetPageSizeList(value []string) types.PaginatorAttribute {
	compo.PageSizeList = value
	return compo
}

func (compo *PaginatorAttribute) SetNextClass(value string) types.PaginatorAttribute {
	compo.NextClass = value
	return compo
}

func (compo *PaginatorAttribute) SetNextUrl(value string) types.PaginatorAttribute {
	compo.NextUrl = value
	return compo
}

func (compo *PaginatorAttribute) SetOption(value map[string]template.HTML) types.PaginatorAttribute {
	compo.Option = value
	return compo
}

func (compo *PaginatorAttribute) SetUrl(value string) types.PaginatorAttribute {
	compo.Url = value
	return compo
}

func (compo *PaginatorAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "paginator")
}
