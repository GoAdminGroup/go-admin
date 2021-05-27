package action

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type FieldFilterAction struct {
	BaseAction
	Field string
}

func FieldFilter(field string) *FieldFilterAction {
	return &FieldFilterAction{Field: field}
}

func (jump *FieldFilterAction) ExtContent() template.HTML {

	options := jump.BtnData.(types.FieldOptions)

	cm := ``
	for _, obejct := range options {
		cm += `if (e.params.data.text === "` + obejct.Text + `") {
		$.pjax({url: setURL("` + jump.Field + `", "` + obejct.Value + `"), container: '#pjax-container'});
	}`
	}

	return template.HTML(`<script>
$("select` + jump.BtnId + `").on("select2:select",function(e){

	let setURL = function(field, value) {
		let vars = window.location.search.substring(1).split("&");
		let params = "";
		let has = false;
		for (let i = 0; i < vars.length; i++) {
			pair = vars[i].split("=");
			if (pair[0] == "` + parameter.Page + `") {
				continue
			} else if (pair[0] === field) {
				has = true
				params += field + "=" + value + "&"
			} else if (pair[0] !== "` + form.NoAnimationKey + `") {
				params += vars[i] + "&"
			}
		}

		if (!has) {
			params += field + "=" + value + "&` + form.NoAnimationKey + `=true"
		} else {
			params +=  "` + form.NoAnimationKey + `=true"
		}

		return window.location.pathname + "?" + params
	}
	
	` + cm + `
})
vv = ""
query = window.location.search.substring(1);
vars = query.split("&");
for (let i = 0; i < vars.length; i++) {
	pair = vars[i].split("=");
	if (pair[0] === "` + jump.Field + `") {
		vv = pair[1];
	}
}
if (vv !== "") {
	$("` + jump.BtnId + `").val(vv).select2()
}
</script>`)
}
