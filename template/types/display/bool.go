package display

import (
	"strings"

	"github.com/GoAdminGroup/go-admin/template/icon"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/html"
)

type Bool struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("bool", new(Bool))
}

func (b *Bool) Get(args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		params := args[0].([]string)
		pass := icon.IconWithStyle(icon.Check, html.Style{
			"color": "green",
		})
		fail := icon.IconWithStyle(icon.Remove, html.Style{
			"color": "red",
		})
		if len(params) == 0 {
			if value.Value == "0" || strings.ToLower(value.Value) == "false" {
				return fail
			}
			return pass
		} else if len(params) == 1 {
			if value.Value == params[0] {
				return pass
			}
			return fail
		} else {
			if value.Value == params[0] {
				return pass
			}
			if value.Value == params[1] {
				return fail
			}
		}
		return ""
	}
}
