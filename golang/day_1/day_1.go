package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {

	input, err := input.GetInput("1")
	if err != nil {
		log.Fatalf("failed to get input for day one: %v", err)
	}

	elfLoads := getElfLoads(input)

	topThree := getTopThreeLoads(elfLoads)

	topThreeSum := 0
	for _, v := range topThree {
		topThreeSum += v
	}

	fmt.Printf("total calories carried by the elf with the most calories is: %v \n", topThree[0])
	fmt.Printf("total calories carried by the top three elves is: %v", topThreeSum)
}

func getElfLoads(input string) []int {
	re := regexp.MustCompile("\n\n")
	elfSplit := re.Split(input, -1)

	elfLoads := []int{}
	for _, v := range elfSplit {
		re := regexp.MustCompile("\n")
		loadSplit := re.Split(v, -1)

		loadSum := 0
		for _, v := range loadSplit {
			intV, _ := strconv.Atoi(v)
			loadSum += intV
		}

		elfLoads = append(elfLoads, loadSum)
		loadSum = 0
	}

	return elfLoads
}

func getTopThreeLoads(loads []int) []int {

	topThree := []int{}
	for _, v := range loads {

		if len(topThree) < 3 {
			topThree = append(topThree, v)
		} else {
			if v > topThree[0] {
				topThree[2] = topThree[1]
				topThree[1] = topThree[0]
				topThree[0] = v
			} else if v > topThree[1] {
				topThree[2] = topThree[1]
				topThree[1] = v
			} else if v > topThree[2] {
				topThree[2] = v
			}
		}
	}

	return topThree
}
