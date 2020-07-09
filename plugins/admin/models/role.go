package models

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
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
	return RoleModel{Base: Base{TableName: "goadmin_roles"}}
}

// RoleWithId return a default role model of given id.
func RoleWithId(id string) RoleModel {
	idInt, _ := strconv.Atoi(id)
	return RoleModel{Base: Base{TableName: "goadmin_roles"}, Id: int64(idInt)}
}

func (t RoleModel) SetConn(con db.Connection) RoleModel {
	t.Conn = con
	return t
}

func (t RoleModel) WithTx(tx *sql.Tx) RoleModel {
	t.Tx = tx
	return t
}

// Find return a default role model of given id.
func (t RoleModel) Find(id interface{}) RoleModel {
	item, _ := t.Table(t.TableName).Find(id)
	return t.MapToModel(item)
}

// IsSlugExist check the row exist with given slug and id.
func (t RoleModel) IsSlugExist(slug string, id string) bool {
	if id == "" {
		check, _ := t.Table(t.TableName).Where("slug", "=", slug).First()
		return check != nil
	}
	check, _ := t.Table(t.TableName).
		Where("slug", "=", slug).
		Where("id", "!=", id).
		First()
	return check != nil
}

// New create a role model.
func (t RoleModel) New(name, slug string) (RoleModel, error) {

	id, err := t.WithTx(t.Tx).Table(t.TableName).Insert(dialect.H{
		"name": name,
		"slug": slug,
	})

	t.Id = id
	t.Name = name
	t.Slug = slug

	return t, err
}

// Update update the role model.
func (t RoleModel) Update(name, slug string) (int64, error) {

	return t.WithTx(t.Tx).Table(t.TableName).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"name":       name,
			"slug":       slug,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
}

// CheckPermission check the permission of role.
func (t RoleModel) CheckPermission(permissionId string) bool {
	checkPermission, _ := t.Table("goadmin_role_permissions").
		Where("permission_id", "=", permissionId).
		Where("role_id", "=", t.Id).
		First()
	return checkPermission != nil
}

// DeletePermissions delete all the permissions of role.
func (t RoleModel) DeletePermissions() error {
	return t.WithTx(t.Tx).Table("goadmin_role_permissions").
		Where("role_id", "=", t.Id).
		Delete()
}

// AddPermission add the permissions to the role.
func (t RoleModel) AddPermission(permissionId string) (int64, error) {
	if permissionId != "" {
		if !t.CheckPermission(permissionId) {
			return t.WithTx(t.Tx).Table("goadmin_role_permissions").
				Insert(dialect.H{
					"permission_id": permissionId,
					"role_id":       t.Id,
				})
		}
	}
	return 0, nil
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
