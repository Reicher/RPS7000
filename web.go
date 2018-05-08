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

type gameSet struct {
	Id        int
	Player    int
	AI        int
	Summary   string
}

type PracticePageData struct {
	Result string
}

type GamePageData struct {
	Stats []gameSet
}

func randomAI() int {
	return 	rand.Intn(3)
}

func battle_to_string(p1 int, p2 int) string {
	desc := "Player 1 " + gesture.ToString(p1) +  " VS: Player 2 " + gesture.ToString(p2)

	switch gesture.Battle(p1, p2) {
	case 0:
		desc += " => Draw!"
	case 1:
		desc += " => Player 1 Wins!"
	case 2:
		desc += " => Player 2 Wins!"
	}

	return desc
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
		fmt.Println(battle_to_string(player, ai))
	}

	t.Execute(w, data)
}

//
func match(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/home/rrr/go/src/github.com/reicher/RPS7000/assets/match.gtpl")

	data := GamePageData{
		Stats:  []gameSet{},
	}

	if r.Method == "POST" {
		r.ParseForm()

		set := gameSet{Id:0}
		set.Player = gesture.FromString(r.Form["Choice"][0])
		set.AI = randomAI()
		set.Summary = battle_to_string(set.Player, set.AI)

		data.Stats = append(data.Stats, set)
		fmt.Println("Sets: " + string(len(data.Stats)))
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
	http.HandleFunc("/match", match)
	err := http.ListenAndServe(":9090", nil) // setting listening port

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
