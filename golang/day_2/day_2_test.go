package main

import "testing"

func TestGetRoundScore(t *testing.T) {

	t.Run("Rock Beats Scissors", func(t *testing.T) {

		round := "C X"
		score := getRoundScore(round)

		want := 7
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

	t.Run("Paper Losses Scissors", func(t *testing.T) {

		round := "C Y"
		score := getRoundScore(round)

		want := 2
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

	t.Run("Tie with Scissors", func(t *testing.T) {

		round := "C Z"
		score := getRoundScore(round)

		want := 6
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

}
