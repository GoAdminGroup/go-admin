package components

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/google/uuid"
)

type TabsAttribute struct {
	Name string
	Id   string
	Data []map[string]template.HTML
	types.Attribute
}

func (compo *TabsAttribute) SetData(value []map[string]template.HTML) types.TabsAttribute {
	compo.Data = value
	return compo
}

func (compo *TabsAttribute) GetContent() template.HTML {
	if compo.Id == "" {
		compo.Id = uuid.New().String()[:8]
	}
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "tabs")
}
