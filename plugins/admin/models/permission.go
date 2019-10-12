package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"strconv"
	"strings"
)

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

func Permission() PermissionModel {
	return PermissionModel{Base: Base{Table: "goadmin_permissions"}}
}

func PermissionWithId(id string) PermissionModel {
	idInt, _ := strconv.Atoi(id)
	return PermissionModel{Base: Base{Table: "goadmin_permissions"}, Id: int64(idInt)}
}

func (t PermissionModel) Find(id interface{}) PermissionModel {
	item, _ := db.Table(t.Table).Find(id)
	return t.MapToModel(item)
}

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
