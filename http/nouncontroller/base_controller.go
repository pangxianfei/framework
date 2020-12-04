package nouncontroller

import (
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request/http/auth"
	"github.com/pangxianfei/framework/validator"
	"github.com/pangxianfei/framework/model"
	//"github.com/gin-gonic/gin"

)

//type View func(code int,tplName string,data gin.H)


type NounController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
	// route controller info
	controllerName string
	//数据模型
	model.BaseModel
	//模板文件名
	TplName        string
	//模板变量
	Data map[string]interface{}
}



