package controllers

import "net/http"

// 一个固定的账户，对密码进行了散列处理 邮箱:密码
var users = map[string]string{
	"admin@yoi.com": "Yoi#QWE@2024!",
}

// Signin 登陆 Post /signin 登陆界面 email:password
func Signin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	if users[email] == password {
		signin(w, r)
		return
	}
	http.Redirect(w, r, "/signin", http.StatusFound)
}

// 通过cookie中的token判断用户是否已经登陆
func authenticated(r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		return false
	}
	return cookie != nil && cookie.Value == "Yoi#QWE@2024!"
}

// 已经登陆则保持回话并且重定向到/home
func signin(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: "Yoi#QWE@2024!",
	})
	http.Redirect(w, r, "/home", http.StatusFound)
}

// AuthMiddleware 验证中间件
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authenticated(r) {
			http.Redirect(w, r, "/signin", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Signout 退出登陆
func Signout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})
}
