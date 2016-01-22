package main

import (
	"net/http"
	"github.com/jschalkwijk/GolangTraining/blog/model/posts"
)

func main() {
	http.HandleFunc("/posts/", posts.PostsHandler)
	http.ListenAndServe(":8080", nil)
}


