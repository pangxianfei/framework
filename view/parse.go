package view

import (
	"bytes"
	"html/template"
	htmlTemplate "html/template"
	"net/url"
	"path/filepath"
	textTemplate "text/template"
)

var htmlTemplateFuncs = template.FuncMap{
	"unescaped": unescaped,
	"urlencode": urlencode,
	"urldecode": urldecode,
}

func TextPath(tplPath string, viewData interface{}, includeFiles ...string) (res string) {
	return textPath(tplPath, viewData, includeFiles...)
}

// 渲染文本类型模版
func textPath(tplPath string, viewData interface{}, includeFiles ...string) (res string) {
	tplName := filepath.Base(tplPath)
	showBuffer := bytes.NewBuffer([]byte{})

	tmpl := textTemplate.Must(textTemplate.New(tplName).ParseFiles(tplPath))
	tmpl.ParseFiles(includeFiles...)

	tmpl.Execute(showBuffer, viewData)
	res = showBuffer.String()
	return
}

// 渲染html类型模版
func htmlPath(tplPath string, viewData interface{}, includeFiles ...string) (res string) {
	tplName := filepath.Base(tplPath)
	showBuffer := bytes.NewBuffer([]byte{})
	tmpl := htmlTemplate.Must(htmlTemplate.New(tplName).Funcs(htmlTemplateFuncs).ParseFiles(tplPath))
	tmpl.ParseFiles(includeFiles...)

	tmpl.Execute(showBuffer, viewData)
	res = showBuffer.String()
	return
}

// 模版函数
func unescaped(x string) interface{} {
	return template.HTML(x)
}

func urlencode(x string) interface{} {
	return template.URLQueryEscaper(x)
}

func urldecode(x string) interface{} {
	decode, _ := url.PathUnescape(x)
	return decode
}
