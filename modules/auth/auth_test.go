package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodePassword(t *testing.T) {
	pwd := EncodePassword([]byte("123456"))
	assert.Equal(t, comparePassword("123456", pwd), true)
}
