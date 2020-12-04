package controller

import (
	//"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/model"
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/request/http/auth"
	"github.com/pangxianfei/framework/validator"
	//tmaic "github.com/pangxianfei/framework"
)




type BaseController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
	model.BaseModel
	Lookup   map[string]string

	//模板变量
	Data     map[string]interface{} //Data map[string]interface{}
	TestData map[interface{}]interface{}

	// template data
	TplName        string
	ViewPath       string
	Layout         string
	LayoutSections map[string]string // the key is the section name and the value is the template name
	TplPrefix      string
	TplExt         string
	EnableRender   bool
}

// ControllerInterface is an interface to uniform all controller handler.
type ControllerInterface interface {
	View()
}

func (c *BaseController) View(request request.Context) {
	//log.Debug(c.Lookup)
	request.View(c.TplName, c.Data)
}
