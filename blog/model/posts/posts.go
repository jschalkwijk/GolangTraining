package posts

import (
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

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
