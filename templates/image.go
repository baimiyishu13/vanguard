package templates

import "net/http"

// Image 通过url访问 ./templates/img 时，返回图片 ./templates/img/signin.jpg
func Image(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/img/signin.jpg")
}
