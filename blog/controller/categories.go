package controller

import (
	"net/http"
	"github.com/jschalkwijk/GolangTraining/blog/model/categories"
)
func Categories(w http.ResponseWriter, r *http.Request) {
	params := SplitURL(r, "/categories/")
	if (len(params) == 2) {
		p := categories.GetSingleCategory(params[0], params[1])
		categories.RenderTemplate(w, "categories", p)
	} else if (len(params) == 3 && params[0] == "edit") {
		p := categories.GetSingleCategory(params[1], params[2])
		categories.RenderTemplate(w, "edit-category", p)
	} else if (len(params) == 3 && params[0] == "save") {
		categories.EditCategory(w, r, params[1], params[2])
	} else if (len(params) == 1 && params[0] == "new") {
		collection := new(categories.Data)
		p := collection
		categories.RenderTemplate(w, "new-category", p)
	} else if (len(params) == 1 && params[0] == "add-category") {
		categories.NewCategory(w, r)
	} else {
		// returns the page struct with the assigned values from the url and file contents
		p := categories.GetCategories()
		categories.RenderTemplate(w, "categories", p)
	}
}