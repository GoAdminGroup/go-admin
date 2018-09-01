package components

import "html/template"

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

func Paninator() *PaninatorAttribute {
	return &PaninatorAttribute{
		Name:    "paninator",
	}
}

func (compo *PaninatorAttribute) SetCurPageStartIndex(value string) *PaninatorAttribute {
	(*compo).CurPageStartIndex = value
	return compo
}

func (compo *PaninatorAttribute) SetCurPageEndIndex(value string) *PaninatorAttribute {
	(*compo).CurPageEndIndex = value
	return compo
}

func (compo *PaninatorAttribute) SetTotal(value string) *PaninatorAttribute {
	(*compo).Total = value
	return compo
}

func (compo *PaninatorAttribute) SetPreviousClass(value string) *PaninatorAttribute {
	(*compo).PreviousClass = value
	return compo
}

func (compo *PaninatorAttribute) SetPreviousUrl(value string) *PaninatorAttribute {
	(*compo).PreviousUrl = value
	return compo
}

func (compo *PaninatorAttribute) SetPages(value []map[string]string) *PaninatorAttribute {
	(*compo).Pages = value
	return compo
}

func (compo *PaninatorAttribute) SetNextClass(value string) *PaninatorAttribute {
	(*compo).NextClass = value
	return compo
}

func (compo *PaninatorAttribute) SetNextUrl(value string) *PaninatorAttribute {
	(*compo).NextUrl = value
	return compo
}

func (compo *PaninatorAttribute) SetOption(value map[string]template.HTML) *PaninatorAttribute {
	(*compo).Option = value
	return compo
}

func (compo *PaninatorAttribute) SetUrl(value string) *PaninatorAttribute {
	(*compo).Url = value
	return compo
}

func (compo *PaninatorAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "paninator")
}
