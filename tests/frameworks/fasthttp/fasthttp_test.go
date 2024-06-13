package fasthttp

import (
	"net/http"
	"testing"

	"github.com/ChenSee/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestFasthttp(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(internalHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
