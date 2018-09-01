package config

type Database struct {
	IP           string
	PORT         string
	USER         string
	PWD          string
	NAME         string
	MAX_IDLE_CON int
	MAX_OPEN_CON int
	DRIVER       string

	FILE string
}

type Config struct {
	DATABASE Database

	AUTH_DOMAIN  string
	LANGUAGE     string
	ADMIN_PREFIX string
}

var GlobalCfg Config

func SetGlobalCfg(cfg Config) {
	GlobalCfg = cfg
}
