package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pangxianfei/framework/request/http"
)

func ConvertHandlers(handlers []HandlerFunc) (ginHandlers []gin.HandlerFunc) {
	for _, h := range handlers {
		handler := h // must new a variable for `range's val`, or the `val` in anonymous funcs will be overwrited in every loop

		ginHandlers = append(ginHandlers, func(c *gin.Context) {
			tmaicContext := http.ConvertContext(c)
			handler(tmaicContext)
		})
	}
	return
}
