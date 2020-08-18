package components

import (
	"encoding/json"
	"html/template"

	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type TreeViewAttribute struct {
	Name      string
	ID        string
	Tree      types.TreeViewData
	TreeJSON  template.JS
	UrlPrefix string
	types.Attribute
}

func (compo *TreeViewAttribute) SetID(id string) types.TreeViewAttribute {
	compo.ID = id
	return compo
}

func (compo *TreeViewAttribute) SetTree(value types.TreeViewData) types.TreeViewAttribute {
	compo.Tree = value
	return compo
}

func (compo *TreeViewAttribute) SetUrlPrefix(value string) types.TreeViewAttribute {
	compo.UrlPrefix = value
	return compo
}

func (compo *TreeViewAttribute) GetContent() template.HTML {
	if compo.ID == "" {
		compo.ID = utils.Uuid(10)
	}
	b, _ := json.Marshal(compo.Tree)
	compo.TreeJSON = template.JS(b)
	return ComposeHtml(compo.TemplateList, compo.Separation, *compo, "treeview")
}
