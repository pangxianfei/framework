package controller

import (
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request/http/auth"
	"github.com/pangxianfei/framework/validator"
)

type Controller interface {
	Validate(c validator.Context, _validator interface{}, onlyFirstError bool) (isAbort bool)

	Authorize(c policy.Context, policies policy.Policier, action policy.Action) (permit bool, user auth.IUser)
}
