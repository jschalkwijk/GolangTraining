// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
    "net/url"
	"path/filepath"
	"fmt"
)


// here we define the absolute path to the view folder it takes the go root until the github folder.
var view, _ = filepath.Abs("../jschalkwijk/GolangTraining/wiki2/view")

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)

}

// Is this an object??
/*
	Here we define the Page structure, we will use a title and a body. The body must be of
	type []byte, and later be converted to a string inside the template. the bites correspond to integers
	that correspond to UTF-8 characters
*/

type Page struct {
	Title string
	Body  []byte
	Name string
}

type Person struct {
	Name string
	Age  int
}

/*
	The Page struct describes how page data will be stored in memory.
	But what about persistent storage? We can address that by creating a save method on Page:

	So this is a method, related to Page, and it returns an error or saves the ne text to the file
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(view + "/" + filename, p.Body, 0600)
}

/*
	loadPage takes a parameter title of type string, and returns a page struct and an error
	the title + .txt is the filename, the contents of the file will be read by ioutil.readfile
	and will return the contents in bytes. if there are no errors, we return the struct with
	assigned values.
*/

func loadPage(title string) (*Page, error,*Person, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(view + "/" + filename)
	if err != nil {
		return nil, err,nil,err
	}

	jorn := Person{Name: "jorn",Age:24}
	return &Page{Title: title, Body: body, Name: "Jorn"}, nil, &jorn, nil
}

/*
  First, this function extracts the page title from r.URL.Path, the path component of the request URL.
  The Path is re-sliced with [len("/view/"):] to drop the leading "/view/" component of the request path.
  This is because the path will invariably begin with "/view/", which is not part of the page's title.


*/
func viewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/view/"):]
	u, err := url.Parse(title)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)

	// returns the page struct with the assigned valus from the url and file contents
	p, err,h,error := loadPage(title)

	if err != nil || error != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p,h)

}

/*
  The function template.ParseFiles will read the contents of edit-post.html and return a *template.Template.
  The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter.
  The .Title and .Body dotted identifiers inside the template refer to p.Title and p.Body.
*/

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page, h *Person) {
	t, err := template.ParseFiles(view + "/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
/*
  The function template.ParseFiles will read the contents of edit-post.html and return a *template.Template.
  The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter.
  The .Title and .Body dotted identifiers refer to p.Title and p.Body.
*/

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p, err,h,error := loadPage(title)
	if err != nil || error != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p,h)
}


