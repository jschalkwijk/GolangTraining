package main

import (
	"net/http"
	//"github.com/jschalkwijk/GolangTraining/blog/model/posts"
	//"github.com/jschalkwijk/GolangTraining/blog/model/categories"
	"github.com/jschalkwijk/GolangTraining/blog/model/home"
	"github.com/jschalkwijk/GolangTraining/blog/controller"
)

var url string
func main() {
//	http.HandleFunc("/",controller.Ctrl)
	http.HandleFunc("/", home.DashboardHandler)
	http.HandleFunc("/posts/", controller.Posts)
	http.HandleFunc("/categories/", controller.Categories)
	//http.Handle("/static/css", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
	http.ListenAndServe(":8080", nil)
}