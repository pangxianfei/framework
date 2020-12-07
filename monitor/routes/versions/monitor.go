package versions

import (
	//"github.com/pangxianfei/framework/config"
	//"github.com/pangxianfei/framework/http/middleware"
	//"github.com/pangxianfei/framework/monitor/routes/groups"

	"github.com/pangxianfei/framework/request"
	//"github.com/pangxianfei/framework/route"
)

func NewMonitor(engine *request.Engine) {
	//ver := route.NewVersion(engine, "monitor")

	//accounts := make(map[string]string)
	//accounts[config.GetString("monitor.username")] = config.GetString("monitor.password")

	// noauth routes
	//ver.NoAuth("", func(grp route.Grouper) {
		//grp.AddGroup("/dashboard", &groups.DashboardGroup{})
	//}, middleware.BasicAuth(accounts))
}
