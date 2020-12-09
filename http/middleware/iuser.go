package middleware

import (
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/request/http/auth"
)

func IUser(userModelPtr auth.IUser) request.HandlerFunc {

	log.Debug(userModelPtr)
	log.Debug("pangxianfei")
	return func(c request.Context) {
		c.SetIUserModel(userModelPtr)

		c.Next()
	}
}
