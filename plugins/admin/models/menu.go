package models

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/db/dialect"
	"strconv"
)

type MenuModel struct {
	Base

	Id        int64
	Title     string
	ParentId  string
	Icon      string
	Uri       string
	CreatedAt string
	UpdatedAt string
}

func Menu() MenuModel {
	return MenuModel{Base: Base{Table: "goadmin_menu"}}
}

func MenuWithId(id string) MenuModel {
	idInt, _ := strconv.Atoi(id)
	return MenuModel{Base: Base{Table: "goadmin_menu"}, Id: int64(idInt)}
}

func (t MenuModel) Find(id interface{}) MenuModel {
	item, _ := db.Table(t.Table).Find(id)
	return t.MapToModel(item)
}

func (t MenuModel) New(title, parentId, icon, uri string, order int64) MenuModel {

	id, _ := db.Table(t.Table).Insert(dialect.H{
		"title":     title,
		"parent_id": parentId,
		"icon":      icon,
		"uri":       uri,
		"order":     order,
	})

	t.Id = id
	t.Title = title
	t.ParentId = parentId
	t.Icon = icon
	t.Uri = uri

	return t
}

func (t MenuModel) Delete() {
	_ = db.Table(t.Table).Where("id", "=", t.Id).Delete()
	_ = db.Table("goadmin_role_menu").
		Where("menu_id", "=", t.Id).Delete()
}

func (t MenuModel) Update(title, parentId, icon, uri string) MenuModel {

	_, _ = db.Table(t.Table).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"title":     title,
			"parent_id": parentId,
			"icon":      icon,
			"uri":       uri,
		})

	t.Title = title
	t.ParentId = parentId
	t.Icon = icon
	t.Uri = uri

	return t
}

func (t MenuModel) ResetOrder(data []map[string]interface{}) {
	count := 1
	for _, v := range data {
		if child, ok := v["children"]; ok {
			_, _ = db.Table(t.Table).
				Where("id", "=", v["id"]).Update(dialect.H{
				"order":     count,
				"parent_id": 0,
			})

			for _, v2 := range child.([]interface{}) {
				_, _ = db.Table(t.Table).
					Where("id", "=", v2.(map[string]interface{})["id"]).Update(dialect.H{
					"order":     count,
					"parent_id": v["id"],
				})
				count++
			}
		} else {
			_, _ = db.Table(t.Table).
				Where("id", "=", v["id"]).Update(dialect.H{
				"order":     count,
				"parent_id": 0,
			})
			count++
		}
	}
}

func (t MenuModel) CheckRole(roleId string) bool {
	checkRole, _ := db.Table("goadmin_role_menu").
		Where("role_id", "=", roleId).
		Where("menu_id", "=", t.Id).
		First()
	return checkRole != nil
}

func (t MenuModel) AddRole(roleId string) {
	if roleId != "" {
		if !t.CheckRole(roleId) {
			_, _ = db.Table("goadmin_role_menu").
				Insert(dialect.H{
					"role_id": roleId,
					"menu_id": t.Id,
				})
		}
	}
}

func (t MenuModel) MapToModel(m map[string]interface{}) MenuModel {
	t.Id = m["id"].(int64)
	t.Title = m["title"].(string)
	t.ParentId = m["parent_id"].(string)
	t.Icon = m["icon"].(string)
	t.Uri = m["uri"].(string)
	t.CreatedAt = m["created_at"].(string)
	t.UpdatedAt = m["updated_at"].(string)
	return t
}
