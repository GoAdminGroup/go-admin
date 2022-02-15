package config

type Ldap struct {
	Enable     bool     `json:"enable,omitempty" yaml:"enable,omitempty" ini:"enable,omitempty"`
	ServerUrls []string `json:"server_urls,omitempty" yaml:"server_urls,omitempty" ini:"server_urls,omitempty"`
	BindDN     string   `json:"bind_dn,omitempty" yaml:"bind_dn,omitempty" ini:"bind_dn,omitempty"`
	BindPwd    string   `json:"bind_pwd,omitempty" yaml:"bind_pwd,omitempty" ini:"bind_pwd,omitempty"`
	BaseDN     string   `json:"base_dn,omitempty" yaml:"base_dn,omitempty" ini:"base_dn,omitempty"`
}
