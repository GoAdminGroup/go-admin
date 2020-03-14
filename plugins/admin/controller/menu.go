package controller

import (
	"encoding/json"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
	"net/http"
)

// ShowMenu show menu info page.
func (h *Handler) ShowMenu(ctx *context.Context) {
	h.getMenuInfoPanel(ctx, "")
}

// ShowNewMenu show new menu page.
func (h *Handler) ShowNewMenu(ctx *context.Context) {

	panel := h.table("menu", ctx)

	formInfo := panel.GetNewForm()

	user := auth.Auth(ctx)

	js := `<script>
$('.icon').iconpicker({placement: 'bottomLeft'});
</script>`

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: formContent(aForm().
			SetContent(formInfo.FieldList).
			SetTabContents(formInfo.GroupFieldList).
			SetTabHeaders(formInfo.GroupFieldHeaders).
			SetPrefix(h.config.PrefixFixSlash()).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetUrl(h.config.Url("/menu/edit")).
			SetHiddenFields(map[string]string{
				form2.TokenKey:    h.authSrv().AddToken(),
				form2.PreviousKey: h.config.Url("/menu"),
			}).
			SetOperationFooter(formFooter("new"))) +
			template2.HTML(js),
		Description: panel.GetForm().Description,
		Title:       panel.GetForm().Title,
	}, h.config, menu.GetGlobalMenu(user, h.conn).SetActiveClass(h.config.URLRemovePrefix(ctx.Path())))

	ctx.HTML(http.StatusOK, buf.String())
}

// ShowEditMenu show edit menu page.
func (h *Handler) ShowEditMenu(ctx *context.Context) {

	if ctx.Query("id") == "" {
		h.getMenuInfoPanel(ctx, template.Get(h.config.Theme).Alert().
			SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML("wrong id")).
			GetContent())
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, h.config.Url("/menu"))
		return
	}

	formInfo, err := h.table("menu", ctx).GetDataWithId(parameter.BaseParam().WithPKs(ctx.Query("id")))

	var alert template2.HTML

	if err != nil {
		alert = aAlert().SetTitle(constant.DefaultErrorMsg).
			SetTheme("warning").
			SetContent(template2.HTML(err.Error())).
			GetContent()
	}

	user := auth.Auth(ctx)

	js := `<script>
$('.icon').iconpicker({placement: 'bottomLeft'});
</script>`

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content: alert + formContent(aForm().
			SetContent(formInfo.FieldList).
			SetTabContents(formInfo.GroupFieldList).
			SetTabHeaders(formInfo.GroupFieldHeaders).
			SetPrefix(h.config.PrefixFixSlash()).
			SetPrimaryKey(h.table("menu", ctx).GetPrimaryKey().Name).
			SetUrl(h.config.Url("/menu/edit")).
			SetOperationFooter(formFooter("edit")).
			SetHiddenFields(map[string]string{
				form2.TokenKey:    h.authSrv().AddToken(),
				form2.PreviousKey: h.config.Url("/menu"),
			})) + template2.HTML(js),
		Description: formInfo.Description,
		Title:       formInfo.Title,
	}, h.config, menu.GetGlobalMenu(user, h.conn).SetActiveClass(h.config.URLRemovePrefix(ctx.Path())))

	ctx.HTML(http.StatusOK, buf.String())
}

// DeleteMenu delete the menu of given id.
func (h *Handler) DeleteMenu(ctx *context.Context) {
	models.MenuWithId(guard.GetMenuDeleteParam(ctx).Id).SetConn(h.conn).Delete()
	response.Ok(ctx)
}

// EditMenu edit the menu of given id.
func (h *Handler) EditMenu(ctx *context.Context) {

	param := guard.GetMenuEditParam(ctx)

	if param.HasAlert() {
		h.getMenuInfoPanel(ctx, param.Alert)
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, h.config.Url("/menu"))
		return
	}

	menuModel := models.MenuWithId(param.Id).SetConn(h.conn)

	menuModel.DeleteRoles()
	for _, roleId := range param.Roles {
		menuModel.AddRole(roleId)
	}

	menuModel.Update(param.Title, param.Icon, param.Uri, param.Header, param.ParentId)

	h.getMenuInfoPanel(ctx, "")
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, h.config.Url("/menu"))
}

// NewMenu create a new menu item.
func (h *Handler) NewMenu(ctx *context.Context) {

	param := guard.GetMenuNewParam(ctx)

	if param.HasAlert() {
		h.getMenuInfoPanel(ctx, param.Alert)
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, h.config.Url("/menu"))
		return
	}

	user := auth.Auth(ctx)

	menuModel := models.Menu().SetConn(h.conn).
		New(param.Title, param.Icon, param.Uri, param.Header, param.ParentId, (menu.GetGlobalMenu(user, h.conn)).MaxOrder+1)

	for _, roleId := range param.Roles {
		menuModel.AddRole(roleId)
	}

	menu.GetGlobalMenu(user, h.conn).AddMaxOrder()

	h.getMenuInfoPanel(ctx, "")
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, h.config.Url("/menu"))
}

// MenuOrder change the order of menu items.
func (h *Handler) MenuOrder(ctx *context.Context) {

	var data []map[string]interface{}
	_ = json.Unmarshal([]byte(ctx.FormValue("_order")), &data)

	models.Menu().SetConn(h.conn).ResetOrder(data)

	response.Ok(ctx)
}

func (h *Handler) getMenuInfoPanel(ctx *context.Context, alert template2.HTML) {
	user := auth.Auth(ctx)

	editUrl := h.config.Url("/menu/edit/show")
	deleteUrl := h.config.Url("/menu/delete")
	orderUrl := h.config.Url("/menu/order")

	tree := aTree().
		SetTree((menu.GetGlobalMenu(user, h.conn)).List).
		SetEditUrl(editUrl).
		SetUrlPrefix(h.config.Prefix()).
		SetDeleteUrl(deleteUrl).
		SetOrderUrl(orderUrl).
		GetContent()

	header := aTree().GetTreeHeader()
	box := aBox().SetHeader(header).SetBody(tree).GetContent()
	col1 := aCol().SetSize(types.SizeMD(6)).SetContent(box).GetContent()

	list := h.table("menu", ctx)

	formInfo := list.GetNewForm()

	newForm := menuFormContent(aForm().
		SetPrefix(h.config.PrefixFixSlash()).
		SetUrl(h.config.Url("/menu/new")).
		SetPrimaryKey(h.table("menu", ctx).GetPrimaryKey().Name).
		SetHiddenFields(map[string]string{
			form2.TokenKey:    h.authSrv().AddToken(),
			form2.PreviousKey: h.config.Url("/menu"),
		}).
		SetOperationFooter(formFooter("menu")).
		SetTitle("New").
		SetContent(formInfo.FieldList).
		SetTabContents(formInfo.GroupFieldList).
		SetTabHeaders(formInfo.GroupFieldHeaders))

	col2 := aCol().SetSize(types.SizeMD(6)).SetContent(newForm).GetContent()

	row := aRow().SetContent(col1 + col2).GetContent()

	tmpl, tmplName := aTemplate().GetTemplate(isPjax(ctx))
	buf := template.Execute(tmpl, tmplName, user, types.Panel{
		Content:     alert + row,
		Description: "Menus Manage",
		Title:       "Menus Manage",
	}, h.config, menu.GetGlobalMenu(user, h.conn).SetActiveClass(h.config.URLRemovePrefix(ctx.Path())))

	ctx.HTML(http.StatusOK, buf.String())
}
