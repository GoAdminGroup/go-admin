package common

import (
	"fmt"
	"github.com/gavv/httpexpect"
	"net/http"
)

func UserTest(e *httpexpect.Expect, sesId *http.Cookie) {

	fmt.Println()
	printlnWithColor("User", "blue")
	fmt.Println("============================")

	printlnWithColor("new", "green")
	e.GET("/ping").Expect().Status(404)
	printlnWithColor("delete", "green")
	e.GET("/pong").Expect().Status(404)
	printlnWithColor("edit", "green")
	e.GET("/pong").Expect().Status(404)
	printlnWithColor("show", "green")
	e.GET("/pong").Expect().Status(404)
}
