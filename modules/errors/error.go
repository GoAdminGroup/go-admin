package errors

import (
	"errors"
	"html/template"

	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template/icon"
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

	PageError404 = errors.New(language.Get("not found"))
	PageError500 = errors.New(language.Get("internal error"))
	PageError403 = errors.New(language.Get("permission denied"))
	PageError401 = errors.New(language.Get("unauthorized"))
}

type PageError error

var (
	PageError404 PageError
	PageError500 PageError
	PageError403 PageError
	PageError401 PageError
)
