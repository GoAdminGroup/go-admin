package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type TabsAttribute struct {
	Name string
	Data []map[string]template.HTML
	types.Attribute
}

func (compo *TabsAttribute) SetData(value []map[string]template.HTML) types.TabsAttribute {
	compo.Data = value
	return compo
}

func (compo *TabsAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "tabs")
}
