package gesture

import (
	"strings"
)

const (
        ROCK int = 0
        PAPER int = 1
        SCISSORS int = 2
)

func Battle(p1 int, p2 int) int {

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

func FromString(g string) int {
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

func ToString(g int) string {
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
