package view

import (
	"html/template"
	"sync"
	"os"
	"io/ioutil"
	//"github.com/pangxianfei/framework"
	"github.com/pangxianfei/framework/request"
	//"github.com/pangxianfei/framework/helpers/log"
	c "github.com/pangxianfei/framework/config"
)
var routes = route{route: c.GetString("TEMPLATE_TPL"), suffix: c.GetString("SUFFIX")}
func Initialize(r *request.Engine) {
	t := template.New("")
	for _, tmpl := range engineTemplateMap.Get() {
		t, _ = t.New(tmpl.name).Parse(tmpl.content)
	}
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

type tmpl struct {
	name    string
	content string
}

type route struct {
	route    string
	suffix     string
}


type engineTemplate struct {
	lock sync.RWMutex
	data []*tmpl
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

var engineTemplateMap *engineTemplate

func init() {
	engineTemplateMap = newEngineTemplate()
}
