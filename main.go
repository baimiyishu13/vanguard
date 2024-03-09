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

	// 	未登陆成功前访问如何路径都会被重定向到/signin
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/signin", http.StatusFound)
	})

	// sign in 登陆界面
	r.Post("/signin", controllers.Signin)
	r.Get("/signin", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS,
			"signin.gohtml",
		))))

	// 登陆成功后才可以访问其他路径
	r.Group(func(r chi.Router) {
		// 验证中间件
		r.Use(controllers.AuthMiddleware)

		// home 主页
		r.Get("/home", controllers.StaticHandler(
			views.Must(views.ParseFS(templates.FS,
				"home.gohtml",
			))))

		// 退出登陆
		r.Get("/signout", controllers.Signout)
	})

	// 通过url访问时，返回图片
	r.Get("/img/signin.jpg", templates.Image)

	fmt.Println("🚀 启动服务器端口:5001 ...")
	err := http.ListenAndServe(":5001", r)
	if err != nil {
		log.Println("Error listening on :5001")
		return
	}
}
