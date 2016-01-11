package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

type Page struct {
	Title string
	Body []byte
}
//The Page struct describes how page data will be stored in memory. But what about persistent storage?
// We can address that by creating a save method on Page:

//"This is a method named save that takes as its receiver p, a pointer to Page .
// It takes no parameters, and returns a value of type error."

// This method will save the Page's Body to a text file. For simplicity, we will use the Title as the file name.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename,p.Body,0600)
}

//The function loadPage constructs the file name from the title parameter,
//reads the file's contents into a new variable body,
//and returns a pointer to a Page literal constructed with the proper title and body values.

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}