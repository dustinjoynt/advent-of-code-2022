package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
)

func main() {

	input, err := input.GetInput("https://adventofcode.com/2022/day/2/input")
	if err != nil {
		log.Fatalf("failed to get input for day one: %v", err)
	}

	re := regexp.MustCompile("\n")
	rounds := re.Split(input, -1)

	totalScore := 0
	for _, v := range rounds {
		if len(v) > 0 {
			totalScore += getRoundScore(v)
		}
	}

	fmt.Printf("total score according to strategy guide: %v", totalScore)
}

func getRoundScore(round string) int {

	scores := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	re := regexp.MustCompile(" ")
	choice := re.Split(round, -1)

	if isTie(choice[0], choice[1]) {
		return 3 + scores[choice[1]]
	} else if youWin(choice[0], choice[1]) {
		return 6 + scores[choice[1]]
	} else {
		return scores[choice[1]]
	}
}

func isTie(elfChoice string, yourChoice string) bool {
	return elfChoice == "A" && yourChoice == "X" ||
		elfChoice == "B" && yourChoice == "Y" ||
		elfChoice == "C" && yourChoice == "Z"
}

func youWin(elfChoice string, yourChoice string) bool {
	return elfChoice == "C" && yourChoice == "X" ||
		elfChoice == "A" && yourChoice == "Y" ||
		elfChoice == "B" && yourChoice == "Z"
}
