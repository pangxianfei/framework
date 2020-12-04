package controllers

import (
	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
)

type Dashboard struct {
	controller.BaseController
}

func (d *Dashboard) Index(c request.Context) {
	c.View("tmaic_dashboard.index", tmaic.V{
		"url": "ws://" + ":" + config.GetString("monitor.port") + "/monitor/dashboard/ws",
	})
	return
}
