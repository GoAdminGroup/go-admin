package ldapAuth

import (
	"crypto/tls"
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

type Ldapparam struct {
	LDAPURL            string //LDAP Server IP or dns
	LDAPSearchDN       string //有读取LDAP树权限的用户，格式：cn=xxx,ou=xxx,dc=xxx,dc=com
	LDAPSearchPassword string
	LDAPBaseDN         string //从哪个分支开始读
	LDAPUid            string //sAMAccountName
	TLS                bool   //是否启用加密
}

func ldapConn(lp Ldapparam) (l *ldap.Conn, err error) {
	ldapUrl := "ldap://" + lp.LDAPURL + ":389"

	l, err = ldap.DialURL(ldapUrl)
	if nil != err {
		return
	}
	if lp.TLS {
		err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if nil != err {
			return
		}
	}
	//bind
	err = l.Bind(lp.LDAPSearchDN, lp.LDAPSearchPassword)
	if nil != err {
		return
	}
	return
}

//search usr
func searchUsr(lp Ldapparam, userName string, l *ldap.Conn) (find bool, err error) {
	find = false
	searchRequest := ldap.NewSearchRequest(
		lp.LDAPBaseDN, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(%s=%s))", lp.LDAPUid, userName),
		//"(&(objectClass=organizationalPerson))", // The filter to apply
		[]string{"dn"}, // A list attributes to retrieve
		nil,
	)
	_, err = l.Search(searchRequest)
	if nil != err {
		return
	}
	find = true
	return
}

func LdapAuth(userName string, pw string, lp Ldapparam) (bool, error) {
	l, err := ldapConn(lp)
	defer l.Close()
	if nil != err {
		return false, err
	}
	find := false
	find, err = searchUsr(lp, userName, l)
	if nil != err {
		return false, err
	}
	err = l.Bind(userName, pw)
	if nil != err {
		return false, err
	}
	return find, err
}
