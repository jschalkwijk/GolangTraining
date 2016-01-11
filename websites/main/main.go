package main


import (
"fmt"
"net/http"
"io/ioutil"
"path/filepath"
)

var html, _ = filepath.Abs("../jschalkwijk/GolangTraining/websites/html")

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func renderPage() (y string) {
	// read whole the file
	test, err := ioutil.ReadFile(html+"/test.html")
	if err != nil {
		panic(err)
	}
	 y = string(test)
	return y

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, renderPage())
}

