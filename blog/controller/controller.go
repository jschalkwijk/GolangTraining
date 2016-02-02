package controller

import (
	"html/template"
	"net/http"
	"strings"
	"path/filepath"
	"github.com/jschalkwijk/GolangTraining/blog/model/posts"
)

var view, _ = filepath.Abs("../jschalkwijk/GolangTraining/blog/view")
var templates, _ = filepath.Abs("../jschalkwijk/GolangTraining/blog/templates")

func RenderTemplate(w http.ResponseWriter,name string, data *posts.Data) {
	t, err := template.ParseFiles(templates+"/"+"header.html",view + "/" + name + ".html",templates+"/"+"footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplate(w,"header",nil)
	t.ExecuteTemplate(w,name,data)
	t.ExecuteTemplate(w,"footer",nil)
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SplitURL(r *http.Request, prefix string) []string {
	//The URL that the user queried, and then slice of the /post/ prefix.
	path := r.URL.Path[len(prefix):]
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

//var path string
//
//func Ctrl(w http.ResponseWriter,r *http.Request){
//	path = r.URL.Path
//	cleanPath := r.URL.Path
//	if strings.HasPrefix(cleanPath, "/") {
//		cleanPath = cleanPath[1:]
//	}
//	//This cuts off the trailing forward slash.
//	if strings.HasSuffix(cleanPath, "/") {
//		removeSlash := len(cleanPath) - 1
//		cleanPath = cleanPath[:removeSlash]
//	}
//	fmt.Println(path,cleanPath)
//	http.HandleFunc(path, Posts)
//}