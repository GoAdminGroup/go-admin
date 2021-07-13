package models

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/db"
)

type GeneralAccount struct {
	Ctx *Context

	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewGeneralAccount() GeneralAccount {
	return GeneralAccount{Ctx: &Context{}}
}

func (GeneralAccount) TableName() string {
	return "goadmin_general_account"
}

func (m GeneralAccount) WithConn(conn db.Connection) GeneralAccount {
	m.Ctx.WithConn(conn)
	return m
}

func (m GeneralAccount) WithTx(tx *sql.Tx) GeneralAccount {
	m.Ctx.WithTx(tx)
	return m
}

func (m GeneralAccount) Find(id interface{}) GeneralAccount {
	item, _ := db.Table(m.TableName()).WithDriver(m.Ctx.Conn).Find(id)
	return m.MapToModel(item)
}

func (m GeneralAccount) FindByUsername(username interface{}) GeneralAccount {
	item, _ := db.Table(m.TableName()).WithDriver(m.Ctx.Conn).Where("username", "=", username).First()
	return m.MapToModel(item)
}

func (m GeneralAccount) IsEmpty() bool {
	return m.UserId == int64(0)
}

// MapToModel get the user model from given map.
func (m GeneralAccount) MapToModel(mm map[string]interface{}) GeneralAccount {
	m.UserId, _ = mm["user_id"].(int64)
	m.Username, _ = mm["username"].(string)
	m.Password, _ = mm["password"].(string)
	m.CreatedAt, _ = mm["created_at"].(string)
	m.UpdatedAt, _ = mm["updated_at"].(string)
	return m
}
