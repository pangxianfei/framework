package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/pangxianfei/framework/request"
)

func BasicAuth(accounts map[string]string) request.HandlerFunc {
	return func(c request.Context) {
		gin.BasicAuth(accounts)(c.GinContext())
	}
}
