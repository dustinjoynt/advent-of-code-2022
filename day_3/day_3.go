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
	input, err := input.GetInput("3")
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
	rs := getRuckSackChunks(input)
	for _, v := range rs {
		if len(v) > 0 {
			sum += getPriority(v)
		}
	}
	return sum
}

func getRuckSackChunks(input string) [][]string {
	var chunks [][]string
	rs := getRuckSacks(input)

	for {
		if len(rs) < 3 {
			break
		}
		chunks = append(chunks, rs[0:3])
		rs = rs[3:]
	}
	return chunks
}

func getRuckSacks(input string) []string {
	re := regexp.MustCompile("\n")
	return re.Split(input, -1)
}

func getPriority(contents []string) int {
	match := getDuplicate(contents)
	return priorityValues[match]
}

func getDuplicate(contents []string) string {
	var items [][]string
	for _, v := range contents {
		items = append(items, strings.Split(v, ""))
	}

	for _, v1 := range items[0] {
		for _, v2 := range items[1] {
			for _, v3 := range items[2] {
				if v1 == v2 && v1 == v3 {
					return v1
				}
			}

		}
	}
	panic("no match found in rucksacks")
}
