package elo

import (
	"math"
)

const (
	//Scores awarded for win, lose, or draw.
	win  = 1.0
	draw = 0.5
	loss = 0.0
)

//ELOCalculator will calculate player scores, and provides kfactors to be used by players.
type ELOCalculator struct {
}

//Score calculates a score based on wins, draws, and losses.
func (e *ELOCalculator) Score(wins, draws, losses int) float64 {
	score := 0.00

	if wins > 0 {
		score += (float64(wins) * win)
	}

	if draws > 0 {
		score += (float64(draws) * draw)
	}

	if losses > 0 {
		score += (float64(losses) * loss)
	}

	return score
}

//ExpectedScores will calculate the expected scores of two players based on their ratings.
func (e *ELOCalculator) ExpectedScores(playerA, playerB *Player) (float64, float64) {
	scoreA := 1 / (1 + math.Pow(10, ((playerB.Rating()-playerA.Rating())/400)))
	scoreB := 1 / (1 + math.Pow(10, ((playerA.Rating()-playerB.Rating())/400)))

	return math.Round(scoreA*100) / 100, math.Round(scoreB*100) / 100
}

//Player represents a player, their rating, and play history.
type Player struct {
	rating  float64
	kfactor float64
}

func NewPlayer(rating, kfactor float64) *Player {
	p := &Player{rating: rating, kfactor: kfactor}
	return p
}

func (p *Player) Rating() float64 {
	return p.rating
}

func (p *Player) CalculateNewRating(expectedScore, actualScore float64) {
	p.rating = p.rating + math.Round((p.kfactor * (actualScore - expectedScore)))
}
