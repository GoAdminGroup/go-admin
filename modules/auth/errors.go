package auth

import "errors"

var (
	ErrGeneralAccountNotFound          = errors.New("general account: not found")
	ErrGeneralAccountIncorrectPassword = errors.New("general account: incorrect password")

	ErrLdapInvalidConfig     = errors.New("ldap: invalid config")
	ErrLdapInvalidUrlSchema  = errors.New("ldap: invalid server url schema")
	ErrLdapNameNotFound      = errors.New("ldap account: name not found")
	ErrLdapIncorrectPassword = errors.New("ldap account: incorrect password")

	ErrUserNotFound        = errors.New("user: not found")
	ErrUserInvalidName     = errors.New("user: invalid name")
	ErrUserInvalidPassword = errors.New("user: invalid password")
)
