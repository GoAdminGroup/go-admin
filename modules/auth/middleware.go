// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package auth

import (
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	gocontext "context"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/page"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"

	"github.com/coreos/go-oidc"
)

/*
var (
	IssuerURL = "https://www.auth.iij.jp/op"
	ClientID  = "d96509bad50fb74fc69b4e1b8c8bd09997173b79575ab39a8768a60a06efadc3"
)
*/

// Invoker contains the callback functions which are used
// in the route middleware.
type Invoker struct {
	prefix                 string
	authFailCallback       MiddlewareCallback
	permissionDenyCallback MiddlewareCallback
	conn                   db.Connection
	// for OIDC
	provider *oidc.Provider
	verifier *oidc.IDTokenVerifier
}

// Middleware is the default auth middleware of plugins.
func Middleware(conn db.Connection) context.Handler {
	return DefaultInvoker(conn).Middleware()
}

// DefaultInvoker return a default Invoker.
func DefaultInvoker(conn db.Connection) *Invoker {
	provider, _ := oidc.NewProvider(gocontext.Background(), config.GetOIDCIssuerURL())
	verifier := provider.Verifier(&oidc.Config{ClientID: config.GetOIDCClientID()})

	return &Invoker{
		provider: provider,
		verifier: verifier,
		prefix:   config.Prefix(),
		authFailCallback: func(ctx *context.Context) {
			if ctx.Request.URL.Path == config.Url(config.GetLoginUrl()) {
				return
			}
			if ctx.Request.URL.Path == config.Url("/logout") {
				ctx.Write(302, map[string]string{
					"Location": config.Url(config.GetLoginUrl()),
				}, ``)
				return
			}
			param := ""
			if ref := ctx.Referer(); ref != "" {
				param = "?ref=" + url.QueryEscape(ref)
			}

			u := config.Url(config.GetLoginUrl() + param)
			_, err := ctx.Request.Cookie(DefaultCookieKey)
			referer := ctx.Referer()

			if (ctx.Headers(constant.PjaxHeader) == "" && ctx.Method() != "GET") ||
				err != nil ||
				referer == "" {
				ctx.Write(302, map[string]string{
					"Location": u,
				}, ``)
			} else {
				msg := language.Get("login overdue, please login again")
				ctx.HTML(http.StatusOK, `<script>
	if (typeof(swal) === "function") {
		swal({
			type: "info",
			title: "`+language.Get("login info")+`",
			text: "`+msg+`",
			showCancelButton: false,
			confirmButtonColor: "#3c8dbc",
			confirmButtonText: '`+language.Get("got it")+`',
        })
		setTimeout(function(){ location.href = "`+u+`"; }, 3000);
	} else {
		alert("`+msg+`")
		location.href = "`+u+`"
    }
</script>`)
			}
		},
		permissionDenyCallback: func(ctx *context.Context) {
			if ctx.Headers(constant.PjaxHeader) == "" && ctx.Method() != "GET" {
				ctx.JSON(http.StatusForbidden, map[string]interface{}{
					"code": http.StatusForbidden,
					"msg":  language.Get(errors.PermissionDenied),
				})
			} else {
				page.SetPageContent(ctx, Auth(ctx), func(ctx interface{}) (types.Panel, error) {
					return template2.WarningPanel(errors.PermissionDenied, template2.NoPermission403Page), nil
				}, conn)
			}
		},
		conn: conn,
	}
}

// SetPrefix return the default Invoker with the given prefix.
func SetPrefix(prefix string, conn db.Connection) *Invoker {
	i := DefaultInvoker(conn)
	i.prefix = prefix
	return i
}

// SetAuthFailCallback set the authFailCallback of Invoker.
func (invoker *Invoker) SetAuthFailCallback(callback MiddlewareCallback) *Invoker {
	invoker.authFailCallback = callback
	return invoker
}

// SetPermissionDenyCallback set the permissionDenyCallback of Invoker.
func (invoker *Invoker) SetPermissionDenyCallback(callback MiddlewareCallback) *Invoker {
	invoker.permissionDenyCallback = callback
	return invoker
}

// MiddlewareCallback is type of callback function.
type MiddlewareCallback func(ctx *context.Context)

// Middleware get the auth middleware from Invoker.
func (invoker *Invoker) Middleware() context.Handler {
	return func(ctx *context.Context) {
		user, authOk, permissionOk := invoker.Filter(ctx, invoker.conn)

		if authOk && permissionOk {
			ctx.SetUserValue("user", user)
			ctx.Next()
			return
		}

		if !authOk {
			invoker.authFailCallback(ctx)
			ctx.Abort()
			return
		}

		if !permissionOk {
			ctx.SetUserValue("user", user)
			invoker.permissionDenyCallback(ctx)
			ctx.Abort()
			return
		}
	}
}

