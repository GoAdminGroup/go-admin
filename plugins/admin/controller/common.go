package controller

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	c "github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	template2 "html/template"
)

var (
	config        c.Config
	captchaConfig map[string]string
	services      service.List
	conn          db.Connection
)

// SetConfig set the config.
func SetCaptcha(cap map[string]string) {
	captchaConfig = cap
}

// SetConfig set the config.
func SetConfig(cfg c.Config) {
	config = cfg
}

// SetServices set the services.
func SetServices(l service.List) {
	services = l
	conn = db.GetConnection(services)
}

func authSrv() *auth.Service {
	return auth.GetService(services.Get("auth"))
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

func aButton() types.ButtonAttribute {
	return aTemplate().Button()
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

func isPjax(ctx *context.Context) bool {
	return ctx.Headers(constant.PjaxHeader) == "true"
}

func formFooter() template2.HTML {
	col1 := aCol().SetSize(map[string]string{"md": "2"}).GetContent()
	btn1 := aButton().SetType("submit").
		SetContent(language.GetFromHtml("Save")).
		SetThemePrimary().
		SetOrientationRight().
		SetLoadingText(`<i class='fa fa-spinner fa-spin '></i> Save`).
		GetContent()
	btn2 := aButton().SetType("reset").
		SetContent(language.GetFromHtml("Reset")).
		SetThemeWarning().
		SetOrientationLeft().
		GetContent()
	col2 := aCol().SetSize(map[string]string{"md": "8"}).
		SetContent(btn1 + btn2).GetContent()
	return col1 + col2
}

func filterFormFooter(infoUrl string) template2.HTML {
	col1 := aCol().SetSize(map[string]string{"md": "2"}).GetContent()
	btn1 := aButton().SetType("submit").
		SetContent(`<i class="fa fa-search"></i>&nbsp;&nbsp;` + language.GetFromHtml("search")).
		SetThemePrimary().
		SetSmallSize().
		SetOrientationLeft().
		SetLoadingText(`<i class='fa fa-spinner fa-spin '></i> ` + language.GetFromHtml("search")).
		GetContent()
	btn2 := aButton().SetType("reset").
		SetContent(`<i class="fa fa-undo"></i>&nbsp;&nbsp;` + language.GetFromHtml("reset")).
		SetThemeDefault().
		SetOrientationLeft().
		SetSmallSize().
		SetHref(infoUrl).
		SetMarginLeft(12).
		GetContent()
	col2 := aCol().SetSize(map[string]string{"md": "8"}).
		SetContent(btn1 + btn2).GetContent()
	return col1 + col2
}

func formContent(form types.FormAttribute) template2.HTML {
	return aBox().
		SetHeader(form.GetBoxHeader()).
		WithHeadBorder().
		SetBody(form.GetContent()).
		GetContent()
}

func detailContent(form types.FormAttribute, editUrl, deleteUrl string) template2.HTML {
	return aBox().
		SetHeader(form.GetDetailBoxHeader(editUrl, deleteUrl)).
		WithHeadBorder().
		SetBody(form.GetContent()).
		GetContent()
}

func menuFormContent(form types.FormAttribute) template2.HTML {
	return aBox().
		SetHeader(form.GetBoxHeaderNoButton()).
		SetStyle(" ").
		WithHeadBorder().
		SetBody(form.GetContent()).
		GetContent()
}
