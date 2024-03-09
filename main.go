package main

import (
	"fmt"
	"github.com/baimiyishu13/vanguard/controllers"
	"github.com/baimiyishu13/vanguard/templates"
	"github.com/baimiyishu13/vanguard/views"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"log"
	"net/http"
)

func main() {
	// test
	r := chi.NewRouter()

	// sign in 登陆界面
	r.Get("/signin", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"signin.gohtml",
		))))

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"home.gohtml",
		))))

	// 通过url访问 ./templates/img 时，返回图片
	r.Get("/img/signin.jpg", templates.Image)

	// CSRF 保护
	var csrfKey = "9IDAuQlSlpBasivx1O5m0xp0nEYkb3bG"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: set this
		csrf.Secure(false),
	)

	fmt.Println("🚀 启动服务器端口:5001 ...")
	err := http.ListenAndServe(":5001", csrfMw(r))
	if err != nil {
		log.Println("Error listening on :5001")
		return
	}
}
