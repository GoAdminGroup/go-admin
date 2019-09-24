package common

import (
	"fmt"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/gavv/httpexpect"
	"net/http"
)

func AuthTest(e *httpexpect.Expect) *http.Cookie {

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
	sesId := e.POST(config.Get().Url("/signin")).WithForm(map[string]string{
		"username": "admin",
		"password": "admin",
	}).Expect().Status(200).Cookie("go_admin_session").Raw()

	// logout: without login

	printlnWithColor("logout: without login", "green")
	e.GET(config.Get().Url("/logout")).Expect().
		Status(200).
		Body().Contains("login")

	// logout

	printlnWithColor("logout", "green")
	e.GET(config.Get().Url("/logout")).WithCookie("go_admin_session", sesId.Value).Expect().
		Status(200).
		Body().Contains("login")

	// login again

	printlnWithColor("login again", "green")
	return e.POST(config.Get().Url("/signin")).WithForm(map[string]string{
		"username": "admin",
		"password": "admin",
	}).Expect().Status(200).Cookie("go_admin_session").Raw()

}
