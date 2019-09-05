package models

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/db/dialect"
	"strconv"
)

type UserModel struct {
	Base

	Id            int64
	Name          string
	UserName      string
	Password      string
	Avatar        string
	RememberToken string
	CreatedAt     string
	UpdatedAt     string
}

func User() UserModel {
	return UserModel{Base: Base{Table: "goadmin_users"}}
}

func UserWithId(id string) UserModel {
	idInt, _ := strconv.Atoi(id)
	return UserModel{Base: Base{Table: "goadmin_users"}, Id: int64(idInt)}
}

func (t UserModel) Find(id interface{}) UserModel {
	item, _ := db.Table(t.Table).Find(id)
	return t.MapToModel(item)
}

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

func (t UserModel) Update(username, password, name, avatar string) UserModel {

	_, _ = db.Table(t.Table).
		Where("id", "=", t.Id).
		Update(dialect.H{
			"username": username,
			"password": password,
			"name":     name,
			"avatar":   avatar,
		})

	t.UserName = username
	t.Password = password
	t.Avatar = avatar
	t.Name = name

	return t
}

func (t UserModel) CheckRole(roleId string) bool {
	checkRole, _ := db.Table("goadmin_role_users").
		Where("role_id", "=", roleId).
		Where("user_id", "=", t.Id).
		First()
	return checkRole != nil
}

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

func (t UserModel) CheckPermission(permissionId string) bool {
	checkPermission, _ := db.Table("goadmin_user_permissions").
		Where("permission_id", "=", permissionId).
		Where("user_id", "=", t.Id).
		First()
	return checkPermission != nil
}

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

func (t UserModel) MapToModel(m map[string]interface{}) UserModel {
	t.Id = m["id"].(int64)
	t.Name = m["name"].(string)
	t.UserName = m["username"].(string)
	t.Password = m["password"].(string)
	t.Avatar = m["avatar"].(string)
	t.RememberToken = m["remember_token"].(string)
	t.CreatedAt = m["created_at"].(string)
	t.UpdatedAt = m["updated_at"].(string)
	return t
}
