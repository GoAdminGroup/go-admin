package display

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type Dot struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("dot", new(Dot))
}

func (d *Dot) Get(args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		icons := args[0].(map[string]types.FieldDotColor)
		defaultDot := types.FieldDotColor("")
		if len(args) > 1 {
			defaultDot = args[1].(types.FieldDotColor)
		}
		for k, style := range icons {
			if k == value.Value {
				return template.HTML(`<span class="label-`+style+`"
					style="width: 8px;height: 8px;padding: 0;border-radius: 50%;display: inline-block;">
					</span>&nbsp;&nbsp;`) +
					template.HTML(value.Value)
			}
		}
		if defaultDot != "" {
			return template.HTML(`<span class="label-`+defaultDot+`"
					style="width: 8px;height: 8px;padding: 0;border-radius: 50%;display: inline-block;">
					</span>&nbsp;&nbsp;`) +
				template.HTML(value.Value)
		}
		return value.Value
	}
}
