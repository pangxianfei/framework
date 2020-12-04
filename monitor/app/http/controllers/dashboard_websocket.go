package controllers

import (
	"errors"

	"github.com/pangxianfei/framework/http/nouncontroller"
	"github.com/pangxianfei/framework/monitor/app/logics/dashboard"
	"github.com/pangxianfei/framework/request/websocket"
	"github.com/gin-gonic/gin"
)

type DashboardWebsocketController struct {
	nouncontroller.NounController
	websocket.BaseHandler
}

func (d *DashboardWebsocketController) View(this *gin.Context, tplName string, data gin.H) {
	panic("implement me")
}


func (d *DashboardWebsocketController) DefaultChannels() []string {
	return []string{"new-test-channel"}
}

func (d *DashboardWebsocketController) OnMessage(hub websocket.Hub, msg *websocket.Msg) {
	mm1 := &websocket.Msg{}

	if msg.String() == "join test channel" {
		hub.JoinChannel("test")
	}
	if msg.String() == "broadcast to test" {
		mm1.SetString("test broadcast")
		hub.BroadcastTo("test", mm1)
	}

	mm := &websocket.Msg{}
	// need login~, just for an example of websocket authentication support
	if err := hub.ScanUser(); err != nil {
		mm.SetString(err.Error())
		hub.Send(mm)
		return
	}

	mm.SetJSON(hub.User().Value())
	hub.Send(mm)
	return
}

func (d *DashboardWebsocketController) Loop(hub websocket.Hub) error {
	select {
	case flow, ok := <-dashboard.Flow.Current():
		if !ok {
			return errors.New("connection closed")
		}
		msg := websocket.Msg{}
		msg.SetJSON(flow)
		hub.Broadcast(&msg)
	default:
		return nil
	}
	return nil
}
