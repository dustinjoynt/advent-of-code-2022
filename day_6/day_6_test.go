package main

import "testing"

func TestGetPacketMarker(t *testing.T) {

	t.Run("sample one", func(t *testing.T) {
		input := "bvwbjplbgvbhsrlpgdmjqwftvncz"

		marker := getPacketMarker(input)

		want := 5
		if marker != want {
			t.Fatalf("failed to get marker. want: %v, got: %v", want, marker)
		}
	})

	t.Run("sample two", func(t *testing.T) {
		input := "nppdvjthqldpwncqszvftbrmjlhg"

		marker := getPacketMarker(input)

		want := 6
		if marker != want {
			t.Fatalf("failed to get marker. want: %v, got: %v", want, marker)
		}
	})

}
