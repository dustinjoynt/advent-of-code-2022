package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var priorityValues map[string]int

func main() {
	input, err := input.GetInput("https://adventofcode.com/2022/day/3/input")
	if err != nil {
		log.Fatalf("failed to get input for day three: %v", err)
	}

	setPriorityValues()
	prioritySum := getPrioritySum(input)
	fmt.Printf("sum of priority: %v", prioritySum)
}

func setPriorityValues() {

	m := make(map[string]int, 52)

	value := 1
	for i := 97; i <= 122; i++ {
		m[string(i)] = value
		value++
	}
	for i := 65; i <= 90; i++ {
		m[string(i)] = value
		value++
	}

	priorityValues = m
}

func getPrioritySum(input string) int {
	sum := 0
	rs := getRuckSacks(input)
	for _, v := range rs {
		if len(v) > 0 {
			sum += getPriority(v)
		}
	}
	return sum
}

func getRuckSacks(input string) []string {
	re := regexp.MustCompile("\n")
	return re.Split(input, -1)
}

func getPriority(contents string) int {

	match := getDuplicate(contents)

	return priorityValues[match]
}

func getDuplicate(contents string) string {

	items := strings.Split(contents, "")

	l := len(items)

	ruckSack1 := items[0:(l / 2)]
	ruckSack2 := items[(l / 2):l]

	for _, v1 := range ruckSack1 {
		for _, v2 := range ruckSack2 {
			if v1 == v2 {
				return v1
			}
		}
	}
	panic("no match found in rucksacks")
}
