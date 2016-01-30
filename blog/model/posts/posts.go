package posts

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
	"strconv"
)

// here we define the absolute path to the view folder it takes the go root until the github folder.
var view, _ = filepath.Abs("../jschalkwijk/GolangTraining/blog/view")
var templates, _ = filepath.Abs("../jschalkwijk/GolangTraining/blog/templates")

// Post struct to create posts which will be added to the collection struct
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
	Posts []Data
}

/*
  The function template.ParseFiles will read the contents of "".html and return a *template.Template.
  The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter.
  The .Title and .Body dotted identifiers inside the template refer to p.Title and p.Body.
*/


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

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	params := splitURL(r,"posts")
	if(len(params) == 2){
		p := getSinglePost(params[0],params[1])
		renderTemplate(w,"posts", p)
	} else if(len(params) == 3 && params[0] == "edit"){
		p := getSinglePost(params[1],params[2])
		renderTemplate(w,"edit-post", p)
	} else if(len(params) == 3 && params[0] == "save"){
		editPost(w,r,params[1],params[2])
	} else if(len(params) == 1 && params[0] == "new"){
		collection := new(Collection)
		p := collection.Posts
		renderTemplate(w,"new-post", p)
	} else if(len(params) == 1 && params[0] == "add-post"){
		newPost(w, r)
	} else {
		// returns the page struct with the assigned values from the url and file contents
		p := getPosts()
		renderTemplate(w,"posts", p)
	}
}

// Get all Posts
func getPosts() []Data {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	checkErr(err)
	fmt.Println("Connection with database Established")
	defer db.Close()
	defer fmt.Println("Connection with database Closed")

	rows, err := db.Query("SELECT * FROM posts ORDER BY post_id DESC")
	checkErr(err)

	collection := new(Collection)
	for rows.Next() {
		err = rows.Scan(&post_id, &title, &description, &content,&keywords,&approved,
			&author,&date,&category_id,&trashed)
		checkErr(err)

		post := Data{post_id,title,description,content,keywords,approved,author,date,category_id,trashed}

		collection.Posts = append(collection.Posts , post)
	}

	return collection.Posts
}

//Get a single Post
func getSinglePost(id string,post_title string) []Data {
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

	post := Data{post_id,title,description,content,keywords,approved,author,date,category_id,trashed}

	collection.Posts = append(collection.Posts , post)

	//fmt.Println(collection.Posts)
	return collection.Posts
}

// Post Methods
func (p *Data) savePost() error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	defer db.Close()
	checkErr(err)
	stmt, err := db.Prepare("UPDATE posts SET title=?, description=?, content=? WHERE post_id=?")
	fmt.Println(stmt)
	checkErr(err)
	res, err := stmt.Exec(p.Title,p.Description,p.Content,p.Post_ID)
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	fmt.Println(res)
	return err
}

func (p *Data) addPost() error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO posts (title,description,content) VALUES(?,?,?) ")
	fmt.Println(stmt)
	checkErr(err)
	res, err := stmt.Exec(p.Title,p.Description,p.Content)
	affect, err := res.RowsAffected()
	fmt.Println(affect)
	fmt.Println(res)
	checkErr(err)
	return err
}
// End Post methods


func editPost(w http.ResponseWriter, r *http.Request,id string,title string) {
	title = r.FormValue("title")
	description := r.FormValue("description")
	//category_id := r.FormValue("category_id")
	content := r.FormValue("content")
	new_id,error := strconv.Atoi(id)
	checkErr(error)
	p := &Data{Post_ID: new_id, Title: title,Description: description, Content: content}
	fmt.Println(p)
	err := p.savePost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/posts/"+id+"/"+title, http.StatusFound)
}
func newPost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	//category_id := r.FormValue("category_id")
	content := r.FormValue("content")

	p := &Data{Title: title ,Description: description, Content: content}
	fmt.Println(p)
	err := p.addPost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/posts", http.StatusFound)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
