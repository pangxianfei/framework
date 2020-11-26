package views

import (
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/view"
)

func Initialize(r *request.Engine) {
	view.Initialize(r)
}
