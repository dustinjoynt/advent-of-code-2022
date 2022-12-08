package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
)

func main() {

	input, err := input.GetInput("7")
	if err != nil {
		log.Fatalf("failed to get input for day seven: %v", err)
	}

	fmt.Println(input)
}
