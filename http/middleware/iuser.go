package middleware

import (
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/request/http/auth"
)

func IUser(userModelPtr auth.IUser) request.HandlerFunc {
	return func(c request.Context) {
		c.SetIUserModel(userModelPtr)

		c.Next()
	}
}
