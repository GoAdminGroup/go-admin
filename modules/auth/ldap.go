package auth

import (
	"crypto/tls"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/go-ldap/ldap/v3"
	"net/http"
	"net/url"
)

type LdapConfig struct {
	urls    []string
	bindDN  string
	BindPwd string
	baseDN  string
}

func NewLdapConfig(urls []string, bindDN, bindPwd, baseDN string) *LdapConfig {
	return &LdapConfig{urls: urls, bindDN: bindDN, BindPwd: bindPwd, baseDN: baseDN}
}

type ldapAuth struct {
	ldapConf *LdapConfig
	dbConn   db.Connection
}

func NewLdapAuth(dbConn db.Connection, conf *LdapConfig) Authenticator {
	return &ldapAuth{dbConn: dbConn, ldapConf: conf}
}

func (auth *ldapAuth) Authenticate(req *http.Request) (user models.UserModel, err error) {
	var (
		username = req.FormValue("username")
		password = req.FormValue("password")
		ldapConn *ldap.Conn
		result   *ldap.SearchResult
	)
	defer user.ReleaseConn()

	if ldapConn, err = ConnectLdap(auth.ldapConf.urls); err != nil {
		return
	}
	defer ldapConn.Close()

	if _, err = ldapConn.SimpleBind(ldap.NewSimpleBindRequest(auth.ldapConf.bindDN, auth.ldapConf.BindPwd, nil)); err != nil {
		return
	}
	searchReq := ldap.NewSearchRequest(auth.ldapConf.baseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(sAMAccountName=%s))", username),
		[]string{"dn", "mail"}, nil)

	if result, err = ldapConn.Search(searchReq); err != nil {
		return
	}
	if len(result.Entries) == 0 {
		err = ErrLdapNameNotFound
		return
	}
	for _, entry := range result.Entries {
		if err = ldapConn.Bind(entry.DN, password); err == nil {
			ldapAccount := models.NewLdapAccount().WithConn(auth.dbConn).FindByUsernameAndDN(username, entry.DN)
			if ldapAccount.IsEmpty() {
				_, user, err = models.NewLdapAccount().WithConn(auth.dbConn).CreateUser(username, entry.DN)
				return
			}
			if user = models.User().SetConn(auth.dbConn).Find(ldapAccount.UserId); user.IsEmpty() {
				err = ErrUserNotFound
			}
			return
		}
	}
	err = ErrLdapIncorrectPassword
	return
}

func ConnectLdap(urls []string) (*ldap.Conn, error) {
	for _, addr := range urls {
		parsedURL, err := url.Parse(addr)
		if err != nil {
			return nil, err
		}
		switch parsedURL.Scheme {
		case "ldap":
			return ldap.Dial("tcp", parsedURL.Host)
		case "ldaps":
			tlsConf := tls.Config{
				ServerName: parsedURL.Hostname(),
			}
			return ldap.DialTLS("tcp", parsedURL.Host, &tlsConf)
		default:
			return nil, ErrLdapInvalidUrlSchema
		}
	}
	return nil, ErrLdapInvalidConfig
}
