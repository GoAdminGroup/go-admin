// +build go1.13

package iris

import (
	"net/http"
	"testing"

	"github.com/GoAdminGroup/go-admin/tests/common"
	// NOTE: iris has its own `kataras/iris/v12/httptest`
	// package which uses gavv/httpexpect under the hoods as well.
	"github.com/gavv/httpexpect"
)

func TestIris(t *testing.T) {
	// TODO: BUG: invalid memory address or nil pointer dereference
	common.Test(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(NewIrisHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
