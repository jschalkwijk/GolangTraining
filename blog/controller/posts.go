package controller

import (
	"net/http"
	"github.com/jschalkwijk/GolangTraining/blog/model/posts"
	//"github.com/gorilla/mux"
)


func Posts(w http.ResponseWriter, r *http.Request) {
//	route := mux.NewRouter()
//	p :=route.HandleFunc("/posts/{id}/{title}", posts.GetSinglePost)
//	RenderTemplate(w,"posts", p)

	params := SplitURL(r,"/posts/")
	if(len(params) == 2){
		p := posts.GetSinglePost(params[0],params[1])
		RenderTemplate(w,"posts", p)
	} else if(len(params) == 3 && params[0] == "edit"){
		p := posts.GetSinglePost(params[1],params[2])
		RenderTemplate(w,"edit-post", p)
	} else if(len(params) == 3 && params[0] == "save"){
		posts.EditPost(w,r,params[1],params[2])
	} else if(len(params) == 1 && params[0] == "new"){
		collection := new(posts.Data)
		p := collection
		RenderTemplate(w,"new-post", p)
	} else if(len(params) == 1 && params[0] == "add-post"){
		posts.NewPost(w, r)
	} else {
		// returns the page struct with the assigned values from the url and file contents
		p := posts.GetPosts()
		RenderTemplate(w,"posts", p)
	}
}
