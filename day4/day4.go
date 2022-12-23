package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	input, err := input.GetInput("4")
	if err != nil {
		log.Fatalf("failed to get input for day four: %v", err)
	}

	count := getCoveredAssignmentPairs(input)
	fmt.Printf("the count of fully covered assignment pairs is %v", count)
}

func getCoveredAssignmentPairs(input string) int {

	count := 0
	assignmentPairs := getAssignmentPairs(input)
	for _, v := range assignmentPairs {
		if len(v) > 0 {
			if isAssignmentCovered(v) {
				count++
			}
		}
	}
	return count
}

func getAssignmentPairs(input string) []string {
	re := regexp.MustCompile("\n")
	return re.Split(input, -1)
}

func isAssignmentCovered(assignmentPair string) bool {

	pairs := splitAssignmentPair(assignmentPair)
	return containsDuplicate(pairs)

}

func splitAssignmentPair(assignmentPair string) []string {
	re := regexp.MustCompile(",")
	return re.Split(assignmentPair, -1)
}

func containsDuplicate(pairs []string) bool {

	r1 := getRange(pairs[0])
	r2 := getRange(pairs[1])

	return valuesOverlap(r1, r2)
}

func getRange(assignment string) []int {

	re := regexp.MustCompile("-")
	r := re.Split(assignment, -1)

	l, _ := strconv.Atoi(r[0])
	m, _ := strconv.Atoi(r[1])

	iR := []int{}
	for i := l; i <= m; i++ {
		iR = append(iR, i)
	}
	return iR
}

func valuesOverlap(a []int, b []int) bool {

	check := make(map[int]bool)
	for _, v := range a {
		check[v] = true
	}
	for _, v := range b {
		if check[v] {
			return true
		}
	}
	return false
}
