package components

import "html/template"

type RowAttribute struct {
	Name    string
	Content template.HTML
}

func Row() *RowAttribute {
	return &RowAttribute{
		"row",
		"",
	}
}

func (compo *RowAttribute) SetContent(value template.HTML) *RowAttribute {
	(*compo).Content = value
	return compo
}

func (compo *RowAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "row")
}
