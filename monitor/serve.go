package monitor

import (
	"context"
	"net/http"
	"sync"

	//"github.com/pangxianfei/framework/monitor/resources/views"

	c "github.com/pangxianfei/framework/config"
	//"github.com/pangxianfei/framework/helpers/log"
	//"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/helpers/zone"
	//"github.com/pangxianfei/framework/monitor/routes"
	"github.com/pangxianfei/framework/request"
)

func HttpMonitorServe(parentCtx context.Context, wg *sync.WaitGroup) {
	r := request.New()



	s := &http.Server{
		Addr:           ":" + c.GetString("monitor.port"),
		Handler:        r,
		ReadTimeout:    zone.Duration(c.GetInt64("app.read_timeout_seconds")) * zone.Second,
		WriteTimeout:   zone.Duration(c.GetInt64("app.write_timeout_seconds")) * zone.Second,
		MaxHeaderBytes: 1 << 20,
	}


	<-parentCtx.Done()


	ctx, cancel := context.WithTimeout(parentCtx, 5*zone.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		//log.Fatal("Monitor Server Shutdown: ", tmaic.V{"error": err})
	}

	wg.Done()
}
