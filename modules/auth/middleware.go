package auth

import (
	"github.com/chenhg5/go-admin/modules/connections"
	"regexp"
	"strings"
	"github.com/chenhg5/go-admin/context"
	"fmt"
)

type User struct {
	ID          string
	Level       string
	Name        string
	LevelName   string
	CreateAt    string
	Avatar      string
	Permissions []Permission
}

type Permission struct {
	Method []string
	Path   []string
}

func AuthMiddleware(h context.Handler) context.Handler {
	return func(ctx *context.Context) {
		var (
			authOk   bool
			filterOk bool
			user     User
		)

		if user, authOk, filterOk = Filter(ctx); authOk && filterOk {
			ctx.SetUserValue("user", user)
			h(ctx)
		}

		if !authOk {
			ctx.Write(302, map[string]string{
				"Location": "/login",
			}, ``)
		}

		if !filterOk {
			fmt.Println("filterOk Path: ", ctx.Path(), "method:", ctx.Method())
			ctx.Write(403, map[string]string{}, `{"code":403, "msg":"权限不够"}`)
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

	return user, true, CheckPermissions(user, ctx)
}

func GetCurUserById(id string) (user User, ok bool) {
	admin, _ := connections.GetConnection().Query("select * from goadmin_users where id = ?", id)

	if len(admin) < 1 {
		ok = false
		return
	}

	roleModel, _ := connections.GetConnection().Query("select r.id, r.name, r.slug from goadmin_role_users as u left join goadmin_roles as r on u.role_id = r.id where user_id = ?", id)

	user.ID = id
	user.Level = roleModel[0]["slug"].(string)
	user.LevelName = roleModel[0]["name"].(string)
	user.Name = admin[0]["name"].(string)
	user.CreateAt = admin[0]["created_at"].(string)
	user.Avatar = admin[0]["avatar"].(string)

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

	ok = true

	return
}

func GetPermissions(role_id interface{}) []map[string]interface{} {
	permissions, _ := connections.GetConnection().Query("select p.http_method, p.http_path from goadmin_role_permissions as rp left join goadmin_permissions as p on rp.permission_id = p.id where role_id = ?", role_id)
	return permissions
}

func CheckPermissions(user User, ctx *context.Context) bool {

	path := string(ctx.Path())
	method := string(ctx.Method())

	for _, v := range user.Permissions {

		if v.Method[0] == "" || InMethodArr(v.Method, method) {

			if v.Path[0] == "*" {
				return true
			}

			for i := 0; i < len(v.Path); i++ {
				if v.Path[i] == path {
					return true
				}

				if ok, err := regexp.Match(v.Path[i], []byte(path)); ok && err == nil {
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