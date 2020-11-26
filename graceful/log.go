package graceful

import (
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/helpers/toto"
)

func panicRecover(quietly bool) {
	if err := recover(); err != nil {
		logFatal(quietly, "tmaic shutting down failed", toto.V{"error": err})
	}
}

func logInfo(quietly bool, msg string, v ...toto.V) {
	if !quietly {
		log.Info(msg, v...)
	}
}
func logFatal(quietly bool, msg string, v ...toto.V) {
	if !quietly {
		log.Fatal(msg, v...)
	}
}
