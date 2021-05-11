package elo

import (
	"testing"
)

func TestExpectedScores(t *testing.T) {
	playerOne := NewPlayer(2300.00, 32)
	playerTwo := NewPlayer(2200.00, 32)
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
	playerOne := NewPlayer(2300.00, 32)
	playerTwo := NewPlayer(2300.00, 32)
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

func TestCalculateNewRating(t *testing.T) {
	p := NewPlayer(1613, 32)
	p.CalculateNewRating(2.88, 2.5)
	if p.Rating() != 1601 {
		t.Errorf("expected rating to be 1601, actual %f\n", p.Rating())
	}
}

func TestCalculateNewRatingDeep(t *testing.T) {
	elo := ELOCalculator{}
	player := NewPlayer(1613, 32)
	expectedScore := 0.0
	s, _ := elo.ExpectedScores(player, NewPlayer(1609, 32))
	expectedScore = s
	s, _ = elo.ExpectedScores(player, NewPlayer(1477, 32))
	expectedScore += s
	s, _ = elo.ExpectedScores(player, NewPlayer(1388, 32))
	expectedScore += s
	s, _ = elo.ExpectedScores(player, NewPlayer(1586, 32))
	expectedScore += s
	s, _ = elo.ExpectedScores(player, NewPlayer(1720, 32))
	expectedScore += s
	actualScore := elo.Score(2, 1, 2)
	player.CalculateNewRating(expectedScore, actualScore)

	t.Run("validate new rating based on actual score...", func(t *testing.T) {
		if player.Rating() != 1601 {
			t.Errorf("expected rating to be 1601, actual %f\n", player.Rating())
		}
	})
}
