package graceful

import (
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework"
)

func panicRecover(quietly bool) {
	if err := recover(); err != nil {
		logFatal(quietly, "tmaic shutting down failed", tmaic.V{"error": err})
	}
}

func logInfo(quietly bool, msg string, v ...tmaic.V) {
	if !quietly {
		log.Info(msg, v...)
	}
}
func logFatal(quietly bool, msg string, v ...tmaic.V) {
	if !quietly {
		log.Fatal(msg, v...)
	}
}
