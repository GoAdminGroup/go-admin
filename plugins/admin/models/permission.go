package models

import (
	"github.com/chenhg5/go-admin/modules/db"
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
	t.Name = m["name"].(string)
	t.Slug = m["slug"].(string)

	methods := m["http_method"].(string)
	if methods != "" {
		t.HttpMethod = strings.Split(methods, ",")
	} else {
		t.HttpMethod = []string{""}
	}

	t.HttpPath = strings.Split(m["http_path"].(string), "\n")
	t.CreatedAt = m["created_at"].(string)
	t.UpdatedAt = m["updated_at"].(string)
	return t
}
