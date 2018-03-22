package gesture

import (
	"testing"
	"github.com/reicher/RPS7000/gesture"
)

func TestBattleLogic(t *testing.T) {

	if gesture.Battle(gesture.ROCK, gesture.SCISSORS) != 1 {
		t.Errorf("Rock should beat scissors")
	}

	if gesture.Battle(gesture.SCISSORS, gesture.PAPER) != 1{
		t.Errorf("Scissors should beat paper")
	}

	if gesture.Battle(gesture.PAPER, gesture.ROCK) != 1{
		t.Errorf("Paper should beat rock")
	}

	if gesture.Battle(gesture.ROCK, gesture.ROCK) != 0{
		t.Errorf("Rock and ROCK should be draw")
	}
}
