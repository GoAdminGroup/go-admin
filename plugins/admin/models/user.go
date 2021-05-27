package models

import (
	"database/sql"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
)

// UserModel is user model structure.
type UserModel struct {
	Base `json:"-"`

	Id            int64             `json:"id"`
	Name          string            `json:"name"`
	UserName      string            `json:"user_name"`
	Password      string            `json:"password"`
	Avatar        string            `json:"avatar"`
	RememberToken string            `json:"remember_token"`
	Permissions   []PermissionModel `json:"permissions"`
	MenuIds       []int64           `json:"menu_ids"`
	Roles         []RoleModel       `json:"role"`
	Level         string            `json:"level"`
	LevelName     string            `json:"level_name"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	cacheReplacer *strings.Replacer
}

// User return a default user model.
func User() UserModel {
	return UserModel{Base: Base{TableName: config.GetAuthUserTable()}}
}

// UserWithId return a default user model of given id.
func UserWithId(id string) UserModel {
	idInt, _ := strconv.Atoi(id)
	return UserModel{Base: Base{TableName: config.GetAuthUserTable()}, Id: int64(idInt)}
}

func (t UserModel) SetConn(con db.Connection) UserModel {
	t.Conn = con
	return t
}

func (t UserModel) WithTx(tx *sql.Tx) UserModel {
	t.Tx = tx
	return t
}

// Find return a default user model of given id.
func (t UserModel) Find(id interface{}) UserModel {
	item, _ := t.Table(t.TableName).Find(id)
	return t.MapToModel(item)
}

// FindByUserName return a default user model of given name.
func (t UserModel) FindByUserName(username interface{}) UserModel {
	item, _ := t.Table(t.TableName).Where("username", "=", username).First()
	return t.MapToModel(item)
}

// IsEmpty check the user model is empty or not.
func (t UserModel) IsEmpty() bool {
	return t.Id == int64(0)
}

// HasMenu check the user has visitable menu or not.
func (t UserModel) HasMenu() bool {
	return len(t.MenuIds) != 0 || t.IsSuperAdmin()
}

// IsSuperAdmin check the user model is super admin or not.
func (t UserModel) IsSuperAdmin() bool {
	for _, per := range t.Permissions {
		if len(per.HttpPath) > 0 && per.HttpPath[0] == "*" && per.HttpMethod[0] == "" {
			return true
		}
	}
	return false
}

func (t UserModel) GetCheckPermissionByUrlMethod(path, method string) string {
	if !t.CheckPermissionByUrlMethod(path, method, url.Values{}) {
		return ""
	}
	return path
}

func (t UserModel) IsVisitor() bool {
	return !t.CheckPermissionByUrlMethod(config.Url("/info/normal_manager"), "GET", url.Values{})
}

func (t UserModel) HideUserCenterEntrance() bool {
	return t.IsVisitor() && config.GetHideVisitorUserCenterEntrance()
}

func (t UserModel) Template(str string) string {
	if t.cacheReplacer == nil {
		t.cacheReplacer = strings.NewReplacer("{{.AuthId}}", strconv.Itoa(int(t.Id)),
			"{{.AuthName}}", t.Name, "{{.AuthUserName}}", t.UserName)
	}
	return t.cacheReplacer.Replace(str)
}

func (t UserModel) CheckPermissionByUrlMethod(path, method string, formParams url.Values) bool {

	// path, _ = url.PathUnescape(path)

	if t.IsSuperAdmin() {
		return true
	}

	if path == "" {
		return false
	}

	logoutCheck, _ := regexp.Compile(config.Url("/logout") + "(.*?)")

	if logoutCheck.MatchString(path) {
		return true
	}

	if path != "/" && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	path = utils.ReplaceAll(path, constant.EditPKKey, "id", constant.DetailPKKey, "id")

	path, params := getParam(path)
	for key, value := range formParams {
		if len(value) > 0 {
			params.Add(key, value[0])
		}
	}

	for _, v := range t.Permissions {

		if v.HttpMethod[0] == "" || inMethodArr(v.HttpMethod, method) {

			if v.HttpPath[0] == "*" {
				return true
			}

			for i := 0; i < len(v.HttpPath); i++ {

				matchPath := config.Url(t.Template(strings.TrimSpace(v.HttpPath[i])))
				matchPath, matchParam := getParam(matchPath)

				if matchPath == path {
					if t.checkParam(params, matchParam) {
						return true
					}
				}

				reg, err := regexp.Compile(matchPath)

				if err != nil {
					logger.Error("CheckPermissions error: ", err)
					continue
				}

				if reg.FindString(path) == path {
					if t.checkParam(params, matchParam) {
						return true
					}
				}
			}
		}
	}

	return false
}

func getParam(u string) (string, url.Values) {
	m := make(url.Values)
	urr := strings.Split(u, "?")
	if len(urr) > 1 {
		m, _ = url.ParseQuery(urr[1])
	}
	return urr[0], m
}

func (t UserModel) checkParam(src, comp url.Values) bool {
	if len(comp) == 0 {
		return true
	}
	if len(src) == 0 {
		return false
	}
	for key, value := range comp {
		v, find := src[key]
		if !find {
			return false
		}
		if len(value) == 0 {
			continue
		}
		if len(v) == 0 {
			return false
		}
		for i := 0; i < len(v); i++ {
			if v[i] == t.Template(value[i]) {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func inMethodArr(arr []string, str string) bool {
	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(arr[i], str) {
			return true
		}
	}
	return false
}

// UpdateAvatar update the avatar of user.
func (t UserModel) ReleaseConn() UserModel {
	t.Conn = nil
	return t
}

// UpdateAvatar update the avatar of user.
func (t UserModel) UpdateAvatar(avatar string) {
	t.Avatar = avatar
}

// WithRoles query the role info of the user.
func (t UserModel) WithRoles() UserModel {
	roleModel, _ := t.Table("goadmin_role_users").
		LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
		Where("user_id", "=", t.Id).
		Select("goadmin_roles.id", "goadmin_roles.name", "goadmin_roles.slug",
			"goadmin_roles.created_at", "goadmin_roles.updated_at").
		All()

	for _, role := range roleModel {
		t.Roles = append(t.Roles, Role().MapToModel(role))
	}

	if len(t.Roles) > 0 {
		t.Level = t.Roles[0].Slug
		t.LevelName = t.Roles[0].Name
	}

	return t
}

func (t UserModel) GetAllRoleId() []interface{} {

	var ids = make([]interface{}, len(t.Roles))

	for key, role := range t.Roles {
		ids[key] = role.Id
	}

	return ids
}

// WithPermissions query the permission info of the user.
func (t UserModel) WithPermissions() UserModel {

	var permissions = make([]map[string]interface{}, 0)

	roleIds := t.GetAllRoleId()

	if len(roleIds) > 0 {
		permissions, _ = t.Table("goadmin_role_permissions").
			LeftJoin("goadmin_permissions", "goadmin_permissions.id", "=", "goadmin_role_permissions.permission_id").
			WhereIn("role_id", roleIds).
			Select("goadmin_permissions.http_method", "goadmin_permissions.http_path",
				"goadmin_permissions.id", "goadmin_permissions.name", "goadmin_permissions.slug",
				"goadmin_permissions.created_at", "goadmin_permissions.updated_at").
			All()
	}

	userPermissions, _ := t.Table("goadmin_user_permissions").
		LeftJoin("goadmin_permissions", "goadmin_permissions.id", "=", "goadmin_user_permissions.permission_id").
		Where("user_id", "=", t.Id).
		Select("goadmin_permissions.http_method", "goadmin_permissions.http_path",
			"goadmin_permissions.id", "goadmin_permissions.name", "goadmin_permissions.slug",
			"goadmin_permissions.created_at", "goadmin_permissions.updated_at").
		All()

	permissions = append(permissions, userPermissions...)

	for i := 0; i < len(permissions); i++ {
		exist := false
		for j := 0; j < len(t.Permissions); j++ {
			if t.Permissions[j].Id == permissions[i]["id"] {
				exist = true
				break
			}
		}
		if exist {
			continue
		}
		t.Permissions = append(t.Permissions, Permission().MapToModel(permissions[i]))
	}

	return t
}

// WithMenus query the menu info of the user.
func (t UserModel) WithMenus() UserModel {

	var menuIdsModel []map[string]interface{}

	if t.IsSuperAdmin() {
		menuIdsModel, _ = t.Table("goadmin_role_menu").
			LeftJoin("goadmin_menu", "goadmin_menu.id", "=", "goadmin_role_menu.menu_id").
			Select("menu_id", "parent_id").
			All()
	} else {
		rolesId := t.GetAllRoleId()
		if len(rolesId) > 0 {
			menuIdsModel, _ = t.Table("goadmin_role_menu").
				LeftJoin("goadmin_menu", "goadmin_menu.id", "=", "goadmin_role_menu.menu_id").
				WhereIn("goadmin_role_menu.role_id", rolesId).
				Select("menu_id", "parent_id").
				All()
		}
	}

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

	t.MenuIds = menuIds
	return t
}

// New create a user model.
func (t UserModel) New(username, password, name, avatar string) (UserModel, error) {

	id, err := t.WithTx(t.Tx).Table(t.TableName).Insert(dialect.H{
		"username": username,
		"password": password,
		"name":     name,
		"avatar":   avatar,
	})

	t.Id = id
	t.UserName = username
	t.Password = password
	t.Avatar = avatar
	t.Name = name

	return t, err
}

// Update update the user model.
func (t UserModel) Update(username, password, name, avatar string, isUpdateAvatar bool) (int64, error) {

	fieldValues := dialect.H{
		"username":   username,
		"name":       name,
		"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	if avatar == "" || isUpdateAvatar {
		fieldValues["avatar"] = avatar
	}

	if password != "" {
		fieldValues["password"] = password
	}

	return t.WithTx(t.Tx).Table(t.TableName).
		Where("id", "=", t.Id).
		Update(fieldValues)
}

// UpdatePwd update the password of the user model.
func (t UserModel) UpdatePwd(password string) UserModel {

	_, _ = t.Table(t.TableName).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"password": password,
		})

	t.Password = password
	return t
}

// CheckRole check the role of the user model.
func (t UserModel) CheckRoleId(roleId string) bool {
	checkRole, _ := t.Table("goadmin_role_users").
		Where("role_id", "=", roleId).
		Where("user_id", "=", t.Id).
		First()
	return checkRole != nil
}

// DeleteRoles delete all the roles of the user model.
func (t UserModel) DeleteRoles() error {
	return t.Table("goadmin_role_users").
		Where("user_id", "=", t.Id).
		Delete()
}

// AddRole add a role of the user model.
func (t UserModel) AddRole(roleId string) (int64, error) {
	if roleId != "" {
		if !t.CheckRoleId(roleId) {
			return t.WithTx(t.Tx).Table("goadmin_role_users").
				Insert(dialect.H{
					"role_id": roleId,
					"user_id": t.Id,
				})
		}
	}
	return 0, nil
}

// CheckRole check the role of the user.
func (t UserModel) CheckRole(slug string) bool {
	for _, role := range t.Roles {
		if role.Slug == slug {
			return true
		}
	}

	return false
}

// CheckPermission check the permission of the user.
func (t UserModel) CheckPermissionById(permissionId string) bool {
	checkPermission, _ := t.Table("goadmin_user_permissions").
		Where("permission_id", "=", permissionId).
		Where("user_id", "=", t.Id).
		First()
	return checkPermission != nil
}

// CheckPermission check the permission of the user.
func (t UserModel) CheckPermission(permission string) bool {
	for _, per := range t.Permissions {
		if per.Slug == permission {
			return true
		}
	}

	return false
}

// DeletePermissions delete all the permissions of the user model.
func (t UserModel) DeletePermissions() error {
	return t.WithTx(t.Tx).Table("goadmin_user_permissions").
		Where("user_id", "=", t.Id).
		Delete()
}

// AddPermission add a permission of the user model.
func (t UserModel) AddPermission(permissionId string) (int64, error) {
	if permissionId != "" {
		if !t.CheckPermissionById(permissionId) {
			return t.WithTx(t.Tx).Table("goadmin_user_permissions").
				Insert(dialect.H{
					"permission_id": permissionId,
					"user_id":       t.Id,
				})
		}
	}
	return 0, nil
}

// MapToModel get the user model from given map.
func (t UserModel) MapToModel(m map[string]interface{}) UserModel {
	t.Id, _ = m["id"].(int64)
	t.Name, _ = m["name"].(string)
	t.UserName, _ = m["username"].(string)
	t.Password, _ = m["password"].(string)
	t.Avatar, _ = m["avatar"].(string)
	t.RememberToken, _ = m["remember_token"].(string)
	t.CreatedAt, _ = m["created_at"].(string)
	t.UpdatedAt, _ = m["updated_at"].(string)
	return t
}
