package guard

import (
	"github.com/chenhg5/go-admin/context"
	"html/template"
	"strconv"
)

type MenuNewParam struct {
	Title    string
	Header   string
	ParentId int64
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

	parentIdInt, _ := strconv.Atoi(parentId)

	ctx.SetUserValue("new_menu_param", &MenuNewParam{
		Title:    ctx.FormValue("title"),
		Header:   ctx.FormValue("header"),
		ParentId: int64(parentIdInt),
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
