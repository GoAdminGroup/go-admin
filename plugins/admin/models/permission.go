package models

import (
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/db"
)

// PermissionModel is permission model structure.
type PermissionModel struct {
	Base

	Id         int64
	Name       string
	Slug       string
	HttpMethod []string
	HttpPath   []string
	CreatedAt  string
	UpdatedAt  string
}

// Permission return a default permission model.
func Permission() PermissionModel {
	return PermissionModel{Base: Base{TableName: "goadmin_permissions"}}
}

// PermissionWithId return a default permission model of given id.
func PermissionWithId(id string) PermissionModel {
	idInt, _ := strconv.Atoi(id)
	return PermissionModel{Base: Base{TableName: "goadmin_permissions"}, Id: int64(idInt)}
}

func (t PermissionModel) SetConn(con db.Connection) PermissionModel {
	t.Conn = con
	return t
}

// IsEmpty check the user model is empty or not.
func (t PermissionModel) IsEmpty() bool {
	return t.Id == int64(0)
}

// IsSlugExist check the row exist with given slug and id.
func (t PermissionModel) IsSlugExist(slug string, id string) bool {
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

// Find return the permission model of given id.
func (t PermissionModel) Find(id interface{}) PermissionModel {
	item, _ := t.Table(t.TableName).Find(id)
	return t.MapToModel(item)
}

// FindBySlug return the permission model of given slug.
func (t PermissionModel) FindBySlug(slug string) PermissionModel {
	item, _ := t.Table(t.TableName).Where("slug", "=", slug).First()
	return t.MapToModel(item)
}

// FindBySlug return the permission model of given slug.
func (t PermissionModel) FindByName(name string) PermissionModel {
	item, _ := t.Table(t.TableName).Where("name", "=", name).First()
	return t.MapToModel(item)
}

// MapToModel get the permission model from given map.
func (t PermissionModel) MapToModel(m map[string]interface{}) PermissionModel {
	t.Id = m["id"].(int64)
	t.Name, _ = m["name"].(string)
	t.Slug, _ = m["slug"].(string)

	methods, _ := m["http_method"].(string)
	if methods != "" {
		t.HttpMethod = strings.Split(methods, ",")
	} else {
		t.HttpMethod = []string{""}
	}

	path, _ := m["http_path"].(string)
	t.HttpPath = strings.Split(path, "\n")
	t.CreatedAt, _ = m["created_at"].(string)
	t.UpdatedAt, _ = m["updated_at"].(string)
	return t
}
