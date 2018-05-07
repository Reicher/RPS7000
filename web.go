package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"math/rand"
	"github.com/reicher/RPS7000/gesture"
)

func randomAI() int {
	return rand.Intn(3)
}

type PracticePageData struct {
	Result string
}

func practice(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/home/rrr/go/src/github.com/reicher/RPS7000/assets/practice.gtpl")

	data := PracticePageData{
		Result: "",
	}

	if r.Method == "POST" {
		r.ParseForm()

		player := gesture.FromString(r.Form["Choice"][0])
		ai := randomAI()

		data.Result = "Player " + gesture.ToString(player)+ " VS: AI " + gesture.ToString(ai)

		switch gesture.Battle(player, ai) {
		case 0:
			data.Result += " => Draw!"
		case 1:
			data.Result += " => Player Wins!"
		case 2:
			data.Result +=" => AI Wins!"
		}
	}
	t.Execute(w, data)
}

func startpage(w http.ResponseWriter, r *http.Request) {

	// ME: Not really needed? Only for debugging
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)

	// // attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	// Actuall page load
	t, err := template.ParseFiles("/home/rrr/go/src/github.com/reicher/RPS7000/assets/startpage.gtpl")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, nil)
	}
}

func main() {
	fmt.Println("localhost:9090")

	rand.Seed(40) // Try changing this number!

	http.HandleFunc("/", startpage) // setting router rule
	http.HandleFunc("/practice", practice)
	err := http.ListenAndServe(":9090", nil) // setting listening port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
