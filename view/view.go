package view

import (
	"errors"
	//"gitee.com/zhucheer/orange/cfg"
	c "github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/utils"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

type AppTemplate struct {
	ViewName    string
	ViewPath    string
	IncludeTmpl []string
	ShowData    map[string]interface{}
	mutex       sync.Mutex
}

func ContextTmpl(viewName string, includeFiles ...string) *AppTemplate {
	viewPath := parseViewName(viewName)
	includeTmpl := make([]string, 0)
	for _, item := range includeFiles {
		includeTmpl = append(includeTmpl, parseViewName(item))
	}
	appTmpl := &AppTemplate{
		ViewName:    viewName,
		ViewPath:    viewPath,
		ShowData:    make(map[string]interface{}),
		IncludeTmpl: includeTmpl,
	}

	return appTmpl
}

func (t *AppTemplate) Assigns(values interface{}) *AppTemplate {
	objT := reflect.TypeOf(values)
	objV := reflect.ValueOf(values)

	if objT.Kind() == reflect.Ptr {
		objT = objT.Elem()
		objV = objV.Elem()
	}
	switch objT.Kind() {
	case reflect.Struct:
		t.mutex.Lock()
		defer t.mutex.Unlock()
		for i := 0; i < objT.NumField(); i++ {
			objName := objT.Field(i).Name
			objValue := objV.Field(i).Interface()
			t.ShowData[objName] = objValue
		}
	case reflect.Map:
		t.mutex.Lock()
		defer t.mutex.Unlock()
		item := objV.MapRange()
		for item.Next() {
			k := item.Key()
			v := item.Value()
			if k.Kind() != reflect.String {
				typePanic()
			}
			t.ShowData[k.String()] = v.Interface()
		}
	default:
		typePanic()
	}

	return t
}

func (t *AppTemplate) Assign(name string, value interface{}) *AppTemplate {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.ShowData[name] = value
	return t
}

func (t *AppTemplate) Render() (viewHtmlRes string, err error) {
	isFile, _ := utils.FileExists(t.ViewPath)
	if isFile == false {
		errMsg := t.ViewPath + " not found"
		return viewHtmlRes, errors.New(errMsg)
	}

	viewHtmlRes = htmlPath(t.ViewPath, t.ShowData, t.IncludeTmpl...)
	return viewHtmlRes, nil
}

func (t *AppTemplate) RenderText() (viewHtmlRes string, err error) {
	isFile, _ := utils.FileExists(t.ViewPath)
	if isFile == false {
		errMsg := t.ViewPath + " not found"
		return viewHtmlRes, errors.New(errMsg)
	}

	viewHtmlRes = textPath(t.ViewPath, t.ShowData, t.IncludeTmpl...)
	return viewHtmlRes, nil
}

func parseViewName(viewName string) string {
	viewPathPre := c.GetString("app.TEMPLATE_TPL", c.GetString("app.TEMPLATE_TPL"))
	if viewPathPre == "" {
		panic("app.viewPath config not found")
	}

	workDir, _ := os.Getwd()
	viewPathPre = filepath.Join(workDir, viewPathPre)
	viewName = strings.Replace(viewName, ".", string(filepath.Separator), -1)
	viewPath := viewPathPre + string(filepath.Separator) + viewName + ".tpl"

	return viewPath
}

func typePanic() {
	panic("viewData must be map[string]interface{} or struct ")
}
