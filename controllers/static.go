package controllers

import (
	"net/http"
)

// StaticHandler 闭包，调用模板的 Execute 方法来将模板呈现给响应
func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
