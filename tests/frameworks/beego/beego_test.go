package beego

import (
	"net/http"
	"testing"

	"github.com/digroad/go-admin/tests/common"
	"github.com/gavv/httpexpect"
)

func TestNewBeego(t *testing.T) {
	common.ExtraTest(httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(newHandler()),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
	}))
}
