package redisclient

import (
	"github.com/chenhg5/go-admin/context"
	c "github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins"
)

type RedisClient struct {
	app *context.App
}

var Plug = new(RedisClient)

var config c.Config

func (redis *RedisClient) InitPlugin() {
	config = c.Get()
	Plug.app = InitRouter(config.Prefix())
}

func NewRedisClient() *RedisClient {
	return Plug
}

func (redis *RedisClient) GetRequest() []context.Path {
	return redis.app.Requests
}

func (redis *RedisClient) GetHandler(url, method string) context.Handlers {
	return plugins.GetHandler(url, method, redis.app)
}
