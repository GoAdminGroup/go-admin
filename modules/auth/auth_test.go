package auth

import (
	"fmt"
	"testing"
)

func TestCSRFToken_AddToken(t *testing.T) {
	TokenHelper.AddToken()
	fmt.Println("TokenHelper", TokenHelper)
	TokenHelper.AddToken()
	fmt.Println("TokenHelper 2", TokenHelper)
	fmt.Println("CheckToken", TokenHelper.CheckToken("123"))
	fmt.Println("TokenHelper 3", TokenHelper)
	fmt.Println("CheckToken", TokenHelper.CheckToken((*TokenHelper)[0]))
	fmt.Println("TokenHelper 4", TokenHelper)
}
