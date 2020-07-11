package controller

import (
	"encoding/json"
	template2 "html/template"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/menu"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/guard"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/response"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// ShowMenu show menu info page.
func (h *Handler) ShowMenu(ctx *context.Context) {
	h.getMenuInfoPanel(ctx, "", "")
}

// ShowNewMenu show new menu page.
func (h *Handler) ShowNewMenu(ctx *context.Context) {
	h.showNewMenu(ctx, nil)
}

func (h *Handler) showNewMenu(ctx *context.Context, err error) {
	panel := h.table("menu", ctx)

	formInfo := panel.GetNewForm()

	user := auth.Auth(ctx)

	var alert template2.HTML

	if err != nil {
		alert = aAlert().Warning(err.Error())
	}

	h.HTML(ctx, user, types.Panel{
		Content: alert + formContent(aForm().
			SetContent(formInfo.FieldList).
			SetTabContents(formInfo.GroupFieldList).
			SetTabHeaders(formInfo.GroupFieldHeaders).
			SetPrefix(h.config.PrefixFixSlash()).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetUrl(h.routePath("menu_edit")).
			SetHiddenFields(map[string]string{
				form2.TokenKey:    h.authSrv().AddToken(),
				form2.PreviousKey: h.routePath("menu"),
			}).
			SetOperationFooter(formFooter("new", false, false, false,
				panel.GetForm().FormNewBtnWord)),
			false, ctx.Query(constant.IframeKey) == "true", false, ""),
		Description: template2.HTML(panel.GetForm().Description),
		Title:       template2.HTML(panel.GetForm().Title),
	})
}

// ShowEditMenu show edit menu page.
func (h *Handler) ShowEditMenu(ctx *context.Context) {

	if ctx.Query("id") == "" {
		h.getMenuInfoPanel(ctx, "", template.Get(h.config.Theme).Alert().Warning(errors.WrongID))

		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, h.routePath("menu"))
		return
	}

	model := h.table("menu", ctx)
	formInfo, err := model.GetDataWithId(parameter.BaseParam().WithPKs(ctx.Query("id")))

	user := auth.Auth(ctx)

	if err != nil {
		h.HTML(ctx, user, types.Panel{
			Content:     aAlert().Warning(err.Error()),
			Description: template2.HTML(model.GetForm().Description),
			Title:       template2.HTML(model.GetForm().Title),
		})
		return
	}

	h.showEditMenu(ctx, "", formInfo, nil)
}

func (h *Handler) showEditMenu(ctx *context.Context, params string, formInfo table.FormInfo, err error) {

	var alert template2.HTML

	if err != nil {
		alert = aAlert().Warning(err.Error())
	}

	panel := h.table("menu", ctx)

	h.HTML(ctx, auth.Auth(ctx), types.Panel{
		Content: alert + formContent(aForm().
			SetContent(formInfo.FieldList).
			SetTabContents(formInfo.GroupFieldList).
			SetTabHeaders(formInfo.GroupFieldHeaders).
			SetPrefix(h.config.PrefixFixSlash()).
			SetPrimaryKey(panel.GetPrimaryKey().Name).
			SetUrl(h.routePath("menu_edit")).
			SetOperationFooter(formFooter("edit", false, false, false,
				panel.GetForm().FormEditBtnWord)).
			SetHiddenFields(map[string]string{
				form2.TokenKey:    h.authSrv().AddToken(),
				form2.PreviousKey: h.routePath("menu") + params,
			}), false, ctx.Query(constant.IframeKey) == "true", false, ""),
		Description: template2.HTML(formInfo.Description),
		Title:       template2.HTML(formInfo.Title),
	})
	return
}

// DeleteMenu delete the menu of given id.
func (h *Handler) DeleteMenu(ctx *context.Context) {
	models.MenuWithId(guard.GetMenuDeleteParam(ctx).Id).SetConn(h.conn).Delete()
	response.OkWithMsg(ctx, language.Get("delete succeed"))
}

