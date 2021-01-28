// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"
	"strings"
	"sync"

	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/logger"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"golang.org/x/crypto/bcrypt"
)

type Method string

const (
	GeneralMethod = "0"
	LdapMethod    = "1"
)

type Authenticator interface {
	Authenticate(req *http.Request) (models.UserModel, error)
}

type generalAuth struct {
	conn db.Connection
}

func (auth *generalAuth) Authenticate(req *http.Request) (user models.UserModel, err error) {
	password := req.FormValue("password")
	username := req.FormValue("username")

	if strings.TrimSpace(username) == "" {
		err = ErrUserInvalidName
		return
	}
	if strings.TrimSpace(password) == "" {
		err = ErrUserInvalidPassword
		return
	}
	account := models.NewGeneralAccount().WithConn(auth.conn).FindByUsername(username)
	if account.IsEmpty() {
		err = ErrGeneralAccountNotFound
		return
	}
	if !comparePassword(password, account.Password) {
		err = ErrGeneralAccountIncorrectPassword
		return
	}
	user = models.User().SetConn(auth.conn).Find(account.UserId)
	return
}

func NewGeneralAuth(conn db.Connection) Authenticator {
	return &generalAuth{conn: conn}
}

// Auth get the user model from Context.
func Auth(ctx *context.Context) models.UserModel {
	return ctx.User().(models.UserModel)
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
	return string(hash)
}

type TokenService struct {
	tokens CSRFToken
	lock   sync.Mutex
	conn   db.Connection
}

func (s *TokenService) Name() string {
	return TokenServiceKey
}

func InitCSRFTokenSrv(conn db.Connection) (string, service.Service) {
	list, err := db.WithDriver(conn).Table("goadmin_session").
		Where("values", "=", "__csrf_token__").
		All()
	if db.CheckError(err, db.QUERY) {
		logger.Error("csrf token query from database error: ", err)
	}
	tokens := make(CSRFToken, len(list))
	for i := 0; i < len(list); i++ {
		tokens[i] = list[i]["sid"].(string)
	}
	return TokenServiceKey, &TokenService{
		tokens: tokens,
		conn:   conn,
	}
}

const (
	TokenServiceKey = "token_csrf_helper"
	ServiceKey      = "auth"
)

func GetTokenService(s interface{}) *TokenService {
	if srv, ok := s.(*TokenService); ok {
		return srv
	}
	panic("wrong service")
}

// AddToken add the token to the CSRFToken.
func (s *TokenService) AddToken() string {
	s.lock.Lock()
	defer s.lock.Unlock()
	tokenStr := modules.Uuid()
	s.tokens = append(s.tokens, tokenStr)
	_, err := db.WithDriver(s.conn).Table("goadmin_session").Insert(dialect.H{
		"sid":    tokenStr,
		"values": "__csrf_token__",
	})
	if db.CheckError(err, db.INSERT) {
		logger.Error("csrf token insert into database error: ", err)
	}
	return tokenStr
}

// CheckToken check the given token with tokens in the CSRFToken, if exist
// return true.
func (s *TokenService) CheckToken(toCheckToken string) bool {
	for i := 0; i < len(s.tokens); i++ {
		if (s.tokens)[i] == toCheckToken {
			s.tokens = append((s.tokens)[:i], (s.tokens)[i+1:]...)
			err := db.WithDriver(s.conn).Table("goadmin_session").
				Where("sid", "=", toCheckToken).
				Where("values", "=", "__csrf_token__").
				Delete()
			if db.CheckError(err, db.DELETE) {
				logger.Error("csrf token delete from database error: ", err)
			}
			return true
		}
	}
	return false
}

// CSRFToken is type of a csrf token list.
type CSRFToken []string

type Processor func(ctx *context.Context) (model models.UserModel, exist bool, msg string)

type Service struct {
	P Processor
}

func (s *Service) Name() string {
	return "auth"
}

func GetService(s interface{}) *Service {
	if srv, ok := s.(*Service); ok {
		return srv
	}
	panic("wrong service")
}

func NewService(processor Processor) *Service {
	return &Service{
		P: processor,
	}
}
