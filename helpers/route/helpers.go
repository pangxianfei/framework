package route

import (
	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/route"
)

func Url(routeName string, param ...toto.S) (url string, err error) {
	if len(param) > 0 {
		return route.RouteNameMap.Get(routeName, param[0])
	}
	return route.RouteNameMap.Get(routeName, toto.S{})
}
