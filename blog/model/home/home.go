package home

import (
	"html/template"
	"net/http"
	"strings"
	"path/filepath"
)

var view, _ = filepath.Abs("../jschalkwijk/GolangTraining/blog/view")
var templates, _ = filepath.Abs("../jschalkwijk/GolangTraining/blog/templates")

type Data struct {
	Post_ID int
	Title string
	Description string
	Content string
	Keywords string
	Approved int
	Author string
	Date string
	Category_ID int
	Trashed int
}

type Collection struct {
	Posts []Data
}

func DashboardHandler(w http.ResponseWriter, r *http.Request){
		//params := splitURL(r,"")
		collection := new(Collection)
		p := collection.Posts
		renderTemplate(w,"index", p)
}

func renderTemplate(w http.ResponseWriter,name string, p []Data) {
	t, err := template.ParseFiles(templates+"/"+"header.html",view + "/" + name + ".html",templates+"/"+"footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w,"header",nil)
	t.ExecuteTemplate(w,name,p)
	t.ExecuteTemplate(w,"footer",nil)
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func splitURL(r *http.Request, prefix string) []string {
	//The URL that the user queried, and then slice of the /post/ prefix.
	path := r.URL.Path[len("/"+ prefix +"/"):]
	path = strings.TrimSpace(path)
	//Cut off the leading and trailing forward slashes, if they exist.
	//This cuts off the leading forward slash.
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	//This cuts off the trailing forward slash.
	if strings.HasSuffix(path, "/") {
		removeSlash := len(path) - 1
		path = path[:removeSlash]
	}
	//We need to isolate the individual parameters of the path.
	params := strings.Split(path, "/")
	return params
}