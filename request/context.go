package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pangxianfei/framework/context"
	"github.com/pangxianfei/framework/request/http/auth"
	"github.com/pangxianfei/framework/utils/jwt"
)

type Context interface {
	context.HttpContextor
	GinContext() *gin.Context

	SetAuthClaim(claims *jwt.UserClaims) //@todo move into a new interface
	SetIUserModel(iUser auth.IUser)      //@todo move into a new interface
	auth.Context
	auth.RequestIUser
}
