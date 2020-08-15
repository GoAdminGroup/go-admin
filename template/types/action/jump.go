package action

import (
	"html/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/utils"
)

type JumpAction struct {
	BaseAction
	Url         string
	Target      string
	Ext         template.HTML
	NewTabTitle string
}

func Jump(url string, ext ...template.HTML) *JumpAction {
	url = utils.ReplaceAll(url, "{%id}", "{{.Id}}", "{%ids}", "{{.Ids}}")
	if len(ext) > 0 {
		return &JumpAction{Url: url, Ext: ext[0]}
	}
	return &JumpAction{Url: url, NewTabTitle: ""}
}

func JumpInNewTab(url, title string, ext ...template.HTML) *JumpAction {
	url = utils.ReplaceAll(url, "{%id}", "{{.Id}}", "{%ids}", "{{.Ids}}")
	if len(ext) > 0 {
		return &JumpAction{Url: url, NewTabTitle: title, Ext: ext[0]}
	}
	return &JumpAction{Url: url, NewTabTitle: title}
}

func JumpWithTarget(url, target string, ext ...template.HTML) *JumpAction {
	url = utils.ReplaceAll(url, "{%id}", "{{.Id}}", "{%ids}", "{{.Ids}}")
	if len(ext) > 0 {
		return &JumpAction{Url: url, Target: target, Ext: ext[0]}
	}
	return &JumpAction{Url: url, Target: target}
}

func (jump *JumpAction) GetCallbacks() context.Node {
	return context.Node{Path: jump.Url, Method: "GET"}
}

func (jump *JumpAction) BtnAttribute() template.HTML {
	html := template.HTML(`href="` + jump.Url + `"`)
	if jump.NewTabTitle != "" {
		html += template.HTML(` data-title="` + jump.NewTabTitle + `"`)
	}
	if jump.Target != "" {
		html += template.HTML(` target="` + jump.Target + `"`)
	}
	return html
}

func (jump *JumpAction) BtnClass() template.HTML {
	if jump.NewTabTitle != "" {
		return "new-tab-link"
	}
	return ""
}

func (jump *JumpAction) ExtContent() template.HTML {
	return jump.Ext
}
