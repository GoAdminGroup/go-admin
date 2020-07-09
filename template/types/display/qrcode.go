package display

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/template/types"
)

type Qrcode struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("qrcode", new(Qrcode))
}

func (q *Qrcode) Get(args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {

		src := `https://api.qrserver.com/v1/create-qr-code/?size=150x150&amp;data=` + value.Value

		return template.HTML(`
<a href="javascript:void(0);" class="grid-column-qrcode text-muted" 
	data-content="<img src='` + src + `' 
style='height:150px;width:150px;'/>" data-toggle="popover" tabindex="0" data-original-title="" title="">
<i class="fa fa-qrcode"></i>
</a>&nbsp;` + value.Value + `
`)
	}
}

func (q *Qrcode) JS() template.HTML {
	return template.HTML(`
$('.grid-column-qrcode').popover({
	html: true,
	container: 'body',
	trigger: 'focus'
});
`)
}
