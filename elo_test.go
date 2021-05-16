package elo

import (
	"testing"
)

func TestExpectedScores(t *testing.T) {
	playerOne := NewPlayer(2300, 32)
	playerTwo := NewPlayer(2200, 32)
	e := ELOCalculator{}
	scoreOne, scoreTwo := e.ExpectedScores(playerOne, playerTwo)

	if scoreOne != 0.64 {
		t.Errorf("epected playerOne score to be 0.64, actual %f\n", scoreOne)
	}

	if scoreTwo != 0.36 {
		t.Errorf("epected playerTwo score to be 0.36, actual %f\n", scoreTwo)
	}
}

func TestExpectedScoresEquallyMatched(t *testing.T) {
	playerOne := NewPlayer(2300, 32)
	playerTwo := NewPlayer(2300, 32)
	e := ELOCalculator{}
	scoreOne, scoreTwo := e.ExpectedScores(playerOne, playerTwo)

	if scoreOne != scoreTwo {
		t.Errorf("expected player scores to match, actual playerOne [%f], actual playerTwo [%f]\n", scoreOne, scoreTwo)
	}
}

func TestScore(t *testing.T) {
	e := ELOCalculator{}
	score := e.Score(3, 1, 1)
	if score != 3.5 {
		t.Errorf("expected score to be 3.5, actual %f\n", score)
	}

	score = e.Score(0, 2, 3)
	if score != 1.0 {
		t.Errorf("expected score to be 1.0, actual %f\n", score)
	}
}
