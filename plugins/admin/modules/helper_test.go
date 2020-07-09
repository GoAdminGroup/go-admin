package modules

import (
	"regexp"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestInArray(t *testing.T) {
	assert.Equal(t, isFormURL("/admin/info/profile/new"), true)
}

func isFormURL(s string) bool {
	reg, _ := regexp.Compile("(.*?)info/(.*)/(new|edit)(.*?)")
	return reg.MatchString(s)
}
