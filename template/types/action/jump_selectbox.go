package action

import (
	"html/template"
)

type JumpSelectBoxAction struct {
	BaseAction
	Options     JumpOptions
	NewTabTitle string
}

type JumpOptions []JumpOption

type JumpOption struct {
	Value string
	Url   string
}

func SelectBoxJump(options JumpOptions) *JumpSelectBoxAction {
	return &JumpSelectBoxAction{Options: options}
}

func (jump *JumpSelectBoxAction) ExtContent() template.HTML {

	cm := ``
	for _, obejct := range jump.Options {
		cm += `if (e.params.data.text === "` + obejct.Value + `") {
		$.pjax({url: "` + obejct.Url + `", container: '#pjax-container'});
	}`
	}

	return template.HTML(`<script>
$("select` + jump.BtnId + `").on("select2:select",function(e){
	` + cm + `
})
</script>`)
}
