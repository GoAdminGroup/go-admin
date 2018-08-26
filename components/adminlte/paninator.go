package adminlte

import "html/template"

type PaninatorAttribute struct {
	Name    string
	Color   string
	Content string
}

func (AdminlteComponents) Paninator() *PaninatorAttribute {
	return &PaninatorAttribute{
		"label",
		"success",
		"",
	}
}

func (compo *PaninatorAttribute) SetContent(value string) *PaninatorAttribute {
	(*compo).Content = value
	return compo
}

func (compo *PaninatorAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "paninator")
}
