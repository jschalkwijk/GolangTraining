package post

import (
	"net/http"
	"github.com/jschalkwijk/GolangTraining/blog/model/posts"
	"github.com/jschalkwijk/GolangTraining/blog/controller"
	"github.com/gorilla/mux"
)


func Posts(w http.ResponseWriter, r *http.Request) {

 if(len(params) == 3 && params[0] == "edit"){
		p := posts.GetSinglePost(params[1],params[2])
		controller.RenderTemplate(w,"edit-post", p)
	} else if(len(params) == 3 && params[0] == "save"){
		posts.EditPost(w,r,params[1],params[2])
	} else if(len(params) == 1 && params[0] == "new"){

	} else if(len(params) == 1 && params[0] == "add-post"){
		posts.NewPost(w, r)
	} else {
		// returns the page struct with the assigned values from the url and file contents
		p := posts.GetPosts()
	 	controller.RenderTemplate(w,"posts", p)
	}
}

func Single(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	post_title := vars["title"]
	p := posts.GetSinglePost(id,post_title)
	controller.RenderTemplate(w,"posts", p)
}

func New(w http.ResponseWriter, r *http.Request){
	collection := new(posts.Data)
	p := collection
	controller.RenderTemplate(w,"new-post", p)
}