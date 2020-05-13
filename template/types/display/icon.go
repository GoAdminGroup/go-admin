package display

import (
	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Icon struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("icon", new(Icon))
}

func (i *Icon) Get(args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		icons := args[0].(map[string]string)
		defaultIcon := ""
		if len(args) > 1 {
			defaultIcon = args[1].(string)
		}
		for k, iconClass := range icons {
			if k == value.Value {
				return icon.Icon(iconClass)
			}
		}
		if defaultIcon != "" {
			return icon.Icon(defaultIcon)
		}
		return value.Value
	}
}
