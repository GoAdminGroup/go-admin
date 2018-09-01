package config

type Config struct {
	Username     string
	Password     string
	Port         string
	Ip           string
	DatabaseName string
	IdleCon      int
	OpenCon      int
}
