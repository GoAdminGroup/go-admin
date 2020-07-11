package models

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
)

// MenuModel is menu model structure.
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

// Menu return a default menu model.
func Menu() MenuModel {
	return MenuModel{Base: Base{TableName: "goadmin_menu"}}
}

// MenuWithId return a default menu model of given id.
func MenuWithId(id string) MenuModel {
	idInt, _ := strconv.Atoi(id)
	return MenuModel{Base: Base{TableName: "goadmin_menu"}, Id: int64(idInt)}
}

func (t MenuModel) SetConn(con db.Connection) MenuModel {
	t.Conn = con
	return t
}

// Find return a default menu model of given id.
func (t MenuModel) Find(id interface{}) MenuModel {
	item, _ := t.Table(t.TableName).Find(id)
	return t.MapToModel(item)
}

// New create a new menu model.
func (t MenuModel) New(title, icon, uri, header, pluginName string, parentId, order int64) (MenuModel, error) {

	id, err := t.Table(t.TableName).Insert(dialect.H{
		"title":       title,
		"parent_id":   parentId,
		"icon":        icon,
		"uri":         uri,
		"order":       order,
		"header":      header,
		"plugin_name": pluginName,
	})

	t.Id = id
	t.Title = title
	t.ParentId = parentId
	t.Icon = icon
	t.Uri = uri
	t.Header = header

	return t, err
}

// Delete delete the menu model.
func (t MenuModel) Delete() {
	_ = t.Table(t.TableName).Where("id", "=", t.Id).Delete()
	_ = t.Table("goadmin_role_menu").Where("menu_id", "=", t.Id).Delete()
	items, _ := t.Table(t.TableName).Where("parent_id", "=", t.Id).All()

	if len(items) > 0 {
		ids := make([]interface{}, len(items))
		for i := 0; i < len(ids); i++ {
			ids[i] = items[i]["id"]
		}
		_ = t.Table("goadmin_role_menu").WhereIn("menu_id", ids).Delete()
	}

	_ = t.Table(t.TableName).Where("parent_id", "=", t.Id).Delete()
}

// Update update the menu model.
func (t MenuModel) Update(title, icon, uri, header, pluginName string, parentId int64) (int64, error) {
	return t.Table(t.TableName).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"title":       title,
			"parent_id":   parentId,
			"icon":        icon,
			"plugin_name": pluginName,
			"uri":         uri,
			"header":      header,
			"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
		})
}

type OrderItems []OrderItem

type OrderItem struct {
	ID       uint       `json:"id"`
	Children OrderItems `json:"children"`
}

// ResetOrder update the order of menu models.
func (t MenuModel) ResetOrder(data []byte) {

	var items OrderItems
	_ = json.Unmarshal(data, &items)

	count := 1
	for _, v := range items {
		if len(v.Children) > 0 {
			_, _ = t.Table(t.TableName).
				Where("id", "=", v.ID).
				Update(dialect.H{
					"order":     count,
					"parent_id": 0,
				})

			for _, v2 := range v.Children {
				if len(v2.Children) > 0 {

					_, _ = t.Table(t.TableName).
						Where("id", "=", v2.ID).
						Update(dialect.H{
							"order":     count,
							"parent_id": v.ID,
						})

					for _, v3 := range v2.Children {
						_, _ = t.Table(t.TableName).
							Where("id", "=", v3.ID).
							Update(dialect.H{
								"order":     count,
								"parent_id": v2.ID,
							})
						count++
					}
				} else {
					_, _ = t.Table(t.TableName).
						Where("id", "=", v2.ID).
						Update(dialect.H{
							"order":     count,
							"parent_id": v.ID,
						})
					count++
				}
			}
		} else {
			_, _ = t.Table(t.TableName).
				Where("id", "=", v.ID).
				Update(dialect.H{
					"order":     count,
					"parent_id": 0,
				})
			count++
		}
	}
}

// CheckRole check the role if has permission to get the menu.
func (t MenuModel) CheckRole(roleId string) bool {
	checkRole, _ := t.Table("goadmin_role_menu").
		Where("role_id", "=", roleId).
		Where("menu_id", "=", t.Id).
		First()
	return checkRole != nil
}

// AddRole add a role to the menu.
func (t MenuModel) AddRole(roleId string) (int64, error) {
	if roleId != "" {
		if !t.CheckRole(roleId) {
			return t.Table("goadmin_role_menu").
				Insert(dialect.H{
					"role_id": roleId,
					"menu_id": t.Id,
				})
		}
	}
	return 0, nil
}

// DeleteRoles delete roles with menu.
func (t MenuModel) DeleteRoles() error {
	return t.Table("goadmin_role_menu").
		Where("menu_id", "=", t.Id).
		Delete()
}

// MapToModel get the menu model from given map.
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
