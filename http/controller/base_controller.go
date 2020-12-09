package controller

import (
	"github.com/gin-gonic/gin"
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

func init() {}

type BaseController struct {
	policy.Authorization
	auth.RequestUser
	validator.Validation
	model.BaseModel
	ShowData   map[string]interface{}
	Output     map[string]string
	TplName    string
	controller string
	method     string
	Ctx        *gin.Context
}

type ControllerInterface interface {
	View()
	SetController()
	SetMethod()
	GetInt64(c *gin.Context, name string)
	GetUint(c *gin.Context, name string)
}

func (c *BaseController) SetController() {}

func (c *BaseController) SetMethod() {}

func CallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func (c *BaseController) View(request request.Context) {
	ActionController := CallerName()
	reg := regexp.MustCompile(`([a-zA-Z0-9_]+)`)
	CallerName := reg.FindAllString(ActionController, -1)
	controllerNum := len(CallerName)
	for i := 0; i < controllerNum; i++ {
		if i == controllerNum-2 {
			c.controller = CallerName[i]
		}
		if i == controllerNum-1 {
			c.method = CallerName[i]
		}
	}

	if len(c.controller) > 0 && len(c.TplName) <= 0 && len(c.method) > 0 {
		c.TplName = tmaic.Strtolower(string("/" + c.controller + "/" + c.method))
	}
	request.HTML(http.StatusOK, c.TplName, c.ShowData)
}

/*
func (c *BaseController) GetInt64(request request.Context,name string) int64 {
	var uintid int
	requestId :=request.Query(name)
	var err interface{}
	uintid,err = strconv.Atoi(requestId)
	if err != nil{
		panic("string -> int64 is fail！")
	}
	return int64(uintid)
}
*/

func (c *BaseController) GetInt64(key string) (i64 int64) {

	i64, _ = strconv.ParseInt(key, 10, 64)

	return i64
}

// GetInt returns the value associated with the key as an integer.
func (c *BaseController) GetUint(request request.Context, key string) (i uint) {
	if val, ok := request.Get(key); ok && val != nil {
		i, _ = val.(uint)
	}
	return
}

/*
func (c *BaseController) GetUint(request request.Context,name string) uint {
	var uintid int
	requestId :=request.Query(name)

	var err interface{}
	uintid,err = strconv.Atoi(requestId)
	if err != nil{
		panic("string -> uint is fail！")
	}
	return uint(uintid)
}
*/
