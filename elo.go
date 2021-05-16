/*
Package elo provides ELOCalculator which can be used to calculate skill levels.
*/
package elo

import (
	"math"
)

//ELOCalculator can be used to calculate scores, and ratings.
type ELOCalculator struct {
	kfactor float64
}

//New creates a new ELOCalculator with the provided kfactor. The kfactor is used when calculating new ratings.
func New(kfactor float64) ELOCalculator {
	return ELOCalculator{kfactor: kfactor}
}

//SetKFactor sets the kfactor used when calculating a new rating.
func (e *ELOCalculator) SetKFactor(k float64) {
	e.kfactor = k
}

//Score calculates an actual score based on wins, draws, and losses.
func (e *ELOCalculator) Score(wins, draws, losses int) float64 {
	//technically we don't need to include losses at all since it will always be 0.
	return (float64(wins) * 1.0) + (float64(draws) * 0.5) + (float64(losses) * 0.0)
}

//ExpectedScores will calculate the expected scores of two players based on their ratings.
func (e *ELOCalculator) ExpectedScores(ratingA, ratingB int) (float64, float64) {
	scoreA := 1 / (1 + math.Pow(10, ((float64(ratingB)-float64(ratingA))/400)))
	scoreB := 1 / (1 + math.Pow(10, ((float64(ratingA)-float64(ratingB))/400)))

	return math.Round(scoreA*100) / 100, math.Round(scoreB*100) / 100
}

//CalculateNewRating will calculate a new rating based on the current rating with the provided scores.
func (e *ELOCalculator) CalculateNewRating(rating int, expectedScore, actualScore float64) int {
	value := math.Round((e.kfactor * (actualScore - expectedScore)))
	return rating + int(value)
}

//SetKFactorFromRating sets kfactor from a rating.
func (e *ELOCalculator) SetKFactorFromRating(rating int) {
	if rating < 2100 {
		e.kfactor = 32
		return
	}
	if rating > 2400 {
		e.kfactor = 16
		return
	}
	e.kfactor = 24
}

//SetKFactorFromGamesPlayed sets kfactor from the amount of games played previously, and today.
func (e *ELOCalculator) SetKFactorFromGamesPlayed(pp, pt int) {
	e.kfactor = 800 / float64(pp+pt)
}
