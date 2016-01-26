package categories

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

// categorie struct to create categories which will be added to the collection struct
type Category struct {
	Category_ID int
	Title string
	Description string
	Content string
	Keywords string
	Approved int
	Author string
	Cat_Type string
	Date string
	Parent_ID int
	Trashed int
}

var category_id int
var title string
var description string
var content string
var keywords string
var approved int
var author string
var cat_type string
var date string
var parent_id int
var trashed int

// Stores a single categorie, or multiple categories which we can then iterate over in the template
type Collection struct {
	Categories []Category
}

/*
  The function template.ParseFiles will read the contents of "".html and return a *template.Template.
  The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter.
  The .Title and .Body dotted identifiers inside the template refer to p.Title and p.Body.
*/


func renderTemplate(w http.ResponseWriter,name string, p []Category) {
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
	//The URL that the user queried, and then slice of the /categorie/ prefix.
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

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	params := splitURL(r,"categories")
	if(len(params) == 2){
		p := getSingleCategory(params[0],params[1])
		renderTemplate(w,"categories", p)
	} else if(len(params) == 3 && params[0] == "edit"){
		p := getSingleCategory(params[1],params[2])
		renderTemplate(w,"edit-category", p)
	} else if(len(params) == 3 && params[0] == "save"){
		editCategory(w,r,params[1],params[2])
	} else if(len(params) == 1 && params[0] == "new"){
		collection := new(Collection)
		p := collection.Categories
		renderTemplate(w,"new-category", p)
	} else if(len(params) == 1 && params[0] == "add-category"){
		newCategory(w, r)
	} else {
	// returns the page struct with the assigned values from the url and file contents
	p := getCategories()
	renderTemplate(w,"categories", p)
}

}

// Get all categories
func getCategories() []Category {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	checkErr(err)
	fmt.Println("Connection with database Established")
	defer db.Close()
	defer fmt.Println("Connection with database Closed")

	rows, err := db.Query("SELECT * FROM categories ORDER BY categorie_id DESC")
	checkErr(err)

	collection := new(Collection)
	for rows.Next() {
		err = rows.Scan(&category_id, &title, &description, &content,&keywords,&approved,
		&author,&cat_type,&date,&parent_id,&trashed)
		checkErr(err)

		category := Category{category_id,title,description,content,keywords,approved,author,cat_type,date,parent_id,trashed}

		collection.Categories = append(collection.Categories , category)
	}

	return collection.Categories
}

//Get a single categorie
func getSingleCategory(id string,category_title string) []Category {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	checkErr(err)
	fmt.Println("Connection established")
	defer db.Close()
	defer fmt.Println("Connection Closed")
	fmt.Println("SELECT * FROM categories WHERE categorie_id="+id+" AND title='"+category_title+"' LIMIT  1")
	rows := db.QueryRow("SELECT * FROM categories WHERE categorie_id=? AND title=? LIMIT 1", id,category_title)

	collection := new(Collection)

	err = rows.Scan(&category_id, &title, &description, &content,&keywords,&approved,
	&author,&date,&parent_id,&trashed)
	checkErr(err)

	category := Category{category_id,title,description,content,keywords,approved,author,cat_type,date,parent_id,trashed}

	collection.Categories = append(collection.Categories , category)

	//fmt.Println(collection.categories)
	return collection.Categories
}

// categorie Methods
func (p *Category) saveCategory() error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	defer db.Close()
	checkErr(err)
	stmt, err := db.Prepare("UPDATE categories SET title=?, description=?, content=? WHERE categorie_id=?")
	fmt.Println(stmt)
	checkErr(err)
	res, err := stmt.Exec(p.Title,p.Description,p.Content,p.Category_ID)
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	fmt.Println(res)
	return err
}

func (p *Category) addCategory() error {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/nerdcms_db?charset=utf8")
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO categories (title,description) VALUES(?,?) ")
	fmt.Println(stmt)
	checkErr(err)
	res, err := stmt.Exec(p.Title,p.Description)
	affect, err := res.RowsAffected()
	fmt.Println(affect)
	fmt.Println(res)
	checkErr(err)
	return err
}
// End category methods


func editCategory(w http.ResponseWriter, r *http.Request,id string,title string) {
	title = r.FormValue("title")
	description := r.FormValue("description")
	//category_id := r.FormValue("category_id")
	content := r.FormValue("content")
	new_id,error := strconv.Atoi(id)
	checkErr(error)
	p := &Category{Category_ID: new_id, Title: title,Description: description, Content: content}
	fmt.Println(p)
	err := p.saveCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/categories/"+id+"/"+title, http.StatusFound)
}

func newCategory(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")

	p := &Category{Title: title ,Description: description}
	fmt.Println(p)
	err := p.addCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/categories", http.StatusFound)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

