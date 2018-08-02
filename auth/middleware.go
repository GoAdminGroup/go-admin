package auth

import (
	"github.com/valyala/fasthttp"
	"goAdmin/connections/mysql"
	"regexp"
	"strings"
)

type User struct {
	ID          string
	Level       string
	Name        string
	LevelName   string
	CreateAt    string
	Permissions []Permission
}

type Permission struct {
	Method []string
	Path   []string
}

func Filter(ctx *fasthttp.RequestCtx) (User, bool, bool) {
	cookieSec := string(ctx.Request.Header.Cookie("go_admin_session")[:])
	id := InitSessionHelper(ctx).GetUserIdFromSession(cookieSec)
	user, ok := GetCurUserById(id)

	if !ok {
		return user, false, false
	}

	return user, true, CheckPermissions(user, ctx)
}

func GetCurUserById(id string) (user User, ok bool) {
	admin, _ := mysql.Query("select * from goadmin_users where id = ?", id)

	if len(admin) < 1 {
		ok = false
		return
	}

	roleModel, _ := mysql.Query("select r.id, r.name, r.slug from goadmin_role_users as u left join goadmin_roles as r on u.role_id = r.id where user_id = ?", id)

	user.ID = id
	user.Level = roleModel[0]["slug"].(string)
	user.LevelName = roleModel[0]["name"].(string)
	user.Name = admin[0]["name"].(string)
	user.CreateAt = admin[0]["created_at"].(string)

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
	permissions, _ := mysql.Query("select p.http_method, p.http_path from goadmin_role_permissions as rp left join goadmin_permissions as p on rp.permission_id = p.id where role_id = ?", role_id)
	return permissions
}

func CheckPermissions(user User, ctx *fasthttp.RequestCtx) bool {

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
