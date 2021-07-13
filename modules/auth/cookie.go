package auth

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
)

type CookieManager interface {
	SetCookie(ctx *context.Context, user models.UserModel) error
	DelCookie(ctx *context.Context) error
}

type cookieManager struct {
	conn db.Connection
}

func (c *cookieManager) SetCookie(ctx *context.Context, user models.UserModel) error {
	ses, err := InitSession(ctx, c.conn)
	if err != nil {
		return err
	}
	return ses.Add("user_id", user.Id)
}

func (c *cookieManager) DelCookie(ctx *context.Context) error {
	ses, err := InitSession(ctx, c.conn)
	if err != nil {
		return err
	}
	return ses.Clear()
}

func NewCookieManger(conn db.Connection) CookieManager {
	return &cookieManager{conn: conn}
}
