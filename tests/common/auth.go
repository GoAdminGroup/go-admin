package common

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/gavv/httpexpect"
	"net/http"
)

func authTest(e *httpexpect.Expect) *http.Cookie {

	printlnWithColor("Auth", "blue")
	fmt.Println("============================")

	// login: show

	printlnWithColor("login: show", "green")
	e.GET(config.Get().Url("/login")).Expect().Status(200)
	printlnWithColor("login: empty password", "green")
	e.POST(config.Get().Url("/signin")).WithJSON(map[string]string{
		"username": "admin",
		"password": "",
	}).Expect().Status(400)

	// login

	printlnWithColor("login", "green")
	sesID := e.POST(config.Get().Url("/signin")).WithForm(map[string]string{
		"username": "admin",
		"password": "admin",
	}).Expect().Status(200).Cookie(auth.DefaultCookieKey).Raw()

	// logout: without login

	printlnWithColor("logout: without login", "green")
	e.GET(config.Get().Url("/logout")).Expect().
		Status(200)

	// logout

	printlnWithColor("logout", "green")
	e.GET(config.Get().Url("/logout")).WithCookie(auth.DefaultCookieKey, sesID.Value).Expect().
		Status(200)

	// login again

	printlnWithColor("login again", "green")
	cookie1 := e.POST(config.Get().Url("/signin")).WithForm(map[string]string{
		"username": "admin",
		"password": "admin",
	}).Expect().Status(200).Cookie(auth.DefaultCookieKey).Raw()

	printlnWithColor("login again：restrict users from logging in at the same time", "green")
	cookie2 := e.POST(config.Get().Url("/signin")).WithForm(map[string]string{
		"username": "admin",
		"password": "admin",
	}).Expect().Status(200).Cookie(auth.DefaultCookieKey).Raw()

	// login success

	printlnWithColor("cookie failure", "green")
	e.GET(config.Get().Url("/")).
		WithCookie(auth.DefaultCookieKey, cookie1.Value).Expect().
		Status(200).
		Body().Contains("login")

	printlnWithColor("login success", "green")
	e.GET(config.Get().Url("/")).
		WithCookie(auth.DefaultCookieKey, cookie2.Value).Expect().
		Status(200).
		Body().Contains("Dashboard")

	return cookie2

}