// EditMenu edit the menu of given id.
func (h *Handler) EditMenu(ctx *context.Context) {

	param := guard.GetMenuEditParam(ctx)
	params := ""
	if param.PluginName != "" {
		params = "?__plugin_name=" + param.PluginName
	}

	if param.HasAlert() {
		h.getMenuInfoPanel(ctx, param.PluginName, param.Alert)
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, h.routePath("menu")+params)
		return
	}

	menuModel := models.MenuWithId(param.Id).SetConn(h.conn)

	// TODO: use transaction
	deleteRolesErr := menuModel.DeleteRoles()
	if db.CheckError(deleteRolesErr, db.DELETE) {
		formInfo, _ := h.table("menu", ctx).GetDataWithId(parameter.BaseParam().WithPKs(param.Id))
		h.showEditMenu(ctx, params, formInfo, deleteRolesErr)
		ctx.AddHeader(constant.PjaxUrlHeader, h.routePath("menu")+params)
		return
	}
	for _, roleId := range param.Roles {
		_, addRoleErr := menuModel.AddRole(roleId)
		if db.CheckError(addRoleErr, db.INSERT) {
			formInfo, _ := h.table("menu", ctx).GetDataWithId(parameter.BaseParam().WithPKs(param.Id))
			h.showEditMenu(ctx, params, formInfo, addRoleErr)
			ctx.AddHeader(constant.PjaxUrlHeader, h.routePath("menu")+params)
			return
		}
	}

	_, updateErr := menuModel.Update(param.Title, param.Icon, param.Uri, param.Header, param.PluginName, param.ParentId)

	if db.CheckError(updateErr, db.UPDATE) {
		formInfo, _ := h.table("menu", ctx).GetDataWithId(parameter.BaseParam().WithPKs(param.Id))
		h.showEditMenu(ctx, params, formInfo, updateErr)
		ctx.AddHeader(constant.PjaxUrlHeader, h.routePath("menu")+params)
		return
	}

	h.getMenuInfoPanel(ctx, param.PluginName, "")
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, h.routePath("menu")+params)
}

// NewMenu create a new menu item.
func (h *Handler) NewMenu(ctx *context.Context) {

	param := guard.GetMenuNewParam(ctx)
	params := ""
	if param.PluginName != "" {
		params = "?__plugin_name=" + param.PluginName
	}

	if param.HasAlert() {
		h.getMenuInfoPanel(ctx, param.PluginName, param.Alert)
		ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
		ctx.AddHeader(constant.PjaxUrlHeader, h.routePath("menu")+params)
		return
	}

	user := auth.Auth(ctx)

	// TODO: use transaction
	menuModel, createErr := models.Menu().SetConn(h.conn).
		New(param.Title, param.Icon, param.Uri, param.Header, param.PluginName, param.ParentId,
			(menu.GetGlobalMenu(user, h.conn, param.PluginName)).MaxOrder+1)

	if db.CheckError(createErr, db.INSERT) {
		h.showNewMenu(ctx, createErr)
		return
	}

	for _, roleId := range param.Roles {
		_, addRoleErr := menuModel.AddRole(roleId)
		if db.CheckError(addRoleErr, db.INSERT) {
			h.showNewMenu(ctx, addRoleErr)
			return
		}
	}

	menu.GetGlobalMenu(user, h.conn, param.PluginName).AddMaxOrder()

	h.getMenuInfoPanel(ctx, param.PluginName, "")
	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")
	ctx.AddHeader(constant.PjaxUrlHeader, h.routePath("menu")+params)
}

// MenuOrder change the order of menu items.
func (h *Handler) MenuOrder(ctx *context.Context) {

	var data []map[string]interface{}
	_ = json.Unmarshal([]byte(ctx.FormValue("_order")), &data)

	models.Menu().SetConn(h.conn).ResetOrder([]byte(ctx.FormValue("_order")))

	response.Ok(ctx)
}

func (h *Handler) getMenuInfoPanel(ctx *context.Context, plugName string, alert template2.HTML) {
	user := auth.Auth(ctx)

	if plugName == "" {
		plugName = ctx.Query("__plugin_name")
	}

	tree := aTree().
		SetTree((menu.GetGlobalMenu(user, h.conn, plugName)).List).
		SetEditUrl(h.routePath("menu_edit_show")).
		SetUrlPrefix(h.config.Prefix()).
		SetDeleteUrl(h.routePath("menu_delete")).
		SetOrderUrl(h.routePath("menu_order")).
		GetContent()

	header := aTree().GetTreeHeader()
	box := aBox().SetHeader(header).SetBody(tree).GetContent()
	col1 := aCol().SetSize(types.SizeMD(6)).SetContent(box).GetContent()
	panel := h.table("menu", ctx)

	formInfo := panel.GetNewForm()

	previousURL := h.routePath("menu")

	if plugName != "" {
		previousURL += "?__plugin_name=" + plugName
	}

	newForm := menuFormContent(aForm().
		SetPrefix(h.config.PrefixFixSlash()).
		SetUrl(h.routePath("menu_new")).
		SetPrimaryKey(panel.GetPrimaryKey().Name).
		SetHiddenFields(map[string]string{
			form2.TokenKey:    h.authSrv().AddToken(),
			form2.PreviousKey: previousURL,
		}).
		SetOperationFooter(formFooter("menu", false, false, false,
			panel.GetForm().FormNewBtnWord)).
		SetTitle("New").
		SetContent(formInfo.FieldList).
		SetTabContents(formInfo.GroupFieldList).
		SetTabHeaders(formInfo.GroupFieldHeaders))

	col2 := aCol().SetSize(types.SizeMD(6)).SetContent(newForm).GetContent()

	row := aRow().SetContent(col1 + col2).GetContent()

	h.HTML(ctx, user, types.Panel{
		Content:     alert + row,
		Description: "Menus Manage",
		Title:       "Menus Manage",
	})
}
