/*
Package elo provides a Calculator which can be used to calculate ratings based on the ELO system.

This package does not enforce any constraints on the type of games, players, events, or rules to be used.
Any constraints should be implemented elsewhere. This package simply calculates a new rating based
on the values provided for a "player's" current rating, expected score, and actual score. The only
configurable value for fine tuning the calculated rating is the kfactor. The kfactor can be set to
any value you desire, the logic for determining a kfactor is left to the user of this package.
Although two common kfactor calculators are provided for convience, but they are very simple in their
implementation.

Convience methods are provided for calculating an actual score, and an expected score. The actual score is
based on the number of wins, losses, or draws over all the "games" played in the current "event". To calculate
an "expected" score two ratings are needed, this typically represents a player vs. player score. The calculator will return
the "expected" score for both players. This allows for recording the outcome of a single game, or tracking multiple
games and players in an event.
*/
package elo

import (
	"math"
)

//Calculator is used to calculate scores, and ratings.
type Calculator struct {
	kfactor float64
}

//New creates a new Calculator with the provided kfactor. The kfactor is used when calculating new ratings.
func New(kfactor float64) Calculator {
	return Calculator{kfactor: kfactor}
}

//SetKFactor sets the kfactor used when calculating a new rating.
func (e *Calculator) SetKFactor(k float64) {
	e.kfactor = k
}

//Score calculates an actual score based on wins, draws, and losses.
func (e *Calculator) Score(wins, draws, losses int) float64 {
	//technically we don't need to include losses at all since it will always be 0.
	return (float64(wins) * 1.0) + (float64(draws) * 0.5) + (float64(losses) * 0.0)
}

//ExpectedScores will calculate the expected scores of two players based on their ratings.
func (e *Calculator) ExpectedScores(ratingA, ratingB int) (float64, float64) {
	scoreA := 1 / (1 + math.Pow(10, ((float64(ratingB)-float64(ratingA))/400)))
	scoreB := 1 / (1 + math.Pow(10, ((float64(ratingA)-float64(ratingB))/400)))

	return math.Round(scoreA*100) / 100, math.Round(scoreB*100) / 100
}

//NewRating will calculate a new rating based on the current rating, expected score, and actual score.
func (e *Calculator) NewRating(currentRating int, expectedScore, actualScore float64) int {
	value := math.Round((e.kfactor * (actualScore - expectedScore)))
	return currentRating + int(value)
}

//SetKFactorFromRating sets kfactor from a rating.
func (e *Calculator) SetKFactorFromRating(rating int) {
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

//SetKFactorFromGamesPlayed sets kfactor from the amount of games played in the past, and the number played at the current "event".
func (e *Calculator) SetKFactorFromGamesPlayed(previous, current int) {
	e.kfactor = 800 / float64(previous+current)
}
