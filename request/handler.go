package request

import (
	"github.com/gin-gonic/gin"
	"github.com/pangxianfei/framework/request/http"
	//"github.com/pangxianfei/framework/helpers/log"
)

func ConvertHandlers(handlers []HandlerFunc) (ginHandlers []gin.HandlerFunc) {
	for _, h := range handlers {
		// 必须为“range val”新建一个变量，否则匿名函数中的“val”将在每个循环中被覆盖
		handler := h
		ginHandlers = append(ginHandlers, func(c *gin.Context) {
			tmaicContext := http.ConvertContext(c)
			//log.Debug(tmaicContext.Keys())
			handler(tmaicContext)
		})
	}
	return
}
