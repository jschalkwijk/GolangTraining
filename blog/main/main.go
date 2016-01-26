package main

import (
	"net/http"
	"github.com/jschalkwijk/GolangTraining/blog/model/posts"
	"github.com/jschalkwijk/GolangTraining/blog/model/categories"
)

func main() {
	http.HandleFunc("/posts/", posts.PostsHandler)
	http.HandleFunc("/categories/", categories.CategoriesHandler)
	http.ListenAndServe(":8080", nil)
}


