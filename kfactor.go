package elo

//KFactorFromRating provides a kfactor from a player's rating.
func KFactorFromRating(rating int) float64 {
	if rating < 2100 {
		return 32
	}
	if rating > 2400 {
		return 16
	}
	return 24
}

//KFactorFromGamesPlayed provides a kfactor from the amount of games played.
func KFactorFromGamesPlayed(pp, pt int) float64 {
	return 800 / float64(pp+pt)
}
