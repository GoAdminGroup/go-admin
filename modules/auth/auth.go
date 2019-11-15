// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

// Auth get the user model from Context.
func Auth(ctx *context.Context) models.UserModel {
	return ctx.User().(models.UserModel)
}

// Check check the password and username and return the user model.
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
	return err == nil
}

// EncodePassword encode the password.
func EncodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash[:])
}

// SetCookie set the cookie.
func SetCookie(ctx *context.Context, user models.UserModel) bool {
	InitSession(ctx).Add("user_id", user.Id)
	return true
}

// DelCookie delete the cookie from Context.
func DelCookie(ctx *context.Context) bool {
	InitSession(ctx).Clear()
	return true
}

// CSRFToken is type of a csrf token list.
type CSRFToken []string

var (
	// TokenHelper helps check the token.
	TokenHelper   = new(CSRFToken)

	// CsrfTokenLock is the a lock of checking.
	CsrfTokenLock sync.Mutex
)

// AddToken add the token to the CSRFToken.
func (csrf *CSRFToken) AddToken() string {
	CsrfTokenLock.Lock()
	defer CsrfTokenLock.Unlock()
	tokenStr := modules.Uuid()
	*csrf = append(*csrf, tokenStr)
	return tokenStr
}

// CheckToken check the given token with tokens in the CSRFToken, if exist
// return true.
func (csrf *CSRFToken) CheckToken(toCheckToken string) bool {
	for i := 0; i < len(*csrf); i++ {
		if (*csrf)[i] == toCheckToken {
			*csrf = append((*csrf)[:i], (*csrf)[i+1:]...)
			return true
		}
	}
	return false
}
