package display

import (
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Link struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("link", new(Link))
}

func (l *Link) Get(args ...interface{}) types.FieldFilterFn {
	prefix := ""
	openInNewTabs := false
	if len(args) > 0 {
		prefix = args[0].(string)
	}
	if len(args) > 1 {
		if openInNewTabsArr, ok := args[1].([]bool); ok {
			openInNewTabs = openInNewTabsArr[0]
		}
	}
	return func(value types.FieldModel) interface{} {
		if openInNewTabs {
			return template.Default().Link().SetURL(prefix + value.Value).OpenInNewTab().GetContent()
		} else {
			return template.Default().Link().SetURL(prefix + value.Value).GetContent()
		}
	}
}
