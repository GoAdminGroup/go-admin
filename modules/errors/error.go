package errors

import (
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template/icon"
	"html/template"
)

var (
	Msg         string
	MsgHTML     template.HTML
	MsgWithIcon template.HTML
)

const (
	PermissionDenied     = "permission denied"
	WrongID              = "wrong id"
	OperationNotAllow    = "operation not allow"
	EditFailWrongToken   = "edit fail, wrong token"
	CreateFailWrongToken = "create fail, wrong token"
	NoPermission         = "no permission"
	SiteOff              = "site is off"
)

func WrongPK(pk string) string {
	return "wrong " + pk
}

func Init() {
	Msg = language.Get("error")
	MsgHTML = language.GetFromHtml("error")
	MsgWithIcon = icon.Icon(icon.Warning, 2) + MsgHTML + `!`
}
