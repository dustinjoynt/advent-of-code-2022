package main

import (
	"advent-of-code-2022/input"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type TreeMap struct {
	treeMap [][]int
}

func main() {

	input, err := input.GetInput("8")
	if err != nil {
		log.Fatalf("failed to get input for day eight: %v", err)
	}

	count := getBestScenicScore(input)

	fmt.Printf("the best scenic score is: %v", count)
}

func getBestScenicScore(input string) int {
	tM := NewMap(input)
	bestScenicScore := 0
	for i, row := range tM.treeMap {
		for j := range row {
			scenicScore := tM.getScenicScore(i, j)
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}
		}
	}
	return bestScenicScore
}

func NewMap(input string) *TreeMap {

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
		trees[i] = intColumn
	}

	return &TreeMap{
		treeMap: trees,
	}
}

func (m *TreeMap) getScenicScore(rowIndex int, columnIndex int) int {
	return m.countVisibleFromLeft(rowIndex, columnIndex) *
		m.countVisibleFromRight(rowIndex, columnIndex) *
		m.countVisibleFromTop(rowIndex, columnIndex) *
		m.countVisibleFromBottom(rowIndex, columnIndex)
}

func (m *TreeMap) countVisibleFromLeft(rowIndex int, columnIndex int) int {
	count := 0
	for i := rowIndex - 1; i >= 0; i-- {
		count++
		if m.treeMap[columnIndex][i] >= m.treeMap[columnIndex][rowIndex] {
			return count
		}
	}
	return count
}

func (m *TreeMap) countVisibleFromRight(rowIndex int, columnIndex int) int {
	count := 0
	for i := rowIndex + 1; i < len(m.treeMap[columnIndex]); i++ {
		count++
		if m.treeMap[columnIndex][i] >= m.treeMap[columnIndex][rowIndex] {
			return count
		}
	}
	return count
}

func (m *TreeMap) countVisibleFromTop(rowIndex int, columnIndex int) int {
	count := 0
	for i := columnIndex - 1; i >= 0; i-- {
		count++
		if m.treeMap[i][rowIndex] >= m.treeMap[columnIndex][rowIndex] {
			return count
		}
	}
	return count
}

func (m *TreeMap) countVisibleFromBottom(rowIndex int, columnIndex int) int {
	count := 0
	for i := columnIndex + 1; i <= len(m.treeMap)-1; i++ {
		count++
		if m.treeMap[i][rowIndex] >= m.treeMap[columnIndex][rowIndex] {
			return count
		}
	}
	return count
}
