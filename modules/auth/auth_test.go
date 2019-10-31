package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCSRFToken_AddToken(t *testing.T) {
	TokenHelper.AddToken()
	assert.Equal(t, len(*TokenHelper), 1)
	TokenHelper.AddToken()
	assert.Equal(t, len(*TokenHelper), 2)
	tt := (*TokenHelper)[1]
	TokenHelper.CheckToken("123")
	assert.Equal(t, TokenHelper.CheckToken((*TokenHelper)[0]), true)
	assert.Equal(t, tt, (*TokenHelper)[0])
}

func TestEncodePassword(t *testing.T) {
	pwd := EncodePassword([]byte("123456"))
	assert.Equal(t, comparePassword("123456", pwd), true)
}
