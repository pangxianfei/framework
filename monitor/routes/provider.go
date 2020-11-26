package routes

import (
	"github.com/pangxianfei/framework/monitor/routes/versions"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/route"
)

func Register(router *request.Engine) {
	defer route.Bind(router)

	versions.NewMonitor(router)
}
