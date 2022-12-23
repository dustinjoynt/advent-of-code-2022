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

	count := countVisibleTrees(input)

	fmt.Println(count)
}

func countVisibleTrees(input string) int {
	tM := NewMap(input)
	count := 0
	for i, row := range tM.treeMap {
		for j := range row {
			if tM.isVisible(i, j) {
				count++
			}
		}
	}
	return count
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

func (m *TreeMap) isVisible(rowIndex int, columnIndex int) bool {
	return m.isVisibleFromLeft(rowIndex, columnIndex) ||
		m.isVisibleFromRight(rowIndex, columnIndex) ||
		m.isVisibleFromTop(rowIndex, columnIndex) ||
		m.isVisibleFromBottom(rowIndex, columnIndex)
}

func (m *TreeMap) isVisibleFromLeft(rowIndex int, columnIndex int) bool {
	for i := rowIndex - 1; i >= 0; i-- {

		// fmt.Printf("%v >= %v \n", m.treeMap[columnIndex][i], m.treeMap[columnIndex][rowIndex])

		if m.treeMap[columnIndex][i] >= m.treeMap[columnIndex][rowIndex] {
			return false
		}
	}
	return true
}

func (m *TreeMap) isVisibleFromRight(rowIndex int, columnIndex int) bool {
	for i := rowIndex + 1; i <= len(m.treeMap[columnIndex])-1; i++ {
		if m.treeMap[columnIndex][i] >= m.treeMap[columnIndex][rowIndex] {
			return false
		}
	}
	return true
}

func (m *TreeMap) isVisibleFromTop(rowIndex int, columnIndex int) bool {
	for i := columnIndex - 1; i >= 0; i-- {
		if m.treeMap[i][rowIndex] >= m.treeMap[columnIndex][rowIndex] {
			return false
		}
	}
	return true
}

func (m *TreeMap) isVisibleFromBottom(rowIndex int, columnIndex int) bool {
	for i := columnIndex + 1; i <= len(m.treeMap)-1; i++ {
		if m.treeMap[i][rowIndex] >= m.treeMap[columnIndex][rowIndex] {
			return false
		}
	}
	return true

}
