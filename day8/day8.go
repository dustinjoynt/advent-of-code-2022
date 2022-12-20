package main

import (
	"advent-of-code-2022/input"
	"log"
)

func main() {

	input, err := input.GetInput("8")
	if err != nil {
		log.Fatalf("failed to get input for day eight: %v", err)
	}

}
