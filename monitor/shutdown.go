package monitor

import "github.com/pangxianfei/framework/monitor/app/logics/dashboard"

func Shutdown() error {
	dashboard.Flow.Close()
	return nil
}
