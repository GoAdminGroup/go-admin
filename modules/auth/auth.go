// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"golang.org/x/crypto/bcrypt"
)

func Auth(ctx *context.Context) models.UserModel {
	return ctx.User().(models.UserModel)
}

func Check(password string, username string) (user models.UserModel, ok bool) {

	user = models.User().FindByUserName(username)

	if user.IsEmpty() {
		ok = false
	} else {
		if comparePassword(password, user.Password) {
			ok = true
			user = user.WithRoles().WithPermissions().WithMenus()
			user.UpdatePwd(EncodePassword([]byte(password)))
		} else {
			ok = false
		}
	}
	return
}

func comparePassword(comPwd, pwdHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(comPwd))
	if err != nil {
		return false
	} else {
		return true
	}
}

func EncodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash[:])
}

func SetCookie(ctx *context.Context, user models.UserModel) bool {
	InitSession(ctx).Set("user_id", user.Id)
	return true
}

func DelCookie(ctx *context.Context) bool {
	InitSession(ctx).Clear()
	return true
}

type CSRFToken []string

var TokenHelper = new(CSRFToken)

func (token *CSRFToken) AddToken() string {
	tokenStr := modules.Uuid()
	if len(*token) == 1 && (*token)[0] == "" {
		(*token)[0] = tokenStr
	} else {
		*token = append(*token, tokenStr)
	}
	return tokenStr
}

func (token *CSRFToken) CheckToken(toCheckToken string) bool {
	for i := 0; i < len(*token); i++ {
		if (*token)[i] == toCheckToken {
			*token = append((*token)[0:i], (*token)[i:len(*token)]...)
			return true
		}
	}
	return false
}
