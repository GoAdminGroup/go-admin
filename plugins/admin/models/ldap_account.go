package models

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"strconv"
)

type LdapAccount struct {
	Ctx *Context

	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	DN       string `json:"dn"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewLdapAccount() LdapAccount {
	return LdapAccount{Ctx: &Context{}}
}

func (LdapAccount) TableName() string {
	return "goadmin_ldap_account"
}

func (m LdapAccount) WithConn(conn db.Connection) LdapAccount {
	m.Ctx.WithConn(conn)
	return m
}

func (m LdapAccount) WithTx(tx *sql.Tx) LdapAccount {
	m.Ctx.WithTx(tx)
	return m
}

func (m LdapAccount) CreateUser(username, dn string) (LdapAccount, UserModel, error) {
	var (
		ldapAccount = NewLdapAccount()
		user        = User()
		role        = Role()
		err         error
	)
	tx := m.Ctx.Conn.BeginTx()

	if user.Id, err = db.Table(user.TableName).WithDriver(m.Ctx.Conn).WithTx(tx).Insert(dialect.H{
		"name": username,
	}); err != nil {
		return ldapAccount, user, tx.Rollback()
	}
	// must have a default permission
	role = role.SetConn(m.Ctx.Conn).FindDefault()
	if role.IsEmpty() {
		return ldapAccount, user, tx.Rollback()
	}
	if _, err = user.SetConn(m.Ctx.Conn).WithTx(tx).AddRole(strconv.FormatInt(role.Id, 10)); db.CheckError(err, db.INSERT) {
		_ = tx.Rollback()
		return ldapAccount, user, err
	}
	user.Roles = []RoleModel{role}
	ldapAccount.UserId, err = db.Table(m.TableName()).WithDriver(m.Ctx.Conn).WithTx(tx).Insert(dialect.H{
		"user_id":  user.Id,
		"username": username,
		"DN":       dn,
	})
	if err != nil {
		return ldapAccount, user, tx.Rollback()
	}
	ldapAccount.Username = username
	ldapAccount.DN = dn

	return ldapAccount, user, tx.Commit()
}

func (m LdapAccount) Find(id interface{}) LdapAccount {
	item, _ := db.Table(m.TableName()).WithDriver(m.Ctx.Conn).Find(id)
	return m.MapToModel(item)
}

func (m LdapAccount) FindByUsernameAndDN(username string, dn string) LdapAccount {
	item, _ := db.Table(m.TableName()).WithDriver(m.Ctx.Conn).Where("username", "=", username).Where("DN", "=", dn).First()
	return m.MapToModel(item)
}

func (m LdapAccount) IsEmpty() bool {
	return m.UserId == int64(0)
}

// MapToModel get the user model from given map.
func (m LdapAccount) MapToModel(mm map[string]interface{}) LdapAccount {
	m.UserId, _ = mm["user_id"].(int64)
	m.Username, _ = mm["username"].(string)
	m.DN, _ = mm["dn"].(string)
	m.CreatedAt, _ = mm["created_at"].(string)
	m.UpdatedAt, _ = mm["updated_at"].(string)
	return m
}
