package controller

import (
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request/http/auth"
	"github.com/pangxianfei/framework/validator"
)

type BaseController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
}
