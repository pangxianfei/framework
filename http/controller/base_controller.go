package controller

import (
	"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/model"
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/request/http/auth"
	"github.com/pangxianfei/framework/validator"
	"net/http"
	"regexp"
	"runtime"
	"strconv"
)



type BaseController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
	model.BaseModel
	ShowData       map[string]interface{}
	Output     map[string]string
	TplName    string
	controller string
	method     string
}

type ControllerInterface interface {
	View()
	SetController()
	SetMethod()
	GetInt64(name string)
}

func (c *BaseController) SetController() {}

func (c *BaseController) SetMethod() {}

func (c *BaseController) View(request request.Context) {

	str := CallerName()
	reg := regexp.MustCompile(`([a-zA-Z0-9_]+)`)
	array := reg.FindAllString(str, -1)
	controllerNum := len(array)
	for i := 0; i < controllerNum; i++ {
		if i == controllerNum-2 {
			c.controller = array[i]
		}
		if i == controllerNum-1 {
			c.method = array[i]
		}
	}

	if len(c.controller) > 0 && len(c.TplName) <= 0 && len(c.method) > 0 {
		c.TplName = tmaic.Strtolower(string("/" + c.controller + "/" + c.method))
	}


	request.HTML(http.StatusOK, c.TplName, c.ShowData)
}



// GetInt64 returns the value associated with the key as an integer.
func (c *BaseController) GetInt64(name string) (i64 int64) {

	if i64, err := strconv.ParseInt(name, 10, 64); err ==nil {
          return i64
	}
   return i64
}



func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func CallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
