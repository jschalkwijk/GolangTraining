package categories
import (
	"net/http"
	"github.com/jschalkwijk/GolangTraining/blog/model/categories"
	//"github.com/jschalkwijk/GolangTraining/blog/controller"
	"github.com/gorilla/mux"
)


func Index(w http.ResponseWriter, r *http.Request) {
	p := categories.GetCategories()
	categories.RenderTemplate(w,"categories", p)
}

func Single(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	post_title := vars["title"]
	p := categories.GetSingleCategory(id,post_title)
	categories.RenderTemplate(w,"categories", p)
}

func New(w http.ResponseWriter, r *http.Request){
collection := new(categories.Data)
p := collection
categories.RenderTemplate(w,"new-category", p)
}

func Edit(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	post_title := vars["title"]
	p := categories.GetSingleCategory(id,post_title)
	categories.RenderTemplate(w,"edit-category", p)
}
func Save(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	post_title := vars["title"]
	categories.EditCategory(w,r,id,post_title)
}

func Add(w http.ResponseWriter, r *http.Request){
	categories.NewCategory(w, r)
}