package middleware

import (
	l "github.com/pangxianfei/framework/helpers/locale"
	"github.com/pangxianfei/framework/request"

	"github.com/pangxianfei/framework/config"
)

func Locale() request.HandlerFunc {
	return func(c request.Context) {
		locale := c.Request().Header.Get("locale")
		if locale == "" {
			locale = c.DefaultQuery("locale", config.GetString("app.locale"))
		}

		l.SetLocale(c, locale)

		c.Next()
	}
}
