package elo

import (
	"math"
)

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
