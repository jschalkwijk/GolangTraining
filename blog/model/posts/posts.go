package posts

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

// here we define the absolute path to the view folder it takes the go root until the github folder.
var view, _ = filepath.Abs("../jschalkwijk/GolangTraining/blog/view")

// Post struct to create posts which will be added to the collection struct
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

// Stores a single post, or multiple posts which we can then iterate over in the template
type Collection struct {
	Posts []Post
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	params := splitURL(r)
	if(len(params) == 2){
		p := getSinglePosts(params[0],params[1])
		renderTemplate(w, p)
	} else {
		// returns the page struct with the assigned values from the url and file contents
		p := getPosts()
		renderTemplate(w, p)
	}
}

func splitURL(r *http.Request) []string {
	//The URL that the user queried.
	path := r.URL.Path[len("/posts/"):]
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

	collection := new(Collection)
	for rows.Next() {
		err = rows.Scan(&post_id, &title, &description, &content,&keywords,&approved,
			&author,&date,&category_id,&trashed)
		checkErr(err)

		post := Post{post_id,title,description,content,keywords,approved,author,date,category_id,trashed}

		collection.Posts = append(collection.Posts , post)
	}

	return collection.Posts
}

func getSinglePosts(id string,post_title string) []Post {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	checkErr(err)
	fmt.Println("Connection established")
	defer db.Close()
	defer fmt.Println("Connection Closed")
	fmt.Println("SELECT * FROM posts WHERE post_id="+id+" AND title='"+post_title+"' LIMIT  1")
	rows := db.QueryRow("SELECT * FROM posts WHERE post_id=? AND title=? LIMIT 1", id,post_title)

	collection := new(Collection)

	err = rows.Scan(&post_id, &title, &description, &content,&keywords,&approved,
		&author,&date,&category_id,&trashed)
	checkErr(err)

	post := Post{post_id,title,description,content,keywords,approved,author,date,category_id,trashed}

	collection.Posts = append(collection.Posts , post)

	//fmt.Println(collection.Posts)
	return collection.Posts
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}