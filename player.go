package elo

import (
	"math"
)

//Player represents a player, their rating, and play history.
type Player struct {
	rating  int
	kfactor float64
}

func NewPlayer(rating int, kfactor float64) *Player {
	p := &Player{rating: rating, kfactor: kfactor}
	return p
}

func NewPlayerUSCF(rating int) *Player {
	p := &Player{rating: rating, kfactor: USCFKFactorFromRating(rating)}
	return p
}

func (p *Player) Rating() int {
	return p.rating
}

func (p *Player) CalculateNewRating(expectedScore, actualScore float64) {
	value := math.Round((p.kfactor * (actualScore - expectedScore)))
	p.rating = p.rating + int(value)
}
