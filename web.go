package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func startpage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	t, _ := template.ParseFiles("startpage.gtpl")
	t.Execute(w, nil)

	// fmt.Fprintf(w, "RPS8000!") // write data to response
	// fmt.Fprintf(w, "");

}

func practice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("practice.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("User choice: ", r.Form["Choice"])
	}
}

func main() {
	http.HandleFunc("/", startpage) // setting router rule
	http.HandleFunc("/practice", practice)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
