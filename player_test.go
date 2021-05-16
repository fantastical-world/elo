package elo

import "testing"

func TestCalculateNewRating(t *testing.T) {
	p := NewPlayer(1613, 32)
	p.CalculateNewRating(2.88, 2.5)
	if p.Rating() != 1601 {
		t.Errorf("expected rating to be 1601, actual %d\n", p.Rating())
	}
}

func TestCalculateNewRatingWithUCSF(t *testing.T) {
	p := NewPlayerKFactorFromRating(1613)
	p.CalculateNewRating(2.88, 2.5)
	if p.Rating() != 1601 {
		t.Errorf("expected rating to be 1601, actual %d\n", p.Rating())
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
			t.Errorf("expected rating to be 1601, actual %d\n", player.Rating())
		}
	})
}
