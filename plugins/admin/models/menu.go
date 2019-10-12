package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"strconv"
)

type MenuModel struct {
	Base

	Id        int64
	Title     string
	ParentId  int64
	Icon      string
	Uri       string
	Header    string
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

func (t MenuModel) New(title, icon, uri, header string, parentId, order int64) MenuModel {

	id, _ := db.Table(t.Table).Insert(dialect.H{
		"title":     title,
		"parent_id": parentId,
		"icon":      icon,
		"uri":       uri,
		"order":     order,
		"header":    header,
	})

	t.Id = id
	t.Title = title
	t.ParentId = parentId
	t.Icon = icon
	t.Uri = uri
	t.Header = header

	return t
}

func (t MenuModel) Delete() {
	_ = db.Table(t.Table).Where("id", "=", t.Id).Delete()
	_ = db.Table("goadmin_role_menu").Where("menu_id", "=", t.Id).Delete()
	items, _ := db.Table(t.Table).Where("parent_id", "=", t.Id).All()

	if len(items) > 0 {
		ids := make([]interface{}, len(items))
		for i := 0; i < len(ids); i++ {
			ids[i] = items[i]["id"]
		}
		_ = db.Table("goadmin_role_menu").WhereIn("menu_id", ids).Delete()
	}

	_ = db.Table(t.Table).Where("parent_id", "=", t.Id).Delete()
}

func (t MenuModel) Update(title, icon, uri, header string, parentId int64) MenuModel {

	_, _ = db.Table(t.Table).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"title":     title,
			"parent_id": parentId,
			"icon":      icon,
			"uri":       uri,
			"header":    header,
		})

	t.Title = title
	t.ParentId = parentId
	t.Icon = icon
	t.Uri = uri
	t.Header = header

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

func (t MenuModel) DeleteRoles() {
	_ = db.Table("goadmin_role_menu").
		Where("menu_id", "=", t.Id).
		Delete()
}

func (t MenuModel) MapToModel(m map[string]interface{}) MenuModel {
	t.Id = m["id"].(int64)
	t.Title, _ = m["title"].(string)
	t.ParentId = m["parent_id"].(int64)
	t.Icon, _ = m["icon"].(string)
	t.Uri, _ = m["uri"].(string)
	t.Header, _ = m["header"].(string)
	t.CreatedAt, _ = m["created_at"].(string)
	t.UpdatedAt, _ = m["updated_at"].(string)
	return t
}
