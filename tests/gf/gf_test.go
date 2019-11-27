package gf

import (
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/test/gtest"
	"testing"
	"time"
)

func TestNewHandler(t *testing.T) {

	s := newHandler()
	s.SetPort(8103)
	s.SetDumpRouteMap(false)
	_ = s.Start()
	defer func() {
		_ = s.Shutdown()
	}()

	time.Sleep(100 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix("http://127.0.0.1:8103")

		resp, _ := client.Post("/admin/signin", map[string]string{
			"username": "admin",
			"password": "",
		})

		gtest.Assert(resp.StatusCode, 400)

		resp, _ = client.Post("/admin/signin", map[string]string{
			"username": "admin",
			"password": "admin",
		})

		gtest.Assert(resp.StatusCode, 200)

		cookie := resp.GetCookie(auth.DefaultCookieKey)
		gtest.Assert(cookie != "", true)
	})
}
