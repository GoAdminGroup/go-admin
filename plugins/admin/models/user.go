package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"strconv"
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
	Role          RoleModel         `json:"role"`
	Level         string            `json:"level"`
	LevelName     string            `json:"level_name"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// User return a default user model.
func User() UserModel {
	return UserModel{Base: Base{Table: "goadmin_users"}}
}

// UserWithId return a default user model of given id.
func UserWithId(id string) UserModel {
	idInt, _ := strconv.Atoi(id)
	return UserModel{Base: Base{Table: "goadmin_users"}, Id: int64(idInt)}
}

// Find return a default user model of given id.
func (t UserModel) Find(id interface{}) UserModel {
	item, _ := db.Table(t.Table).Find(id)
	return t.MapToModel(item)
}

// FindByUserName return a default user model of given name.
func (t UserModel) FindByUserName(username interface{}) UserModel {
	item, _ := db.Table(t.Table).Where("username", "=", username).First()
	return t.MapToModel(item)
}

// IsEmpty check the user model is empty or not.
func (t UserModel) IsEmpty() bool {
	return t.Id == int64(0)
}

// IsSuperAdmin check the user model is super admin or not.
func (t UserModel) IsSuperAdmin() bool {
	for _, per := range t.Permissions {
		if len(per.HttpPath) > 0 && per.HttpPath[0] == "*" {
			return true
		}
	}
	return false
}

// UpdateAvatar update the avatar of user.
func (t UserModel) UpdateAvatar(avatar string) {
	t.Avatar = avatar
}

// WithRoles query the role info of the user.
func (t UserModel) WithRoles() UserModel {
	roleModel, _ := db.Table("goadmin_role_users").
		LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
		Where("user_id", "=", t.Id).
		Select("goadmin_roles.id", "goadmin_roles.name", "goadmin_roles.slug",
			"goadmin_roles.created_at", "goadmin_roles.updated_at").
		First()

	t.Role = Role().MapToModel(roleModel)
	t.Level = roleModel["slug"].(string)
	t.LevelName = roleModel["name"].(string)
	return t
}

// WithPermissions query the permission info of the user.
func (t UserModel) WithPermissions() UserModel {

	permissions, _ := db.Table("goadmin_role_permissions").
		LeftJoin("goadmin_permissions", "goadmin_permissions.id", "=", "goadmin_role_permissions.permission_id").
		Where("role_id", "=", t.Role.Id).
		Select("goadmin_permissions.http_method", "goadmin_permissions.http_path",
			"goadmin_permissions.id", "goadmin_permissions.name", "goadmin_permissions.slug",
			"goadmin_permissions.created_at", "goadmin_permissions.updated_at").
		All()

	for i := 0; i < len(permissions); i++ {
		t.Permissions = append(t.Permissions, Permission().MapToModel(permissions[i]))
	}

	return t
}

// WithMenus query the menu info of the user.
func (t UserModel) WithMenus() UserModel {

	menuIdsModel, _ := db.Table("goadmin_role_menu").
		LeftJoin("goadmin_menu", "goadmin_menu.id", "=", "goadmin_role_menu.menu_id").
		Where("goadmin_role_menu.role_id", "=", t.Role.Id).
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

	t.MenuIds = menuIds
	return t
}

// New create a user model.
func (t UserModel) New(username, password, name, avatar string) UserModel {

	id, _ := db.Table(t.Table).Insert(dialect.H{
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

	return t
}

// Update update the user model.
func (t UserModel) Update(username, password, name, avatar string) UserModel {

	fieldValues := dialect.H{
		"username": username,
		"name":     name,
	}

	if avatar != "" {
		fieldValues["avatar"] = avatar
		t.Avatar = avatar
	}

	if password != "" {
		fieldValues["password"] = password
		t.Avatar = avatar
	}

	_, _ = db.Table(t.Table).
		Where("id", "=", t.Id).
		Update(fieldValues)

	t.UserName = username
	t.Password = password
	t.Name = name

	return t
}

// UpdatePwd update the password of the user model.
func (t UserModel) UpdatePwd(password string) UserModel {

	_, _ = db.Table(t.Table).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"password": password,
		})

	t.Password = password
	return t
}

// CheckRole check the role of the user model.
func (t UserModel) CheckRole(roleId string) bool {
	checkRole, _ := db.Table("goadmin_role_users").
		Where("role_id", "=", roleId).
		Where("user_id", "=", t.Id).
		First()
	return checkRole != nil
}

// DeleteRoles delete all the roles of the user model.
func (t UserModel) DeleteRoles() {
	_ = db.Table("goadmin_role_users").
		Where("user_id", "=", t.Id).
		Delete()
}

// AddRole add a role of the user model.
func (t UserModel) AddRole(roleId string) {
	if roleId != "" {
		if !t.CheckRole(roleId) {
			_, _ = db.Table("goadmin_role_users").
				Insert(dialect.H{
					"role_id": roleId,
					"user_id": t.Id,
				})
		}
	}
}

// CheckPermission check the permission of the user.
func (t UserModel) CheckPermission(permissionId string) bool {
	checkPermission, _ := db.Table("goadmin_user_permissions").
		Where("permission_id", "=", permissionId).
		Where("user_id", "=", t.Id).
		First()
	return checkPermission != nil
}

// DeletePermissions delete all the permissions of the user model.
func (t UserModel) DeletePermissions() {
	_ = db.Table("goadmin_user_permissions").
		Where("user_id", "=", t.Id).
		Delete()
}

// AddPermission add a permission of the user model.
func (t UserModel) AddPermission(permissionId string) {
	if permissionId != "" {
		if !t.CheckPermission(permissionId) {
			_, _ = db.Table("goadmin_user_permissions").
				Insert(dialect.H{
					"permission_id": permissionId,
					"user_id":       t.Id,
				})
		}
	}
}

// MapToModel get the user model from given map.
func (t UserModel) MapToModel(m map[string]interface{}) UserModel {
	t.Id = m["id"].(int64)
	t.Name, _ = m["name"].(string)
	t.UserName, _ = m["username"].(string)
	t.Password = m["password"].(string)
	t.Avatar, _ = m["avatar"].(string)
	t.RememberToken, _ = m["remember_token"].(string)
	t.CreatedAt, _ = m["created_at"].(string)
	t.UpdatedAt, _ = m["updated_at"].(string)
	return t
}
