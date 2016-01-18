package main

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// here we define the absolute path to the view folder it takes the go root until the github folder.
var view, _ = filepath.Abs("../jschalkwijk/GolangTraining/blog/view")

type Post struct {
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
	Posts []Post
}

func main() {
	http.HandleFunc("/posts/", postsHandler)
	http.ListenAndServe(":8080", nil)
}

func postsHandler(w http.ResponseWriter, r *http.Request) {

	// returns the page struct with the assigned valus from the url and file contents
	p := getPosts()
	renderTemplate(w, p)

}

//func (p *Post) save() error {
//
//}

/*
  The function template.ParseFiles will read the contents of edit.html and return a *template.Template.
  The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter.
  The .Title and .Body dotted identifiers inside the template refer to p.Title and p.Body.
*/


func renderTemplate(w http.ResponseWriter, p []Post) {
	t, err := template.ParseFiles(view + "/posts.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getPosts() []Post {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	checkErr(err)
	fmt.Println("Connection with database Established")
	defer db.Close()
	defer fmt.Println("Connection with database Closed")

	rows, err := db.Query("SELECT * FROM posts")
	checkErr(err)

	//var posts Post
	collection := new(Collection)
	for rows.Next() {
		var post_id int
		var title string
		var description string
		var content string
		var keywords string
		var approved int
		var author string
		var date string
		var category_id int
		var trashed int

		err = rows.Scan(&post_id, &title, &description, &content,&keywords,&approved,
			&author,&date,&category_id,&trashed)
		checkErr(err)

		post := Post{post_id,title,description,content,keywords,approved,author,date,category_id,trashed}

		collection.Posts = append(collection.Posts , post)
	}
	//fmt.Println(collection.Posts)
	return collection.Posts
	//return Post{post_id,title,description,content,keywords,approved,author,date,category_id,trashed}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

