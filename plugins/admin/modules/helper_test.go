package modules

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInArray(t *testing.T) {
	assert.Equal(t, InArray([]string{"2"}, "2"), true)
}
