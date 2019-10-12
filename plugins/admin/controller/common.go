package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

var config c.Config

func SetConfig(cfg c.Config) {
	config = cfg
}

func aAlert() types.AlertAttribute {
	return aTemplate().Alert()
}

func aForm() types.FormAttribute {
	return aTemplate().Form()
}

func aRow() types.RowAttribute {
	return aTemplate().Row()
}

func aCol() types.ColAttribute {
	return aTemplate().Col()
}

func aTree() types.TreeAttribute {
	return aTemplate().Tree()
}

func aDataTable() types.DataTableAttribute {
	return aTemplate().DataTable()
}

func aBox() types.BoxAttribute {
	return aTemplate().Box()
}

func aTab() types.TabsAttribute {
	return aTemplate().Tabs()
}

func aTemplate() template.Template {
	return template.Get(config.Theme)
}

func loginComponent() template.Component {
	return template.GetComp("login")
}

func isPjax(ctx *context.Context) bool {
	return ctx.Headers(constant.PjaxHeader) == "true"
}
