package common

import (
	"fmt"
	"net/http"

	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gavv/httpexpect"
)

func authTest(e *httpexpect.Expect) *http.Cookie {

	printlnWithColor("Auth", "blue")
	fmt.Println("============================")

	// login: show

	printlnWithColor("login: show", "green")
	e.GET(config.Url(config.GetLoginUrl())).Expect().Status(200)
	printlnWithColor("login: empty password", "green")
	e.POST(config.Url("/signin")).WithJSON(map[string]string{
		"username": "admin",
		"password": "",
	}).Expect().Status(400)

	// login

	printlnWithColor("login", "green")
	sesID := e.POST(config.Url("/signin")).WithForm(map[string]string{
		"username": "admin",
		"password": "admin",
	}).Expect().Status(200).Cookie(auth.DefaultCookieKey).Raw()

	// logout: without login

	printlnWithColor("logout: without login", "green")
	e.GET(config.Url("/logout")).Expect().
		Status(200)

	// logout

	printlnWithColor("logout", "green")
	e.GET(config.Url("/logout")).WithCookie(auth.DefaultCookieKey, sesID.Value).Expect().
		Status(200)

	// login again

	printlnWithColor("login again", "green")
	cookie1 := e.POST(config.Url("/signin")).WithForm(map[string]string{
		"username": "admin",
		"password": "admin",
	}).Expect().Status(200).Cookie(auth.DefaultCookieKey).Raw()

	printlnWithColor("login again：restrict users from logging in at the same time", "green")
	cookie2 := e.POST(config.Url("/signin")).WithForm(map[string]string{
		"username": "admin",
		"password": "admin",
	}).Expect().Status(200).Cookie(auth.DefaultCookieKey).Raw()

	// login success

	printlnWithColor("cookie failure", "green")
	e.GET(config.Url("/")).
		WithCookie(auth.DefaultCookieKey, cookie1.Value).Expect().
		Status(200).
		Body().Contains("login")

	printlnWithColor("login success", "green")
	e.GET(config.Url("/")).
		WithCookie(auth.DefaultCookieKey, cookie2.Value).Expect().
		Status(200).
		Body().Contains("Dashboard")

	return cookie2

}
