package main

import (
	"fmt"
	"github.com/baimiyishu13/vanguard/controllers"
	"github.com/baimiyishu13/vanguard/templates"
	"github.com/baimiyishu13/vanguard/views"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	// test
	r := chi.NewRouter()

	// sign in 登陆界面
	r.Get("/signin", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "signin.gohtml"))))

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"signin.gohtml",
		))))

	fmt.Println("🚀 启动服务器端口:5001 ...")
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Println("Error listening on :5001")
		return
	}
}
