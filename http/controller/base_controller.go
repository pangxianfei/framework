package controller

import (

	"github.com/pangxianfei/framework/model"
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/request/http/auth"
	"github.com/pangxianfei/framework/validator"
	"net/http"
)
var Output = make(map[string]string)



type BaseController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
	model.BaseModel
	//模板变量
	Data     map[string]interface{}
	Output   map[string]string
	TplName  string
}

type ControllerInterface interface {
	View()
}

func (c *BaseController) View(request request.Context) {
	request.HTML(http.StatusOK,c.TplName,c.Data)
}
