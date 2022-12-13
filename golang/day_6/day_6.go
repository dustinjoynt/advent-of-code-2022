package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"strings"
)

func main() {

	input, err := input.GetInput("6")
	if err != nil {
		log.Fatalf("failed to get input for day six: %v", err)
	}

	marker := getPacketMarker(input)

	fmt.Printf("start of packet marker is found at %v", marker)
}

func getPacketMarker(input string) int {

	strSlice := strings.Split(input, "")

	for i := 0; i < len(strSlice); i++ {
		markerRange := strSlice[i : 4+i]
		if checkForMarker(markerRange) {
			return i + 4
		}
	}

	panic("no marker found")
}

func checkForMarker(markerRange []string) bool {
	check := make(map[string]bool)
	for _, v := range markerRange {
		check[v] = true
	}
	return len(check) == 4
}
