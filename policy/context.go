package policy

import (
	"github.com/pangxianfei/framework/context"
	"github.com/pangxianfei/framework/request/http/auth"
)

type Context interface {
	context.LifeCycleContextor
	context.ResponseContextor
	auth.Context
	auth.RequestIUser
}
