package logs

import (
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/sentry"
)

var log *logrus.Logger
var logLevel Level

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.Out = os.Stdout
}

func Initialize() {
	levelStr := config.GetString("app.log_level")
	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		panic(err)
	}

	logLevel = level
	log.SetLevel(logLevel)
}

type Field = toto.V

func Println(level Level, msg interface{}, fields Field) {

	if fields == nil {
		log.Log(level, msg)
	} else {
		log.WithFields(logrus.Fields(fields)).Log(level, msg)
	}

	if level <= logLevel {
		_fields := make(map[string]interface{})
		if fields != nil {
			_fields = fields
		}

		switch level {
		case PANIC:
			sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
		case FATAL:
			sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
		case ERROR:
			sentry.CaptureError(errors.New(fmt.Sprintf("%v - %v", msg, fields)))
		case WARN:
			_fields["level"] = "WARN"
			sentry.CaptureMsg(fmt.Sprintf("%v", msg), _fields)
			//case INFO:
			//	_fields["level"] = "INFO"
			//	sentry.CaptureMsg(msg, _fields)
			//case DEBUG:
			//	_fields["level"] = "DEBU"
			//	sentry.CaptureMsg(msg, _fields)
			//case TRACE:
			//	_fields["level"] = "TRAC"
			//	sentry.CaptureMsg(msg, _fields)
		}
	}
}
