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

func (user User) UpdateMenus() User {
	roleModel, _ := db.Table("goadmin_role_users").
		LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
		Where("user_id", "=", user.ID).
		Select("goadmin_roles.id", "goadmin_roles.name", "goadmin_roles.slug").
		First()

	menuIdsModel, _ := db.Table("goadmin_role_menu").
		LeftJoin("goadmin_menu", "goadmin_menu.id", "=", "goadmin_role_menu.menu_id").
		Where("goadmin_role_menu.role_id", "=", roleModel["id"]).
		Select("menu_id", "parent_id").
		All()

	var menuIds []int64

	for _, mid := range menuIdsModel {
		if parentId, ok := mid["parent_id"].(int64); ok && parentId != 0 {
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

	return user
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
	admin, _ := db.Table("goadmin_users").Find(id)

	if admin == nil {
		ok = false
		return
	}

	roleModel, _ := db.Table("goadmin_role_users").
		LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
		Where("user_id", "=", id).
		Select("goadmin_roles.id", "goadmin_roles.name", "goadmin_roles.slug").
		First()

	user.ID = id
	user.Level = roleModel["slug"].(string)
	user.LevelName = roleModel["name"].(string)
	user.Name = admin["name"].(string)
	user.CreateAt = admin["created_at"].(string)
	if admin["avatar"].(string) == "" || config.Get().STORE.PREFIX == "" {
		user.Avatar = ""
	} else {
		user.Avatar = "/" + config.Get().STORE.PREFIX + "/" + admin["avatar"].(string)
	}

	// TODO: 支持多角色
	permissionModel := GetPermissions(roleModel["id"])
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

	menuIdsModel, _ := db.Table("goadmin_role_menu").
		LeftJoin("goadmin_menu", "goadmin_menu.id", "=", "goadmin_role_menu.menu_id").
		Where("goadmin_role_menu.role_id", "=", roleModel["id"]).
		Select("menu_id", "parent_id").
		All()

	var menuIds []int64

	for _, mid := range menuIdsModel {
		if parentId, ok := mid["parent_id"].(int64); ok && parentId != 0 {
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

func GetPermissions(roleId interface{}) []map[string]interface{} {
	permissions, _ := db.Table("goadmin_role_permissions").
		LeftJoin("goadmin_permissions", "goadmin_permissions.id", "=", "goadmin_role_permissions.permission_id").
		Where("role_id", "=", roleId).
		Select("goadmin_permissions.http_method", "goadmin_permissions.http_path").
		All()

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
