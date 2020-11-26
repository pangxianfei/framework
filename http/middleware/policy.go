package middleware

import (
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"
)

func Policy(_policy policy.Policier, action policy.Action) request.HandlerFunc {
	return func(c request.Context) {
		policy.Middleware(_policy, action, c, c.Params())
	}
}
