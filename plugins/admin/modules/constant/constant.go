package constant

import (
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template/icon"
)

const (
	// PjaxHeader is default pjax http header key.
	PjaxHeader = constant.PjaxHeader

	// PjaxUrlHeader is default pjax url http header key.
	PjaxUrlHeader = constant.PjaxUrlHeader

	EditPKKey   = "__goadmin_edit_pk"
	DetailPKKey = "__goadmin_detail_pk"
	PrefixKey   = "__prefix"

	PostTypeKey           = "__go_admin_post_type"
	PostIsSingleUpdateKey = "__go_admin_is_single_update"
)

var (
	DefaultErrorMsg = icon.Icon(icon.Warning, 2) + language.GetFromHtml("error") + `!`
)
