package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/modules/menu"
	"github.com/chenhg5/go-admin/template/types"
)

type TreeAttribute struct {
	Name    string
	Tree    []menu.MenuItem
	EditUrl string
}

func (compo *TreeAttribute) SetTree(value []menu.MenuItem) types.TreeAttribute {
	(*compo).Tree = value
	return compo
}

func (compo *TreeAttribute) SetEditUrl(value string) types.TreeAttribute {
	(*compo).EditUrl = value
	return compo
}

func (compo *TreeAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "tree")
}

func (compo *TreeAttribute) GetTreeHeader() template.HTML {
	return ComposeHtml(*compo, "tree-header")
}
