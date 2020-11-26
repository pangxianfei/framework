package cache

import (
	"github.com/pangxianfei/framework/cache/driver"
)

type cacher interface {
	driver.ProtoCacher
	driver.BasicCacher
}
