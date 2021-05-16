package elo

import (
	"testing"
)

func TestExpectedScores(t *testing.T) {
	playerOne := 2300
	playerTwo := 2200
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
	playerOne := 2300
	playerTwo := 2300
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

func TestKFactorFromRating(t *testing.T) {
	e := ELOCalculator{}
	e.SetKFactorFromRating(2000)
	t.Run("validate kfactor for rating below 2100...", func(t *testing.T) {
		if e.kfactor != 32 {
			t.Errorf("expected kfactor to be 32, actual %f\n", e.kfactor)
		}
	})
	e.SetKFactorFromRating(2500)
	t.Run("validate kfactor for rating above 2400...", func(t *testing.T) {
		if e.kfactor != 16 {
			t.Errorf("expected kfactor to be 16, actual %f\n", e.kfactor)
		}
	})
	e.SetKFactorFromRating(2300)
	t.Run("validate kfactor when rating is between 2100  and  2400 (inclusive)...", func(t *testing.T) {
		if e.kfactor != 24 {
			t.Errorf("expected kfactor to be 24, actual %f\n", e.kfactor)
		}
	})
}

func TestKFactorGamesPlayed(t *testing.T) {
	e := ELOCalculator{}
	e.SetKFactorFromGamesPlayed(30, 5)
	t.Run("validate kfactor from rating based on 30 games previously played, and 5 played today...", func(t *testing.T) {
		if e.kfactor != 22.857142857142858 {
			t.Errorf("expected kfactor to be 22.857142857142858, actual %f\n", e.kfactor)
		}
	})
}

func TestCalculateNewRating(t *testing.T) {
	e := ELOCalculator{}
	e.SetKFactor(32)
	rating := e.CalculateNewRating(1613, 2.88, 2.5)
	if rating != 1601 {
		t.Errorf("expected rating to be 1601, actual %d\n", rating)
	}
}

func TestCalculateNewRatingWithRating(t *testing.T) {
	e := ELOCalculator{}
	e.SetKFactorFromRating(1613)
	rating := e.CalculateNewRating(1613, 2.88, 2.5)
	if rating != 1601 {
		t.Errorf("expected rating to be 1601, actual %d\n", rating)
	}
}

func TestCalculateNewRatingDeep(t *testing.T) {
	elo := ELOCalculator{}
	elo.SetKFactorFromRating(1613)
	expectedScore := 0.0
	s, _ := elo.ExpectedScores(1613, 1609)
	expectedScore = s
	s, _ = elo.ExpectedScores(1613, 1477)
	expectedScore += s
	s, _ = elo.ExpectedScores(1613, 1388)
	expectedScore += s
	s, _ = elo.ExpectedScores(1613, 1586)
	expectedScore += s
	s, _ = elo.ExpectedScores(1613, 1720)
	expectedScore += s
	actualScore := elo.Score(2, 1, 2)
	rating := elo.CalculateNewRating(1613, expectedScore, actualScore)

	t.Run("validate new rating based on actual score...", func(t *testing.T) {
		if rating != 1601 {
			t.Errorf("expected rating to be 1601, actual %d\n", rating)
		}
	})
}
