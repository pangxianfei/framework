package view

import (
	"html/template"
	"io/ioutil"
	"os"
	"sync"
	c "github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/request"
	//"github.com/pangxianfei/framework/helpers/log" //log.Debug(t.Name())

)
var routes = route{route: c.GetString("TEMPLATE_TPL"), suffix: c.GetString("SUFFIX")}
var engineTemplateMap *engineTemplate



type view interface {
	Initialize(r *request.Engine)
	Show(c request.Context,data map[string]interface{})
}




type tmpl struct {
	name    string
	content string
}

type Template struct {
	//模板文件名
	TplName        string
	//模板变量
	Data map[string]interface{}
}

type route struct {
	route    string
	suffix   string
}


type engineTemplate struct {
	lock sync.RWMutex
	data []*tmpl
}


func init() {
	engineTemplateMap = newEngineTemplate()
	//log.Debug("init")
}

func Initialize(r *request.Engine) {
	t := template.New("")
	for _, tmpl := range engineTemplateMap.Get() {
		t, _ = t.New(tmpl.name).Parse(tmpl.content)
	}

	//log.Debug(t.Name())
	r.SetHTMLTemplate(t)
}

func AddView(name string) {

	FILE_PATH := routes.route+name+routes.suffix
	_, err := os.Stat(FILE_PATH)
	if err == nil {
		content, err := ioutil.ReadFile(routes.route+name+routes.suffix)
		if err != nil {
			panic("Template does not exist!")
		}

		engineTemplateMap.Set(&tmpl{
			name:    name,
			content: string(content),
		})

	}
	//else {
	//	panic("System Template does not exist!")
	//}

}

func newEngineTemplate() *engineTemplate {
	return &engineTemplate{
		data: []*tmpl{},
	}
}
func (et *engineTemplate) Get() []*tmpl {
	et.lock.RLock()
	defer et.lock.RUnlock()
	return et.data
}
func (et *engineTemplate) Set(tmpl *tmpl) {
	et.lock.Lock()
	defer et.lock.Unlock()
	et.data = append(et.data, tmpl)
}
//自定义方法  与类方无关
func (T *Template) Show(c request.Context) {
	 c.View(T.TplName,T.Data)
}
