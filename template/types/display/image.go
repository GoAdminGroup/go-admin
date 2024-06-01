package display

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Image struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("image", new(Image))
}

func (image *Image) Get(ctx *context.Context, args ...interface{}) types.FieldFilterFn {
	param := args[2].([]string)
	return func(value types.FieldModel) interface{} {
		if value.Value == "" {
			return ""
		}
		if len(param) > 0 {
			return template.Default(ctx).Image().SetWidth(args[0].(string)).SetHeight(args[1].(string)).
				SetSrc(template.HTML(param[0] + value.Value)).GetContent()

		} else {
			return template.Default(ctx).Image().SetWidth(args[0].(string)).SetHeight(args[1].(string)).
				SetSrc(template.HTML(value.Value)).GetContent()
		}
	}
}
