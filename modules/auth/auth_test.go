package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodePassword(t *testing.T) {
	pwd := EncodePassword([]byte("123456"))
	assert.Equal(t, comparePassword("123456", pwd), true)
}
