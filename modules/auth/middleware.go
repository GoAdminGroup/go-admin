// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"net/url"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/page"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// Invoker contains the callback functions which are used
// in the route middleware.
type Invoker struct {
	prefix                 string
	authFailCallback       MiddlewareCallback
	permissionDenyCallback MiddlewareCallback
	conn                   db.Connection
}

// Middleware is the default auth middleware of plugins.
func Middleware(conn db.Connection) context.Handler {
	return DefaultInvoker(conn).Middleware()
}

// DefaultInvoker return a default Invoker.
func DefaultInvoker(conn db.Connection) *Invoker {
	return &Invoker{
		prefix: config.Prefix(),
		authFailCallback: func(ctx *context.Context) {
			if ctx.Request.URL.Path == config.Url(config.GetLoginUrl()) {
				return
			}
			if ctx.Request.URL.Path == config.Url("/logout") {
				ctx.Write(302, map[string]string{
					"Location": config.Url(config.GetLoginUrl()),
				}, ``)
				return
			}
			param := ""
			if ref := ctx.Referer(); ref != "" {
				param = "?ref=" + url.QueryEscape(ref)
			}

			u := config.Url(config.GetLoginUrl() + param)
			_, err := ctx.Request.Cookie(DefaultCookieKey)
			referer := ctx.Referer()

			if (ctx.Headers(constant.PjaxHeader) == "" && ctx.Method() != "GET") ||
				err != nil ||
				referer == "" {
				ctx.Write(302, map[string]string{
					"Location": u,
				}, ``)
			} else {
				msg := language.Get("login overdue, please login again")
				ctx.HTML(http.StatusOK, `<script>
	if (typeof(swal) === "function") {
		swal({
			type: "info",
			title: "`+language.Get("login info")+`",
			text: "`+msg+`",
			showCancelButton: false,
			confirmButtonColor: "#3c8dbc",
			confirmButtonText: '`+language.Get("got it")+`',
        })
		setTimeout(function(){ location.href = "`+u+`"; }, 3000);
	} else {
		alert("`+msg+`")
		location.href = "`+u+`"
    }
</script>`)
			}
		},
		permissionDenyCallback: func(ctx *context.Context) {
			if ctx.Headers(constant.PjaxHeader) == "" && ctx.Method() != "GET" {
				ctx.JSON(http.StatusForbidden, map[string]interface{}{
					"code": http.StatusForbidden,
					"msg":  language.Get(errors.PermissionDenied),
				})
			} else {
				page.SetPageContent(ctx, Auth(ctx), func(ctx interface{}) (types.Panel, error) {
					return template2.WarningPanel(errors.PermissionDenied, template2.NoPermission403Page), nil
				}, conn)
			}
		},
		conn: conn,
	}
}

// SetPrefix return the default Invoker with the given prefix.
func SetPrefix(prefix string, conn db.Connection) *Invoker {
	i := DefaultInvoker(conn)
	i.prefix = prefix
	return i
}

// SetAuthFailCallback set the authFailCallback of Invoker.
func (invoker *Invoker) SetAuthFailCallback(callback MiddlewareCallback) *Invoker {
	invoker.authFailCallback = callback
	return invoker
}

// SetPermissionDenyCallback set the permissionDenyCallback of Invoker.
func (invoker *Invoker) SetPermissionDenyCallback(callback MiddlewareCallback) *Invoker {
	invoker.permissionDenyCallback = callback
	return invoker
}

// MiddlewareCallback is type of callback function.
type MiddlewareCallback func(ctx *context.Context)

// Middleware get the auth middleware from Invoker.
func (invoker *Invoker) Middleware() context.Handler {
	return func(ctx *context.Context) {
		user, authOk, permissionOk := Filter(ctx, invoker.conn)

		if authOk && permissionOk {
			ctx.SetUserValue("user", user)
			ctx.Next()
			return
		}

		if !authOk {
			invoker.authFailCallback(ctx)
			ctx.Abort()
			return
		}

		if !permissionOk {
			ctx.SetUserValue("user", user)
			invoker.permissionDenyCallback(ctx)
			ctx.Abort()
			return
		}
	}
}

// Filter retrieve the user model from Context and check the permission
// at the same time.
func Filter(ctx *context.Context, conn db.Connection) (models.UserModel, bool, bool) {
	var (
		id float64
		ok bool

		user     = models.User()
		ses, err = InitSession(ctx, conn)
	)

	if err != nil {
		logger.Error("retrieve auth user failed", err)
		return user, false, false
	}

	if id, ok = ses.Get("user_id").(float64); !ok {
		return user, false, false
	}

	user, ok = GetCurUserByID(int64(id), conn)

	if !ok {
		return user, false, false
	}

	return user, true, CheckPermissions(user, ctx.Request.URL.String(), ctx.Method(), ctx.PostForm())
}

const defaultUserIDSesKey = "user_id"

// GetUserID return the user id from the session.
func GetUserID(sesKey string, conn db.Connection) int64 {
	id, err := GetSessionByKey(sesKey, defaultUserIDSesKey, conn)
	if err != nil {
		logger.Error("retrieve auth user failed", err)
		return -1
	}
	if idFloat64, ok := id.(float64); ok {
		return int64(idFloat64)
	}
	return -1
}

// GetCurUser return the user model.
func GetCurUser(sesKey string, conn db.Connection) (user models.UserModel, ok bool) {

	if sesKey == "" {
		ok = false
		return
	}

	id := GetUserID(sesKey, conn)
	if id == -1 {
		ok = false
		return
	}
	return GetCurUserByID(id, conn)
}

// GetCurUserByID return the user model of given user id.
func GetCurUserByID(id int64, conn db.Connection) (user models.UserModel, ok bool) {

	user = models.User().SetConn(conn).Find(id)

	if user.IsEmpty() {
		ok = false
		return
	}

	if user.Avatar == "" || config.GetStore().Prefix == "" {
		user.Avatar = ""
	} else {
		user.Avatar = config.GetStore().URL(user.Avatar)
	}

	user = user.WithRoles().WithPermissions().WithMenus()

	ok = user.HasMenu()

	return
}

// CheckPermissions check the permission of the user.
func CheckPermissions(user models.UserModel, path, method string, param url.Values) bool {
	return user.CheckPermissionByUrlMethod(path, method, param)
}
