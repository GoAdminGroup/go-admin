package iris

import (
	"net/http"
	"testing"

	"github.com/GoAdminGroup/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestIris(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(newHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
