package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/pangxianfei/framework/request"
)

func Logger() request.HandlerFunc {
	return func(c request.Context) {
		gin.Logger()(c.GinContext())
	}
}
