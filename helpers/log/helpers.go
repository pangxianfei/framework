package log

import (
	"github.com/pangxianfei/framework/errors"
	"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/logs"
)

func Error(err error, v ...tmaic.V) error {
	var fields tmaic.V
	if len(v) > 0 {
		fields = v[0]
	}
	errors.ErrPrintln(err, fields)
	return err
}

func Info(msg interface{}, v ...tmaic.V) {
	var fields tmaic.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.INFO, msg, fields)
}
func Warn(msg interface{}, v ...tmaic.V) {
	var fields tmaic.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.WARN, msg, fields)
}
func Fatal(msg interface{}, v ...tmaic.V) {
	var fields tmaic.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.FATAL, msg, fields)
}
func Debug(msg interface{}, v ...tmaic.V) {
	var fields tmaic.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.DEBUG, msg, fields)
}
func Panic(msg interface{}, v ...tmaic.V) {
	var fields tmaic.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.PANIC, msg, fields)
}
func Trace(msg interface{}, v ...tmaic.V) {
	var fields tmaic.V
	if len(v) > 0 {
		fields = v[0]
	}
	logs.Println(logs.TRACE, msg, fields)
}
func ErrorStr(err error, v ...tmaic.V) string {
	var fields tmaic.V
	if len(v) > 0 {
		fields = v[0]
	}
	return errors.ErrPrint(err, 2, fields)
}
