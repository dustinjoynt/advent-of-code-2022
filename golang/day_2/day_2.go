package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
)

var scores = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {

	input, err := input.GetInput("https://adventofcode.com/2022/day/2/input")
	if err != nil {
		log.Fatalf("failed to get input for day two: %v", err)
	}

	re := regexp.MustCompile("\n")
	rounds := re.Split(input, -1)

	totalScore := 0
	revisedScore := 0
	for _, v := range rounds {
		if len(v) > 0 {
			totalScore += getRoundScorePart1(v)
			revisedScore += getRoundScorePart2(v)
		}
	}

	fmt.Printf("total score according to strategy guide: %v \n", totalScore)
	fmt.Printf("total score according to revised strategy guide: %v", revisedScore)
}

func getRoundScorePart1(round string) int {

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

func getRoundScorePart2(round string) int {

	re := regexp.MustCompile(" ")
	choice := re.Split(round, -1)

	if choice[1] == "X" {
		return scores[getLosingChoice(choice[0])]
	} else if choice[1] == "Y" {
		return 3 + scores[getTyingChoice(choice[0])]
	} else {
		return 6 + scores[getWinningChoice(choice[0])]
	}
}

func getLosingChoice(elfChoice string) string {
	if elfChoice == "A" {
		return "Z"
	} else if elfChoice == "B" {
		return "X"
	} else {
		return "Y"
	}
}

func getTyingChoice(elfChoice string) string {
	if elfChoice == "A" {
		return "X"
	} else if elfChoice == "B" {
		return "Y"
	} else {
		return "Z"
	}
}

func getWinningChoice(elfChoice string) string {
	if elfChoice == "A" {
		return "Y"
	} else if elfChoice == "B" {
		return "Z"
	} else {
		return "X"
	}
}
