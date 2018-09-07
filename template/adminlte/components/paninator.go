package components

import (
	"html/template"
	"goAdmin/template/types"
)

type PaninatorAttribute struct {
	Name              string
	CurPageStartIndex string
	CurPageEndIndex   string
	Total             string
	PreviousClass     string
	PreviousUrl       string
	Pages             []map[string]string
	NextClass         string
	NextUrl           string
	Option            map[string]template.HTML
	Url               string
}

func (compo *PaninatorAttribute) SetCurPageStartIndex(value string) types.PaninatorAttribute {
	(*compo).CurPageStartIndex = value
	return compo
}

func (compo *PaninatorAttribute) SetCurPageEndIndex(value string) types.PaninatorAttribute {
	(*compo).CurPageEndIndex = value
	return compo
}

func (compo *PaninatorAttribute) SetTotal(value string) types.PaninatorAttribute {
	(*compo).Total = value
	return compo
}

func (compo *PaninatorAttribute) SetPreviousClass(value string) types.PaninatorAttribute {
	(*compo).PreviousClass = value
	return compo
}

func (compo *PaninatorAttribute) SetPreviousUrl(value string) types.PaninatorAttribute {
	(*compo).PreviousUrl = value
	return compo
}

func (compo *PaninatorAttribute) SetPages(value []map[string]string) types.PaninatorAttribute {
	(*compo).Pages = value
	return compo
}

func (compo *PaninatorAttribute) SetNextClass(value string) types.PaninatorAttribute {
	(*compo).NextClass = value
	return compo
}

func (compo *PaninatorAttribute) SetNextUrl(value string) types.PaninatorAttribute {
	(*compo).NextUrl = value
	return compo
}

func (compo *PaninatorAttribute) SetOption(value map[string]template.HTML) types.PaninatorAttribute {
	(*compo).Option = value
	return compo
}

func (compo *PaninatorAttribute) SetUrl(value string) types.PaninatorAttribute {
	(*compo).Url = value
	return compo
}

func (compo *PaninatorAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "paninator")
}
