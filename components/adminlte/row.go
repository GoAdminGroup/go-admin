package adminlte

import "html/template"

type RowAttribute struct {
	Name    string
	Content template.HTML
}

func (AdminlteComponents) Row() *RowAttribute {
	return &RowAttribute{
		"row",
		"",
	}
}

func (compo *RowAttribute) SetContent(value string) *RowAttribute {
	(*compo).Content = template.HTML(value)
	return compo
}

func (compo *RowAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "row")
}
