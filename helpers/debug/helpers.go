package debug

import (
	"errors"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/ztrue/tracerr"

	"github.com/pangxianfei/framework/console"
	"github.com/pangxianfei/framework/helpers/log"
)

func Dump(v ...interface{}) {
	console.Println(console.CODE_ERROR, spew.Sdump(v...))
	debugPrint(errors.New("====== tmaic Debug ======"))
}

func DD(v ...interface{}) {
	Dump(v...)
	os.Exit(1)
}

func debugPrint(err error) {
	startFrom := 2
	traceErr := tracerr.Wrap(err)
	frameList := tracerr.StackTrace(traceErr)
	if startFrom > len(frameList) || len(frameList)-2 <= 0 {
		_ = log.Error(err)
	}
	traceErr = tracerr.CustomError(err, frameList[startFrom:len(frameList)-2])
	tracerr.PrintSourceColor(traceErr, 0)
}
