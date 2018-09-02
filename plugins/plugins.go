package plugins

import (
	"github.com/chenhg5/go-admin/context"
	"github.com/chenhg5/go-admin/modules/config"
	"strings"
	"regexp"
)

type Plugin interface {
	GetRequest() []context.Path
	GetHandler(url, method string) context.Handler
	InitPlugin(config config.Config)
}

func GetHandler(url, method string, handleList *map[context.Path]context.Handler) context.Handler {
	for path, handler := range *handleList {
		if path.Method == method {
			if path.RegUrl == "" {
				if path.URL == url || path.URL + "/" == url || path.URL == url + "/" {
					return handler
				}
			} else {
				if strings.Count(path.RegUrl, "/") == strings.Count(url, "/") {
					if ok, err := regexp.MatchString(path.RegUrl, url); ok && err == nil {
						return handler
					}
				}
			}
		}
	}

	panic("handler not found")
}