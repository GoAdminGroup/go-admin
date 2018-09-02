package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type RowAttribute struct {
	Name    string
	Content template.HTML
}

func (*AdminlteStruct) Row() types.RowAttribute {
	return &RowAttribute{
		"row",
		"",
	}
}

func (compo *RowAttribute) SetContent(value template.HTML) types.RowAttribute {
	(*compo).Content = value
	return compo
}

func (compo *RowAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "row")
}