// Filter retrieve the user model from Context and check the permission
// at the same time.
func (invoker *Invoker) Filter(ctx *context.Context, conn db.Connection) (models.UserModel, bool, bool) {
	var (
		id float64
		ok bool

		user     = models.User()
		ses, err = InitSession(ctx, conn)
	)

	if err != nil {
		logger.Error("retrieve auth user failed", err)
		return user, false, false
	}

	// すでにセッションクッキーからIDを取得できるときは、そのIDを使う
	if id, ok = ses.Get("user_id").(float64); ok {
		if user, ok = GetCurUserByID(int64(id), conn); ok {
			return user, true, CheckPermissions(user, ctx.Request.URL.String(), ctx.Method(), ctx.PostForm())
		}
	}

	// まだセッションがなくても、AuthorizationヘッダでIDトークンを入手できれば使う
	authorization := ctx.Request.Header.Get("Authorization")
	if !strings.HasPrefix(strings.ToLower(authorization), "bearer ") {
		logger.Info("authorization header invalid: ", authorization)
		return user, false, false
	}

	// IDトークンが失効しているなど、不正であればエラー
	rawIDtoken := authorization[7:]
	idtoken, err := invoker.verifier.Verify(gocontext.Background(), rawIDtoken)
	if err != nil {
		logger.Info("idtoken verification failed: ", rawIDtoken)
		return user, false, false
	}

	logger.Infof("authenticated user: %s", idtoken.Subject)

	// goadmin_usersテーブルのnameとIDトークンのsubjectが一致することを期待している
	user.Conn = conn
	user = user.FindByUserName(idtoken.Subject)

	if user.Id == 0 {
		// Authorizationヘッダに正常なIDトークンが指定されたが、対応するユーザーが存在しない
		// 自動的に goadmin_users と users にアカウントを作成する
		if user, err = createUser(user, idtoken); err != nil {
			logger.Info("unknown user and creatin failed: ", idtoken.Subject)
			return user, false, false
		}
	}

	if user.Id == 0 {
		logger.Info("unknown user: ", idtoken.Subject)
		return user, false, false
	}

	user = user.WithRoles().WithPermissions().WithMenus()

	// すべて正常なのでセッションを作成してクッキーに保存する
	if err := SetCookie(ctx, user, conn); err != nil {
		logger.Error("set cookie failed: ", err)
		return user, false, false
	}

	return user, true, CheckPermissions(user, ctx.Request.URL.String(), ctx.Method(), ctx.PostForm())
}

// ベリファイ済のIDTokenに基づいてgoadmin userを作成する
func createUser(user models.UserModel, idtoken *oidc.IDToken) (models.UserModel, error) {
	var claims struct {
		Email             string `json:"email"`
		PrefferedUsername string `json:"preferred_username"`
	}

	if err := idtoken.Claims(&claims); err != nil {
		return user, err
	}

	// どのクレームを取得できるかわからないので、以下の優先順位であるものを使う
	nickname := ""
	if len(claims.Email) > 0 {
		nickname = claims.Email
	} else if len(claims.PrefferedUsername) > 0 {
		nickname = claims.PrefferedUsername
	} else {
		nickname = idtoken.Subject
	}

	// goadmin_usersにアカウントを作成する。パスワードは使わないのでランダムな文字列を設定する
	// avatarは面倒なので設定しない
	u, err := user.New(idtoken.Subject, generatePassword(40), nickname, "")
	if err != nil {
		return u, err
	}

	// ToDo: 管理職であれば、ChargecodeOwner, ProjectOwner, GroupOwnerロールを設定する。暫定的に全員につける
	// ToDo: SlugからRoleIDを検索するとエラーになってしまうため、いったんIDを直書きしている。あとで調査
	if err := bindRole(u, []string{"6", "3", "4", "5"}); err != nil {
		// LastInsertId is not supported ?
		// 先頭の "6" （user role）を設定した後、次の"3"を設定するところでエラーになる
		logger.Error("bind role failed", err)
		return u, err
	}

	return u, err
}

func bindRole(u models.UserModel, roles []string) error {
	for _, role := range roles {
		if _, err := u.AddRole(role); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("!@#$%&*0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generatePassword(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

const defaultUserIDSesKey = "user_id"

// GetUserID return the user id from the session.
func GetUserID(sesKey string, conn db.Connection) int64 {
	id, err := GetSessionByKey(sesKey, defaultUserIDSesKey, conn)
	if err != nil {
		logger.Error("retrieve auth user failed", err)
		return -1
	}
	if idFloat64, ok := id.(float64); ok {
		return int64(idFloat64)
	}
	return -1
}

// GetCurUser return the user model.
func GetCurUser(sesKey string, conn db.Connection) (user models.UserModel, ok bool) {

	if sesKey == "" {
		ok = false
		return
	}

	id := GetUserID(sesKey, conn)
	if id == -1 {
		ok = false
		return
	}
	return GetCurUserByID(id, conn)
}

// GetCurUserByID return the user model of given user id.
func GetCurUserByID(id int64, conn db.Connection) (user models.UserModel, ok bool) {

	user = models.User().SetConn(conn).Find(id)

	if user.IsEmpty() {
		ok = false
		return
	}

	if user.Avatar == "" || config.GetStore().Prefix == "" {
		user.Avatar = ""
	} else {
		user.Avatar = config.GetStore().URL(user.Avatar)
	}

	user = user.WithRoles().WithPermissions().WithMenus()

	ok = user.HasMenu()

	return
}

// CheckPermissions check the permission of the user.
func CheckPermissions(user models.UserModel, path, method string, param url.Values) bool {
	return user.CheckPermissionByUrlMethod(path, method, param)
}
