package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"math/rand"
)

func startpage(w http.ResponseWriter, r *http.Request) {

	// ME: Not really needed? Only for debugging
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

	// Actuall page load
	t, _ := template.ParseFiles("startpage.gtpl")
	t.Execute(w, nil)
}

func battle(){

}

func randomAI(){
	types := []string{
		"Rock",
		"Paper",
		"Scissors",
	}
	fmt.Println(types[rand.Intn(len(types))])
}

func practice(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	t, _ := template.ParseFiles("practice.gtpl")
	t.Execute(w, nil)

	if r.Method == "POST" {
		r.ParseForm()
		// logic part of log in
		choice := r.Form["Choice"]
		fmt.Println("User choice: ", choice)
		randomAI()
	}
}

func main() {
	rand.Seed(40) // Try changing this number!
	http.HandleFunc("/", startpage) // setting router rule
	http.HandleFunc("/practice", practice)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
