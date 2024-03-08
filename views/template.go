package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

// Must 确保在没有错误的情况下返回 Template 结构
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

// ParseFS 搜索与模式匹配的文件，然后该函数将这些文件的内容解析为模版
func ParseFS(fs fs.FS, pattern ...string) (Template, error) {
	tpl, err := template.ParseFS(fs, pattern...)
	if err != nil {
		return Template{}, fmt.Errorf("paresing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

// Execute 该方法将内容类型设置为text/html； charset=utf-8，然后使用提供的数据执行模板。
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("执行模版: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
