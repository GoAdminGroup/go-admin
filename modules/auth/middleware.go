// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/page"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"net/http"
	"net/url"
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
			ctx.Write(302, map[string]string{
				"Location": config.Url("/login"),
			}, ``)
		},
		permissionDenyCallback: func(ctx *context.Context) {
			if ctx.Headers(constant.PjaxHeader) == "" && ctx.Method() != "GET" {
				ctx.JSON(http.StatusForbidden, map[string]interface{}{
					"code": http.StatusForbidden,
					"msg":  language.Get(errors.PermissionDenied),
				})
			} else {
				page.SetPageContent(ctx, Auth(ctx), func(ctx interface{}) (types.Panel, error) {
					return template2.WarningPanel(errors.PermissionDenied), nil
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
		id   float64
		ok   bool
		user = models.User()
	)

	if id, ok = InitSession(ctx, conn).Get("user_id").(float64); !ok {
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
	id := GetSessionByKey(sesKey, defaultUserIDSesKey, conn)
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
