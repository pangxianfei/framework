package redis

import (
	r "github.com/go-redis/redis"

	. "github.com/pangxianfei/framework/cache/driver/redis"
	"github.com/pangxianfei/framework/config"
)

type redis = r.Client

func Connection(name string) *redis {
	return NewRedis(
		config.GetString("database.redis."+name+".host"),
		config.GetString("database.redis."+name+".port"),
		config.GetString("database.redis."+name+".password"),
		config.GetInt("database.redis."+name+".database"),
		config.GetString("database.redis.options.prefix"),
	).Cache()
}
