package ldapAuth

import (
	"fmt"
	"testing"
)

func TestLdapAuth(t *testing.T) {
	lp := Ldapparam{
		"192.168.1.1",
		"cn=xxx,ou=xxx,dc=test,dc=com,dc=cn",
		"sdfer",
		"ou=xxx,dc=test,dc=com,dc=cn",
		"sAMAccountName",
		true,
	}
	t1, err := LdapAuth("test11", "Abc123456", lp)
	if nil != err {
		fmt.Printf("出错信息：%v, %s\n", t1, err)
	} else {
		fmt.Printf("认证通过否：%v, %s\n", t1, err)
	}
}
