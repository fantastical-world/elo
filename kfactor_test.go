package elo

import "testing"

func TestUSCFKFactorFromRating(t *testing.T) {
	kfactor := USCFKFactorFromRating(2000)
	t.Run("validate kfactor for rating below 2100...", func(t *testing.T) {
		if kfactor != 32 {
			t.Errorf("expected kfactor to be 32, actual %f\n", kfactor)
		}
	})
	kfactor = USCFKFactorFromRating(2500)
	t.Run("validate kfactor for rating above 2400...", func(t *testing.T) {
		if kfactor != 16 {
			t.Errorf("expected kfactor to be 16, actual %f\n", kfactor)
		}
	})
	kfactor = USCFKFactorFromRating(2300)
	t.Run("validate kfactor when rating is between 2100  and  2400 (inclusive)...", func(t *testing.T) {
		if kfactor != 24 {
			t.Errorf("expected kfactor to be 24, actual %f\n", kfactor)
		}
	})
}

func TestUSCFKFactorGamesPlayed(t *testing.T) {
	kfactor := USCFKFactorFromGamesPlayed(30, 5)
	t.Run("validate kfactor from rating based on 30 games previously played, and 5 played today...", func(t *testing.T) {
		if kfactor != 22.857142857142858 {
			t.Errorf("expected kfactor to be 22.857142857142858, actual %f\n", kfactor)
		}
	})
}
