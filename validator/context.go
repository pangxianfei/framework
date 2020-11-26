package validator

import (
	"github.com/pangxianfei/framework/context"
	"github.com/pangxianfei/framework/resources/lang"
)

type Context interface {
	context.RequestBindingContextor
	context.ResponseContextor
	lang.Context
}
