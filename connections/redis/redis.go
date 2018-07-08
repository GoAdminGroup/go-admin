package redis

import (
	"github.com/go-redis/redis"
	"goAdmin/config"
	"time"
)

type ClientType struct {
	RedisCon *redis.Client
}

var RedisClient *ClientType

func InitRedis() {
	RedisClient = &ClientType{
		RedisCon: redis.NewClient(&redis.Options{
			Addr:     config.EnvConfig["REDIS_IP"].(string) + ":" + config.EnvConfig["REDIS_PORT"].(string),
			Password: config.EnvConfig["REDIS_PASSWORD"].(string), // no password set
			DB:       config.EnvConfig["REDIS_DB"].(int),          // use default DB
		}),
	}
}

func (Client *ClientType) Set(key string, value interface{}, expiration time.Duration) *redis.Client {
	err := (*Client).RedisCon.Set(key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	return (*Client).RedisCon
}

func (Client *ClientType) Get(key string) (string, *redis.Client) {
	val, err := (*Client).RedisCon.Get(key).Result()

	if err == redis.Nil {
		return "", (*Client).RedisCon
	}

	if err != nil {
		panic(err)
	}

	return val, (*Client).RedisCon
}

func (Client *ClientType) Del(key string) *redis.Client {
	_, err := (*Client).RedisCon.Del(key).Result()
	if err != nil {
		panic(err)
	}
	return (*Client).RedisCon
}
