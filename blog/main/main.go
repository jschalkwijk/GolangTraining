package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jschalkwijk/GolangTraining/blog/model/home"
	"github.com/jschalkwijk/GolangTraining/blog/controller/posts"
	"github.com/jschalkwijk/GolangTraining/blog/controller/categories"
)

func main() {
	r := mux.NewRouter()

	// Index
	r.HandleFunc("/", home.DashboardHandler)
	// Posts
	r.HandleFunc("/posts/", posts.Index)
		p := r.PathPrefix("/posts").Subrouter()
		p.HandleFunc("/{id:[0-9]+}/{title}", posts.Single)
		p.HandleFunc("/new", posts.New)
		p.HandleFunc("/edit/{id:[0-9]+}/{title}", posts.Edit)
		p.HandleFunc("/save/{id:[0-9]+}/{title}", posts.Save)
		p.HandleFunc("/add-post", posts.Add)
	// Categories
	r.HandleFunc("/categories/", categories.Index)
		c := r.PathPrefix("/categories").Subrouter()
		c.HandleFunc("/{id:[0-9]+}/{title}", categories.Single)
		c.HandleFunc("/new", categories.New)
		c.HandleFunc("/edit/{id:[0-9]+}/{title}", categories.Edit)
		c.HandleFunc("/save/{id:[0-9]+}/{title}", categories.Save)
		c.HandleFunc("/add-category", categories.Add)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)

}