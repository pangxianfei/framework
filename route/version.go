package route

import (
	"github.com/gin-gonic/gin"

	"github.com/pangxianfei/framework/http/middleware"
	"github.com/pangxianfei/framework/request"
)

type version struct {
	engine *request.Engine
	group  *gin.RouterGroup
	prefix string
}

func NewVersion(engine *request.Engine, prefix string) *version {
	ver := &version{engine: engine, prefix: prefix}
	ver.group = ver.engine.Group(prefix)
	return ver
}

func (v *version) Auth(signKey string, relativePath string, groupFunc func(grp Grouper), handlers ...request.HandlerFunc) {
	ginGroup := v.group.Group(relativePath, request.ConvertHandlers(append([]request.HandlerFunc{middleware.AuthRequired(signKey)}, handlers...))...)
	groupFunc(&group{engineHash: v.engine.Hash(), RouterGroup: ginGroup})
}

func (v *version) NoAuth(relativePath string, groupFunc func(grp Grouper), handlers ...request.HandlerFunc) {
	ginGroup := v.group.Group(relativePath, request.ConvertHandlers(handlers)...)
	groupFunc(&group{engineHash: v.engine.Hash(), RouterGroup: ginGroup})
}
