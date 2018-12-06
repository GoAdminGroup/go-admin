package components

import (
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
)

type ProductListAttribute struct {
	Name string
	Data []map[string]string
	types.Attribute
}

func (compo *ProductListAttribute) SetData(value []map[string]string) types.ProductListAttribute {
	compo.Data = value
	return compo
}

func (compo *ProductListAttribute) GetContent() template.HTML {
	return ComposeHtml(compo.TemplateList, *compo, "productlist")
}
