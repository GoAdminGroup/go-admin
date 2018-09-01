package config

type Config struct {
	DATABASE_IP           string
	DATABASE_PORT         string
	DATABASE_USER         string
	DATABASE_PWD          string
	DATABASE_NAME         string
	DATABASE_MAX_IDLE_CON string
	DATABASE_MAX_OPEN_CON string
	DATABASE_DRIVER       string

	AUTH_DOMAIN  string
	LANGUAGE     string
	ADMIN_PREFIX string
}

var GlobalCfg Config

func SetGlobalCfg(cfg Config)  {
	GlobalCfg = cfg
}