package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"math/rand"
	// "strconv"
)

/////  Gesture stuff, put somewhere else? //////
const (
        ROCK int = 0
        PAPER int = 1
        SCISSORS int = 2
)

// Returns true if p1 wins over p2
func battle(p1 int, p2 int) int {

	// Draw
	if p1 == p2 {
		return 0
	}

	// p1 win conditions
	if (p1 == ROCK && p2 == SCISSORS) ||
		(p1 == PAPER && p2 == ROCK) ||
		(p1 == SCISSORS && p2 == PAPER){
		return 1
	}

	return 2
}

func stringGesture(g string) int {
 	switch strings.ToUpper(g) {
	case "ROCK":
		return ROCK
	case "PAPER":
		return PAPER
	case "SCISSORS":
		return SCISSORS
	}
	return ROCK // WHAT TO DO ?
}


func gestureString(g int) string {
	switch g {
	case 0:
		return "Rock"
	case 1:
		return "Paper"
	case 2:
		return "Scissors"
	}
	return "None"
}

/////////////////////////////////////////////

func startpage(w http.ResponseWriter, r *http.Request) {

	// ME: Not really needed? Only for debugging
	// r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// // attention: If you do not call ParseForm method, the following data can not be obtained form
	// fmt.Println(r.Form) // print information on server side.
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])
	// for k, v := range r.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }

	// Actuall page load
	t, _ := template.ParseFiles("startpage.gtpl")
	t.Execute(w, nil)
}

func randomAI() int {
	return rand.Intn(3)
}

func practice(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("practice.gtpl")
	t.Execute(w, nil)

	if r.Method == "POST" {
		r.ParseForm()

		player := stringGesture(r.Form["Choice"][0])
		ai := randomAI()

		result := gestureString(player) + " VS: " + gestureString(ai)

		switch battle(player, ai) {
		case 0:
			result += " -> Draw!"
		case 1:
			result += " -> Player Wins!"
		case 2:
			result +=" -> AI Wins!"
		}

		fmt.Println(result)
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
