package route

import (
	"errors"
	"fmt"
	"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/app"
	"github.com/pangxianfei/framework/helpers/log"
	"github.com/pangxianfei/framework/http/middleware"
	"github.com/pangxianfei/framework/model"
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"
	"github.com/pangxianfei/framework/request/websocket"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

const maxRouteMapLength = 1000

var RouteNameMap *routeNameMap

type route struct {
	name              string
	bindFunc          func(handlers ...request.HandlerFunc)
	handlers          []request.HandlerFunc
	relativePath      string
	httpMethod        string
	basicPath         string
	prefixHandlersNum int
	wsHandler         websocket.Handler
	controller        string
	function          string
}

type routeNameMap struct {
	lock       sync.RWMutex
	data       map[string]string
	controller string
	function   string
}

func newRouteNameMap() *routeNameMap {
	return &routeNameMap{
		data: make(map[string]string),
	}
}

func (rnm *routeNameMap) set(routeName, routeUrl string) {
	rnm.lock.Lock()
	defer rnm.lock.Unlock()
	rnm.data[routeName] = routeUrl
}
func (rnm *routeNameMap) Get(routeName string, param tmaic.S) (url string, err error) {
	rnm.lock.RLock()
	defer rnm.lock.RUnlock()
	routeUrl, ok := rnm.data[routeName]
	if !strings.Contains(routeUrl, "/:") {
		// simple url
		return routeUrl, nil
	}

	// param url
	if !ok {
		return "", errors.New("Cannot find route by Name " + routeName)
	}
	for k, v := range param {
		routeUrl = strings.Replace(routeUrl, "/:"+k, v, 1)
	}
	return routeUrl, nil
}

type engineRoute struct {
	lock       sync.RWMutex
	data       map[request.EngineHash]chan *route
	name       string
	controller string
	Function   string
}

func newEngineRoute() *engineRoute {
	return &engineRoute{
		data: make(map[request.EngineHash]chan *route),
	}
}
func (erm *engineRoute) initEngine(eh request.EngineHash) {
	if erm.data[eh] == nil {
		erm.data[eh] = make(chan *route, maxRouteMapLength)
	}
}
func (erm *engineRoute) Get(eh request.EngineHash) chan *route {
	erm.lock.RLock()
	defer erm.lock.RUnlock()
	return erm.data[eh]
}
func (erm *engineRoute) Set(eh request.EngineHash, r *route) {
	//@todo doesn't process the situation that for multi serve
	erm.lock.Lock()
	defer erm.lock.Unlock()
	erm.initEngine(eh)
	erm.data[eh] <- r
}
func (erm *engineRoute) Close(eh request.EngineHash) {
	close(erm.data[eh])
}

var engineRouteMap *engineRoute

func init() {
	engineRouteMap = newEngineRoute()
	RouteNameMap = newRouteNameMap()
}

func newRoute(httpMethod string, g *group, relativePath string, bindFunc func(ginHandlers ...request.HandlerFunc), handlers ...request.HandlerFunc) *route {
	r := route{httpMethod: httpMethod, prefixHandlersNum: len(g.RouterGroup.Handlers), basicPath: g.RouterGroup.BasePath(), relativePath: relativePath, bindFunc: bindFunc, handlers: handlers}

	engineRouteMap.Set(g.engineHash, &r)

	return &r
}
func newWsRoute(httpMethod string, g *group, relativePath string, bindWsFunc func(wsHandler websocket.Handler, ginHandlers ...request.HandlerFunc), wsHandler websocket.Handler, handlers ...request.HandlerFunc) *route {
	r := route{wsHandler: wsHandler, httpMethod: httpMethod, prefixHandlersNum: len(g.RouterGroup.Handlers), basicPath: g.RouterGroup.BasePath(), relativePath: relativePath, bindFunc: func(handlers ...request.HandlerFunc) {
		// normalHandler -> wshandler
		httpHs := append(handlers, func(context request.Context) {
			//@todo websocket handler placeholder
		})
		bindWsFunc(wsHandler, httpHs...)
	}, handlers: handlers}

	engineRouteMap.Set(g.engineHash, &r)

	return &r
}
func (r *route) Name(routeName string) policy.RoutePolicier {
	r.name = routeName
	return r
}
func (r *route) Can(policies policy.Policier, action policy.Action) {
	r.handlers = append([]request.HandlerFunc{middleware.Policy(policies, action)}, r.handlers...)
}
func (r *route) handlerNum() int {
	return r.prefixHandlersNum + len(r.handlers)
}
func (r *route) absolutePath() string {
	return fmt.Sprintf("%s%s", r.basicPath, r.relativePath)
}
func (r *route) lastHandlerName() string {

	if r.httpMethod == httpMethodWebsocket {
		h := reflect.ValueOf(r.wsHandler).Elem().Type()
		return h.PkgPath() + "." + h.Name()
	}

	if len(r.handlers) > 0 {
		str := runtime.FuncForPC(reflect.ValueOf(r.handlers[len(r.handlers)-1]).Pointer()).Name()
		str = strings.TrimSuffix(str, "-fm")
		str = strings.TrimPrefix(str, "tmaic/app/http/controllers.")
		reg := regexp.MustCompile(`([a-zA-Z0-9_]+)`)
		array := reg.FindAllString(str, -1)
		controllerNum := len(array)
		for i := 0; i < controllerNum; i++ {
			if i == 0 {
				r.controller = array[i]
			}
			if i == 1 {
				r.function = array[i]
			}
		}
		return runtime.FuncForPC(reflect.ValueOf(r.handlers[len(r.handlers)-1]).Pointer()).Name()
	}
	return "nil"
}

func Bind(engine *request.Engine) {
	hash := engine.Hash()
	defer engineRouteMap.Close(hash)

	var routeNum int
	routeNum = 0
	for r := range engineRouteMap.Get(hash) {
		r.bindFunc(r.handlers...)

		// for Route() function to easy retrieve url by name
		RouteNameMap.set(r.name, r.absolutePath())

		/*************************保存路由 start *********************************/
		r.lastHandlerName()
		webroute := model.Webroutes{
			RouteName:  r.name,
			Controller: r.controller,
			Function:   r.function,
			Path:       r.absolutePath(),
		}
		db := webroute.DB()
		if routeNum == 0 {
			//清空表记录  记录新的路由 pangxianfei by add
			db.Exec("truncate table " + webroute.TableName())
		}
		db.Create(&webroute)
		/**************************保存路由 end **********************************/

		if app.GetMode() != app.ModeProduction {
			log.Info(fmt.Sprintf("%-6s %-30s --> controller:%s       -> function:%s", r.httpMethod, r.absolutePath(), r.controller, r.function))
		}

		if len(engineRouteMap.Get(hash)) <= 0 {
			break
		}
		routeNum++
	}
}
