// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/db/dialect"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

func Auth(ctx *context.Context) User {
	return ctx.User().(User)
}

func Check(password string, username string) (user User, ok bool) {

	admin, _ := db.Table("goadmin_users").Where("username", "=", username).First()

	if admin == nil {
		ok = false
	} else {
		if ComparePassword(password, admin["password"].(string)) {
			ok = true

			roleModel, _ := db.Table("goadmin_role_users").
				LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
				Where("user_id", "=", admin["id"]).
				Select("goadmin_roles.id", "goadmin_roles.name", "goadmin_roles.slug").
				First()

			user.ID = strconv.FormatInt(admin["id"].(int64), 10)
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

			newPwd := EncodePassword([]byte(password))
			_, _ = db.Table("goadmin_users").
				Where("id", "=", user.ID).Update(dialect.H{
				"password": newPwd,
			})

		} else {
			ok = false
		}
	}
	return
}

func ComparePassword(comPwd, pwdHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(comPwd))
	if err != nil {
		return false
	} else {
		return true
	}
}

func EncodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash[:])
}

func SetCookie(ctx *context.Context, user User) bool {
	InitSession(ctx).Set("user_id", user.ID)
	return true
}

func DelCookie(ctx *context.Context) bool {
	InitSession(ctx).Clear()
	return true
}

type CSRFToken []string

var TokenHelper = new(CSRFToken)

func (token *CSRFToken) AddToken() string {
	tokenStr := modules.Uuid(35)
	if len(*token) == 1 && (*token)[0] == "" {
		(*token)[0] = tokenStr
	} else {
		*token = append(*token, tokenStr)
	}
	return tokenStr
}

func (token *CSRFToken) CheckToken(tocheck string) bool {
	for i := 0; i < len(*token); i++ {
		if (*token)[i] == tocheck {
			*token = append((*token)[0:i], (*token)[i:len(*token)]...)
			return true
		}
	}
	return false
}
