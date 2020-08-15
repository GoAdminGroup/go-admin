package types

import (
	"html/template"
	"strconv"

	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/utils"
)

type DefaultSelection struct {
	*BaseButton
	Options     FieldOptions
	Placeholder string
	Width       int
}

func btnUUID() string {
	return "info-btn-" + utils.Uuid(10)
}

func GetDefaultSelection(placeholder string, options FieldOptions, action Action, widths ...int) *DefaultSelection {

	id := btnUUID()
	action.SetBtnId("." + id)

	var width = 100
	if len(widths) > 0 {
		width = widths[0]
	}
	node := action.GetCallbacks()
	return &DefaultSelection{
		BaseButton: &BaseButton{
			Id:     id,
			Action: action,
			Url:    node.Path,
			Method: node.Method,
		},
		Placeholder: placeholder,
		Width:       width,
		Options:     options,
	}

}

func (b *DefaultSelection) Content() (template.HTML, template.JS) {

	optionsHtml := `<option value='__go_admin_all__'>` + language.Get("All") + `</option>`

	for _, op := range b.Options {
		optionsHtml += `<option value='` + op.Value + `'>` + op.Text + `</option>`
	}

	h := template.HTML(`<div class="btn-group pull-right" style="margin-right: 10px">
<div style="width:`+strconv.Itoa(b.Width)+`px;">
<select style="width:100%;height:30px;" class="`+b.Id+` select2-hidden-accessible" name="`+b.Id+`"
            data-multiple="false"  data-placeholder="`+b.Placeholder+`" tabindex="-1" aria-hidden="true">
	<option></option>
    `+optionsHtml+`
</select>
</div>
</div>
<style type="text/css">
	.box-header .select2-container .select2-selection--single {
		height: 30px;
		line-height: 24px;
	}
	.box-header .select2-container--default .select2-selection--single .select2-selection__rendered
	{
		line-height: 24px;
	}
</style>`) + b.Action.ExtContent()

	return h, b.Action.Js() + template.JS(`
	$(".`+b.Id+`").select2();
`)
}
