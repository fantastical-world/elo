package elo

import (
	"testing"
)

func Test_New(t *testing.T) {
	e := New(32)
	t.Run("validate the calculator's kfactor is set correctly...", func(t *testing.T) {
		if e.kfactor != 32 {
			t.Errorf("expected kfactor to be 32, actual %f\n", e.kfactor)
		}
	})
}
func TestELOCalculator_ExpectedScores(t *testing.T) {
	e := ELOCalculator{}
	t.Run("validate scores are correct for players with different ratings...", func(t *testing.T) {
		playerOne, playerTwo := 2300, 2200
		scoreOne, scoreTwo := e.ExpectedScores(playerOne, playerTwo)
		if scoreOne != 0.64 {
			t.Errorf("epected playerOne score to be 0.64, actual %f\n", scoreOne)
		}
		if scoreTwo != 0.36 {
			t.Errorf("epected playerTwo score to be 0.36, actual %f\n", scoreTwo)
		}
	})

	t.Run("validate scores are correct for players with the same rating...", func(t *testing.T) {
		playerOne, playerTwo := 2300, 2300
		scoreOne, scoreTwo := e.ExpectedScores(playerOne, playerTwo)
		if scoreOne != scoreTwo {
			t.Errorf("expected player scores to match, actual playerOne [%f], actual playerTwo [%f]\n", scoreOne, scoreTwo)
		}
	})
}

func TestELOCalculator_Score(t *testing.T) {
	e := ELOCalculator{}
	t.Run("validate score for 3 wins, 1 draw, and 1 loss...", func(t *testing.T) {
		score := e.Score(3, 1, 1)
		if score != 3.5 {
			t.Errorf("expected score to be 3.5, actual %f\n", score)
		}
	})

	t.Run("validate score for 0 wins, 2 draws, 3 losses...", func(t *testing.T) {
		score := e.Score(0, 2, 3)
		if score != 1.0 {
			t.Errorf("expected score to be 1.0, actual %f\n", score)
		}
	})

	t.Run("validate score for 0 wins, 0 draws, 0 losses...", func(t *testing.T) {
		score := e.Score(0, 0, 0)
		if score != 0.0 {
			t.Errorf("expected score to be 0.0, actual %f\n", score)
		}
	})
}

func TestELOCalculator_SetKFactorFromRating(t *testing.T) {
	e := ELOCalculator{}
	t.Run("validate kfactor for rating below 2100...", func(t *testing.T) {
		e.SetKFactorFromRating(2000)
		if e.kfactor != 32 {
			t.Errorf("expected kfactor to be 32, actual %f\n", e.kfactor)
		}
	})

	t.Run("validate kfactor for rating above 2400...", func(t *testing.T) {
		e.SetKFactorFromRating(2500)
		if e.kfactor != 16 {
			t.Errorf("expected kfactor to be 16, actual %f\n", e.kfactor)
		}
	})

	t.Run("validate kfactor when rating is between 2100  and  2400 (inclusive)...", func(t *testing.T) {
		e.SetKFactorFromRating(2300)
		if e.kfactor != 24 {
			t.Errorf("expected kfactor to be 24, actual %f\n", e.kfactor)
		}
	})
}

func TestELOCalculator_SetKFactorFromGamesPlayed(t *testing.T) {
	e := ELOCalculator{}
	t.Run("validate kfactor from rating based on 30 games previously played, and 5 played today...", func(t *testing.T) {
		e.SetKFactorFromGamesPlayed(30, 5)
		if e.kfactor != 22.857142857142858 {
			t.Errorf("expected kfactor to be 22.857142857142858, actual %f\n", e.kfactor)
		}
	})
}

func TestELOCalculator_CalculateNewRating(t *testing.T) {
	e := ELOCalculator{}
	t.Run("validate rating drops if score is lower than expected...", func(t *testing.T) {
		e.SetKFactor(32)
		rating := e.CalculateNewRating(1613, 2.88, 2.5)
		if rating != 1601 {
			t.Errorf("expected rating to be 1601, actual %d\n", rating)
		}
	})

	t.Run("validate rating increases if score is higher than expected...", func(t *testing.T) {
		e.SetKFactor(32)
		rating := e.CalculateNewRating(1613, 2.0, 2.5)
		if rating != 1629 {
			t.Errorf("expected rating to be 1629, actual %d\n", rating)
		}
	})

	t.Run("validate rating stays the same if score is the same as expected...", func(t *testing.T) {
		e.SetKFactor(32)
		rating := e.CalculateNewRating(1613, 2.88, 2.88)
		if rating != 1613 {
			t.Errorf("expected rating to be 1613, actual %d\n", rating)
		}
	})
}

func Test_Simulated(t *testing.T) {
	elo := ELOCalculator{}
	t.Run("validate rating drops if score is lower than expected...", func(t *testing.T) {
		playerRating := 1613
		elo.SetKFactorFromRating(playerRating)
		expectedScore := 0.0
		s, _ := elo.ExpectedScores(playerRating, 1609)
		expectedScore = s
		s, _ = elo.ExpectedScores(playerRating, 1477)
		expectedScore += s
		s, _ = elo.ExpectedScores(playerRating, 1388)
		expectedScore += s
		s, _ = elo.ExpectedScores(playerRating, 1586)
		expectedScore += s
		s, _ = elo.ExpectedScores(playerRating, 1720)
		expectedScore += s
		actualScore := elo.Score(2, 1, 2)
		rating := elo.CalculateNewRating(playerRating, expectedScore, actualScore)
		if rating != 1601 {
			t.Errorf("expected rating to be 1601, actual %d\n", rating)
		}
	})
}
