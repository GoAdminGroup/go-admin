// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/page"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"regexp"
	"strings"
)

type Invoker struct {
	prefix                 string
	authFailCallback       MiddlewareCallback
	permissionDenyCallback MiddlewareCallback
}

var Middleware = DefaultInvoker().Middleware()

func DefaultInvoker() *Invoker {
	return &Invoker{
		prefix: config.Get().Prefix(),
		authFailCallback: func(ctx *context.Context) {
			ctx.Write(302, map[string]string{
				"Location": config.Get().Url("/login"),
			}, ``)
		},
		permissionDenyCallback: func(ctx *context.Context) {
			page.SetPageContent(ctx, Auth(ctx), func(ctx interface{}) (types.Panel, error) {
				alert := template2.Get(config.Get().Theme).Alert().
					SetTitle(template.HTML(`<i class="icon fa fa-warning"></i> ` + language.Get("error") + `!`)).
					SetTheme("warning").SetContent(template.HTML("permission denied")).GetContent()

				return types.Panel{
					Content:     alert,
					Description: "Error",
					Title:       "Error",
				}, nil
			})
		},
	}
}

func SetPrefix(prefix string) *Invoker {
	i := DefaultInvoker()
	i.prefix = prefix
	return i
}

func (invoker *Invoker) SetAuthFailCallback(callback MiddlewareCallback) *Invoker {
	invoker.authFailCallback = callback
	return invoker
}

func (invoker *Invoker) SetPermissionDenyCallback(callback MiddlewareCallback) *Invoker {
	invoker.permissionDenyCallback = callback
	return invoker
}

type MiddlewareCallback func(ctx *context.Context)

func (invoker *Invoker) Middleware() context.Handler {
	return func(ctx *context.Context) {
		user, authOk, permissionOk := Filter(ctx)

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

func Filter(ctx *context.Context) (models.UserModel, bool, bool) {
	var (
		id   float64
		ok   bool
		user = models.User()
	)

	if id, ok = InitSession(ctx).Get("user_id").(float64); !ok {
		return user, false, false
	}

	user, ok = GetCurUserById(int64(id))

	if !ok {
		return user, false, false
	}

	return user, true, CheckPermissions(user, ctx.Path(), ctx.Method())
}

func GetCurUserById(id int64) (user models.UserModel, ok bool) {

	user = models.User().Find(id)

	if user.IsEmpty() {
		ok = false
		return
	}

	if user.Avatar == "" || config.Get().Store.Prefix == "" {
		user.Avatar = ""
	} else {
		user.Avatar = "/" + config.Get().Store.Prefix + "/" + user.Avatar
	}

	user = user.WithRoles().WithPermissions().WithMenus()

	ok = true

	return
}

func CheckPermissions(user models.UserModel, path string, method string) bool {

	if path == config.Get().Url("/logout") {
		return true
	}

	for _, v := range user.Permissions {

		if v.HttpMethod[0] == "" || InMethodArr(v.HttpMethod, method) {

			if v.HttpPath[0] == "*" {
				return true
			}

			for i := 0; i < len(v.HttpPath); i++ {

				matchPath := config.Get().Url(strings.TrimSpace(v.HttpPath[i]))

				if matchPath == path {
					return true
				}

				reg, err := regexp.Compile(matchPath)

				if err != nil {
					fmt.Println("err", err)
					continue
				}

				if reg.FindString(path) == path {
					return true
				}
			}
		}
	}

	return false
}

func InMethodArr(arr []string, str string) bool {
	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(arr[i], str) {
			return true
		}
	}
	return false
}
