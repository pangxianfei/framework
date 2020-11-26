package controllers

import (
	"net/http"

	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/request"
)

type Dashboard struct {
	controller.BaseController
}

func (d *Dashboard) Index(c request.Context) {
	c.HTML(http.StatusOK, "totoval_dashboard.index", toto.V{
		"url": "ws://" + ":" + config.GetString("monitor.port") + "/monitor/dashboard/ws",
	})
	return
}
