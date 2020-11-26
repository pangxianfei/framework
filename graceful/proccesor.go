package graceful

import (
	"github.com/pangxianfei/framework/cache"
	//"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/monitor"
	"github.com/pangxianfei/framework/queue"
)

func closeQueue(quietly bool) {
	defer panicRecover(quietly)
	logInfo(quietly, "Queue closing")
	if err := queue.Queue().Close(); err != nil {
		logFatal(quietly, "queue close failed", toto.V{"error": err})
	}
	logInfo(quietly, "Queue closed")
}
func closeDB(quietly bool) {
	defer panicRecover(quietly)
	logInfo(quietly, "Database closing")
   /* pangxianfei by add Close
	if err := m.H().DB().Close(); err != nil {
		logFatal(quietly, "database close failed", toto.V{"error": err})
	}
	logInfo(quietly, "Database closed")
	*/
}
func closeCache(quietly bool) {
	defer panicRecover(quietly)
	logInfo(quietly, "Cache closing")
	if err := cache.Cache().Close(); err != nil {
		logFatal(quietly, "cache close failed", toto.V{"error": err})
	}
	logInfo(quietly, "Cache closed")
}
func closeMonitor(quietly bool) {
	defer panicRecover(quietly)
	logInfo(quietly, "Monitor closing")
	if err := monitor.Shutdown(); err != nil {
		logFatal(quietly, "monitor close failed", toto.V{"error": err})
	}
	logInfo(quietly, "Monitor closed")
}
