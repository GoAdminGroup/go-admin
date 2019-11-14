package iris

import (
	"net/http"
	"testing"

	"github.com/GoAdminGroup/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestIris(t *testing.T) {
	// TODO: BUG: invalid memory address or nil pointer dereference
	common.Test(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(newIrisHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
