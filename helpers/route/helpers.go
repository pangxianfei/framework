package route

import (
	"github.com/pangxianfei/framework/route"
)

func Url(routeName string, param ...tmaic.S) (url string, err error) {
	if len(param) > 0 {
		return route.RouteNameMap.Get(routeName, param[0])
	}
	return route.RouteNameMap.Get(routeName, tmaic.S{})
}
