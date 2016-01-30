package main

import (
	"net/http"
	"github.com/jschalkwijk/GolangTraining/blog/model/posts"
	"github.com/jschalkwijk/GolangTraining/blog/model/categories"
	"github.com/jschalkwijk/GolangTraining/blog/model/home"
)

func main() {
	http.HandleFunc("/", home.DashboardHandler)
	http.HandleFunc("/posts/", posts.PostsHandler)
	http.HandleFunc("/categories/", categories.CategoriesHandler)
	//http.Handle("/static/css", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css"))))
	http.ListenAndServe(":8080", nil)
}


