package main

import "testing"

func TestGetRoundScorePart1(t *testing.T) {

	t.Run("Rock Beats Scissors", func(t *testing.T) {

		round := "C X"
		score := getRoundScorePart1(round)

		want := 7
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

	t.Run("Paper Losses Scissors", func(t *testing.T) {

		round := "C Y"
		score := getRoundScorePart1(round)

		want := 2
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

	t.Run("Tie with Scissors", func(t *testing.T) {

		round := "C Z"
		score := getRoundScorePart1(round)

		want := 6
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

}

func TestGetRoundScorePart2(t *testing.T) {

	t.Run("Lose", func(t *testing.T) {

		round := "C X"
		score := getRoundScorePart2(round)

		want := 2
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

	t.Run("Tie", func(t *testing.T) {

		round := "C Y"
		score := getRoundScorePart2(round)

		want := 6
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

	t.Run("Win", func(t *testing.T) {

		round := "C Z"
		score := getRoundScorePart2(round)

		want := 7
		if score != want {
			t.Fatalf("incorrect score. want: %v, got %v", want, score)
		}
	})

}
