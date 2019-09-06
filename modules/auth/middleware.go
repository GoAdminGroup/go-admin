// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/page"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	template2 "github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
	"regexp"
)

type Invoker struct {
	prefix                 string
	authFailCallback       MiddlewareCallback
	permissionDenyCallback MiddlewareCallback
}

func Middleware() context.Handler {
	return DefaultInvoker().Middleware()
}

func DefaultInvoker() *Invoker {
	return &Invoker{
		prefix: "/" + config.Get().PREFIX,
		authFailCallback: func(ctx *context.Context) {
			ctx.Write(302, map[string]string{
				"Location": "/" + config.Get().PREFIX + "/login",
			}, ``)
		},
		permissionDenyCallback: func(ctx *context.Context) {
			page.SetPageContent(ctx, Auth(ctx), func() types.Panel {
				alert := template2.Get(config.Get().THEME).Alert().SetTitle(template.HTML(`<i class="icon fa fa-warning"></i> Error!`)).
					SetTheme("warning").SetContent(template.HTML("permission denied")).GetContent()

				return types.Panel{
					Content:     alert,
					Description: "Error",
					Title:       "Error",
				}
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

	if user.Avatar == "" || config.Get().STORE.PREFIX == "" {
		user.Avatar = ""
	} else {
		user.Avatar = "/" + config.Get().STORE.PREFIX + "/" + user.Avatar
	}

	user = user.WithRoles().WithPermissions().WithMenus()

	ok = true

	return
}

func CheckPermissions(user models.UserModel, path string, method string) bool {

	prefix := "/" + config.Get().PREFIX

	if path == prefix+"/logout" {
		return true
	}

	for _, v := range user.Permissions {

		if v.HttpMethod[0] == "" || InMethodArr(v.HttpMethod, method) {

			if v.HttpPath[0] == "*" {
				return true
			}

			for i := 0; i < len(v.HttpPath); i++ {

				matchPath := ""

				if v.HttpPath[i] == "/" {
					matchPath = prefix
				} else {
					matchPath = prefix + v.HttpPath[i]
				}

				if matchPath == path {
					return true
				}

				reg, err := regexp.Compile(matchPath)

				if err != nil {
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
		if arr[i] == str {
			return true
		}
	}
	return false
}
