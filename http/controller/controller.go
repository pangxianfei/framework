package controller

import (
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request/http/auth"
	"github.com/pangxianfei/framework/validator"
	//"github.com/gin-gonic/gin"

)

type Controller interface {
	Validate(c validator.Context, _validator interface{}, onlyFirstError bool) (isAbort bool)
	Authorize(c policy.Context, policies policy.Policier, action policy.Action) (permit bool, user auth.IUser)
	//view 未实现方法
	//View (this *gin.Context,tplName string,data gin.H)
}

