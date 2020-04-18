package display

import (
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Image struct{}

func init() {
	types.RegisterDisplayFnGenerator("image", new(Image))
}

func (image *Image) Get(args ...interface{}) types.FieldFilterFn {

	param := args[2].([]interface{})
	return func(value types.FieldModel) interface{} {
		if len(param) > 0 {
			return template.Default().Image().SetWidth(args[0].(string)).SetHeight(args[1].(string)).SetSrc(template.HTML(param[0].(string) + value.Value)).GetContent()

		} else {
			return template.Default().Image().SetWidth(args[0].(string)).SetHeight(args[1].(string)).SetSrc(template.HTML(value.Value)).GetContent()
		}
	}

}
