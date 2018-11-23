package auth

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

func Auth(ctx *context.Context) User {
	return ctx.User().(User)
}

func Check(password string, username string) (user User, ok bool) {

	admin, _ := db.Query("select * from goadmin_users where username = ?", username)

	if len(admin) < 1 {
		ok = false
	} else {
		if ComparePassword(password, admin[0]["password"].(string)) {
			ok = true

			roleModel, _ := db.Query("select r.id, r.name, r.slug from goadmin_role_users "+
				"as u left join goadmin_roles as r on u.role_id = r.id where user_id = ?", admin[0]["id"])

			user.ID = strconv.FormatInt(admin[0]["id"].(int64), 10)
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

			newPwd := EncodePassword([]byte(password))
			db.Exec("update goadmin_users set password = ? where id = ?", newPwd, user.ID)

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
