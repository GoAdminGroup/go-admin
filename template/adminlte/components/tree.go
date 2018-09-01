package components

import (
	"html/template"
	"goAdmin/modules/menu"
)

type TreeAttribute struct {
	Name    string
	Tree   []menu.MenuItem
}

func Tree() *TreeAttribute {
	return &TreeAttribute{
		"tree",
		[]menu.MenuItem{},
	}
}

func (compo *TreeAttribute) SetTree(value []menu.MenuItem) *TreeAttribute {
	(*compo).Tree = value
	return compo
}

func (compo *TreeAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "tree")
}

func (compo *TreeAttribute) GetTreeHeader() template.HTML {
	return ComposeHtml(*compo, "tree-header")
}