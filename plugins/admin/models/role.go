package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"strconv"
)

// RoleModel is role model structure.
type RoleModel struct {
	Base

	Id        int64
	Name      string
	Slug      string
	CreatedAt string
	UpdatedAt string
}

// Role return a default role model.
func Role() RoleModel {
	return RoleModel{Base: Base{Table: "goadmin_roles"}}
}

// RoleWithId return a default role model of given id.
func RoleWithId(id string) RoleModel {
	idInt, _ := strconv.Atoi(id)
	return RoleModel{Base: Base{Table: "goadmin_roles"}, Id: int64(idInt)}
}

// Find return a default role model of given id.
func (t RoleModel) Find(id interface{}) RoleModel {
	item, _ := db.Table(t.Table).Find(id)
	return t.MapToModel(item)
}

// New create a role model.
func (t RoleModel) New(name, slug string) RoleModel {

	id, _ := db.Table(t.Table).Insert(dialect.H{
		"name": name,
		"slug": slug,
	})

	t.Id = id
	t.Name = name
	t.Slug = slug

	return t
}

// Update update the role model.
func (t RoleModel) Update(name, slug string) RoleModel {

	_, _ = db.Table(t.Table).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"name": name,
			"slug": slug,
		})

	t.Name = name
	t.Slug = slug

	return t
}

// CheckPermission check the permission of role.
func (t RoleModel) CheckPermission(permissionId string) bool {
	checkPermission, _ := db.Table("goadmin_role_permissions").
		Where("permission_id", "=", permissionId).
		Where("role_id", "=", t.Id).
		First()
	return checkPermission != nil
}

// DeletePermissions delete all the permissions of role.
func (t RoleModel) DeletePermissions() {
	_ = db.Table("goadmin_role_permissions").
		Where("role_id", "=", t.Id).
		Delete()
}

// AddPermission add the permissions to the role.
func (t RoleModel) AddPermission(permissionId string) {
	if permissionId != "" {
		if !t.CheckPermission(permissionId) {
			_, _ = db.Table("goadmin_role_permissions").
				Insert(dialect.H{
					"permission_id": permissionId,
					"role_id":       t.Id,
				})
		}
	}
}

// MapToModel get the role model from given map.
func (t RoleModel) MapToModel(m map[string]interface{}) RoleModel {
	t.Id = m["id"].(int64)
	t.Name, _ = m["name"].(string)
	t.Slug, _ = m["slug"].(string)
	t.CreatedAt, _ = m["created_at"].(string)
	t.UpdatedAt, _ = m["updated_at"].(string)
	return t
}
