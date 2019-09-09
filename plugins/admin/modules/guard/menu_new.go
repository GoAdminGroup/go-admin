package guard

import (
	"github.com/chenhg5/go-admin/context"
	"html/template"
)

type MenuNewParam struct {
	Title    string
	ParentId string
	Icon     string
	Uri      string
	Roles    []string
	Alert    template.HTML
}

func (e MenuNewParam) HasAlert() bool {
	return e.Alert != template.HTML("")
}

func MenuNew(ctx *context.Context) {

	parentId := ctx.FormValue("parent_id")
	if parentId == "" {
		parentId = "0"
	}

	ctx.SetUserValue("new_menu_param", &MenuNewParam{
		Title:    ctx.FormValue("title"),
		ParentId: parentId,
		Icon:     ctx.FormValue("icon"),
		Uri:      ctx.FormValue("uri"),
		Roles:    ctx.Request.Form["roles[]"],
		Alert:    checkEmpty(ctx, "title", "icon", "uri"),
	})
	ctx.Next()
	return
}

func GetMenuNewParam(ctx *context.Context) *MenuNewParam {
	return ctx.UserValue["new_menu_param"].(*MenuNewParam)
}
