package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type Tree struct {
}

func main() {

	input, err := input.GetInput("8")
	if err != nil {
		log.Fatalf("failed to get input for day eight: %v", err)
	}

	trees := buildMap(input)
	fmt.Println(trees)
}

func buildMap(input string) [][]int {

	rowRe := regexp.MustCompile(`\n`)
	rows := rowRe.Split(input, -1)
	rows = rows[:len(rows)-1]

	trees := make([][]int, len(rows))

	for i, row := range rows {
		columnRe := regexp.MustCompile(``)
		columns := columnRe.Split(row, -1)

		intColumn := []int{}
		for _, column := range columns {
			i, _ := strconv.Atoi(column)
			intColumn = append(intColumn, i)
		}
		// trees[i] = make([]int, len(columns[i]))
		trees[i] = intColumn
	}

	return trees
}
