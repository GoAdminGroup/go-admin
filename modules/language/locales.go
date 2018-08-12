package language

import "goAdmin/config"

var locales = map[string]map[string]string{
	"cn" : cn,
	"en" : en,
}

var Lang = locales[config.EnvConfig["LANGUAGE"].(string)]