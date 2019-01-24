// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"regexp"
	"strings"
)

type User struct {
	ID          string
	Level       string
	Name        string
	LevelName   string
	CreateAt    string
	Avatar      string
	Permissions []Permission
	Menus       []int64
}

type Permission struct {
	Method []string
	Path   []string
}

type Invoker struct {
	prefix                 string
	authFailCallback       MiddlewareCallback
	permissionDenyCallback MiddlewareCallback
}

func (user User) IsSuperAdmin() bool {
	for _, per := range user.Permissions {
		if len(per.Path) > 0 && per.Path[0] == "*" {
			return true
		}
	}
	return false
}

func SetPrefix(prefix string) *Invoker {
	return &Invoker{
		prefix: prefix,
	}
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

func (invoker *Invoker) Middleware(h context.Handler) context.Handler {
	return func(ctx *context.Context) {
		user, authOk, permissionOk := Filter(ctx)

		if authOk && permissionOk {
			ctx.SetUserValue("user", user)
			h(ctx)
			return
		}

		if !authOk {
			invoker.authFailCallback(ctx)
			return
		}

		if !permissionOk {
			ctx.SetUserValue("user", user)
			invoker.permissionDenyCallback(ctx)
			return
		}
	}
}

func Filter(ctx *context.Context) (User, bool, bool) {
	var (
		id   string
		ok   bool
		user User
	)
	if id, ok = InitSession(ctx).Get("user_id").(string); !ok {
		return user, false, false
	}

	user, ok = GetCurUserById(id)

	if !ok {
		return user, false, false
	}

	return user, true, CheckPermissions(user, ctx.Path(), ctx.Method())
}

func GetCurUserById(id string) (user User, ok bool) {
	admin, _ := db.Query("select * from goadmin_users where id = ?", id)

	if len(admin) < 1 {
		ok = false
		return
	}

	roleModel, _ := db.Query("select r.id, r.name, r.slug from goadmin_role_users as u "+
		"left join goadmin_roles as r on u.role_id = r.id where user_id = ?", id)

	user.ID = id
	user.Level = roleModel[0]["slug"].(string)
	user.LevelName = roleModel[0]["name"].(string)
	user.Name = admin[0]["name"].(string)
	user.CreateAt = admin[0]["created_at"].(string)
	if admin[0]["avatar"].(string) == "" || config.Get().STORE.PREFIX == "" {
		user.Avatar = ""
	} else {
		user.Avatar = "/" + config.Get().STORE.PREFIX + "/" + admin[0]["avatar"].(string)
	}

	// TODO: 支持多角色
	permissionModel := GetPermissions(roleModel[0]["id"])
	var permissions []Permission
	for i := 0; i < len(permissionModel); i++ {

		var methodArr []string

		if permissionModel[i]["http_method"].(string) != "" {
			methodArr = strings.Split(permissionModel[i]["http_method"].(string), ",")
		} else {
			methodArr = []string{""}
		}
		permissions = append(permissions, Permission{
			methodArr,
			strings.Split(permissionModel[i]["http_path"].(string), "\n"),
		})
	}

	user.Permissions = permissions

	menuIdsModel, _ := db.Query("select menu_id, parent_id from goadmin_role_menu left join "+
		"goadmin_menu on goadmin_menu.id = goadmin_role_menu.menu_id where goadmin_role_menu.role_id = ?", roleModel[0]["id"])

	var menuIds []int64

	for _, mid := range menuIdsModel {
		if parent_id, ok := mid["parent_id"].(int64); ok && parent_id != 0 {
			for _, mid2 := range menuIdsModel {
				if mid2["menu_id"].(int64) == mid["parent_id"].(int64) {
					menuIds = append(menuIds, mid["menu_id"].(int64))
					break
				}
			}
		} else {
			menuIds = append(menuIds, mid["menu_id"].(int64))
		}
	}

	user.Menus = menuIds

	ok = true

	return
}

func GetPermissions(role_id interface{}) []map[string]interface{} {
	permissions, _ := db.Query("select p.http_method, p.http_path from goadmin_role_permissions "+
		"as rp left join goadmin_permissions as p on rp.permission_id = p.id where role_id = ?", role_id)
	return permissions
}

func CheckPermissions(user User, path string, method string) bool {

	prefix := "/" + config.Get().PREFIX

	if path == prefix+"/logout" {
		return true
	}

	for _, v := range user.Permissions {

		if v.Method[0] == "" || InMethodArr(v.Method, method) {

			if v.Path[0] == "*" {
				return true
			}

			for i := 0; i < len(v.Path); i++ {

				matchPath := ""

				if v.Path[i] == "/" {
					matchPath = prefix
				} else {
					matchPath = prefix + v.Path[i]
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
