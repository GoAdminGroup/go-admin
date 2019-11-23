package modules

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInArray(t *testing.T) {
	assert.Equal(t, InArray([]string{"2"}, "2"), true)
}

func TestIsInfoUrl(t *testing.T) {
	assert.Equal(t, IsInfoUrl("/admin/info/user?id=asfas"), true)
}
