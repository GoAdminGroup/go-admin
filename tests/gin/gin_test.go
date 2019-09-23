package gin

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func ginHTTPTester(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(NewHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	})
}

func TestGin(t *testing.T) {
	e := ginHTTPTester(t)

	e.GET("/ping").Expect().Status(404)
}
