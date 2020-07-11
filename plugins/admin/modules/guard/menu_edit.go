package guard

import (
	"html/template"
	"strconv"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
)

type MenuEditParam struct {
	Id         string
	Title      string
	Header     string
	PluginName string
	ParentId   int64
	Icon       string
	Uri        string
	Roles      []string
	Alert      template.HTML
}

func (e MenuEditParam) HasAlert() bool {
	return e.Alert != template.HTML("")
}

func (g *Guard) MenuEdit(ctx *context.Context) {

	parentId := ctx.FormValue("parent_id")
	if parentId == "" {
		parentId = "0"
	}

	var (
		parentIdInt, _ = strconv.Atoi(parentId)
		token          = ctx.FormValue(form.TokenKey)
		alert          template.HTML
	)

	if !auth.GetTokenService(g.services.Get(auth.TokenServiceKey)).CheckToken(token) {
		alert = getAlert(errors.EditFailWrongToken)
	}

	if alert == "" {
		alert = checkEmpty(ctx, "id", "title", "icon")
	}

	ctx.SetUserValue(editMenuParamKey, &MenuEditParam{
		Id:         ctx.FormValue("id"),
		Title:      ctx.FormValue("title"),
		Header:     ctx.FormValue("header"),
		PluginName: ctx.FormValue("plugin_name"),
		ParentId:   int64(parentIdInt),
		Icon:       ctx.FormValue("icon"),
		Uri:        ctx.FormValue("uri"),
		Roles:      ctx.Request.Form["roles[]"],
		Alert:      alert,
	})
	ctx.Next()
}

func GetMenuEditParam(ctx *context.Context) *MenuEditParam {
	return ctx.UserValue[editMenuParamKey].(*MenuEditParam)
}

func checkEmpty(ctx *context.Context, key ...string) template.HTML {
	for _, k := range key {
		if ctx.FormValue(k) == "" {
			return getAlert("wrong " + k)
		}
	}
	return template.HTML("")
}
