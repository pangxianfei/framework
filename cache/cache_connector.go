package cache

import (
	"github.com/pangxianfei/framework/cache/driver/memory"
	"github.com/pangxianfei/framework/cache/driver/redis"

	"github.com/pangxianfei/framework/config"
)

var cer cacher

func Initialize() {
	cer = setStore("default")
}
func setStore(store string) (cer cacher) {

	conn := store
	if store == "default" {
		conn = config.GetString("cache." + store)
		if conn == "" {
			panic("cache connection parse error")
		}
	}

	// get driver instance and connect cache store
	switch conn {
	//memory 驱动
	case "memory":
		cer = memory.NewMemory(
			config.GetString("cache.stores.memory.prefix"),
			config.GetUint("cache.stores.memory.default_expiration_minute"),
			config.GetUint("cache.stores.memory.cleanup_interval_minute"),
		)
	//redis 缓存驱动
	case "redis":
		connection := config.GetString("cache.stores.redis.connection")
		cer = redis.NewRedis(
			config.GetString("database.redis."+connection+".host"),
			config.GetString("database.redis."+connection+".port"),
			config.GetString("database.redis."+connection+".password"),
			config.GetInt("database.redis."+connection+".database"),
			config.GetString("database.redis.options.prefix"),
		)
	default:
		panic("incorrect cache connection provided")
	}

	return cer
}

func Store(store string) cacher {
	return setStore(store)
}

func Cache() cacher {
	return cer
}
