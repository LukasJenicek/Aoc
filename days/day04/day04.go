package day04

import (
	"fmt"
	"strings"
)

var directions = [][]int{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // down
	{-1, 0},  // up
	{1, 1},   // diagonal down-right
	{-1, -1}, // diagonal up-left
	{1, -1},  // diagonal down-left
	{-1, 1},  // diagonal up-right
}

var foundWordCells = [][]int{}

type FourthDay struct{}

func (*FourthDay) RunFirstPart(input string) (string, error) {
	grid := buildGrid(input)

	occurrences := 0
	word := "XMAS"
	reversedWord := "SAMX"

	rowsCount := len(grid)
	colsCount := len(grid[0])

	for i := 0; i < rowsCount; i++ {
		for j := 0; j < colsCount; j++ {
			foundWordCells = append(foundWordCells, make([]int, colsCount))
		}
	}

	fmt.Println("rows count", rowsCount)
	fmt.Println("cols count", colsCount)

	for i, row := range grid {
		for j, _ := range row {
			for _, dc := range directions {
				if searchWord(grid, rowsCount, colsCount, i, j, dc[0], dc[1], word) {
					occurrences++
				}

				if searchWord(grid, rowsCount, colsCount, i, j, dc[0], dc[1], reversedWord) {
					occurrences++
				}
			}
		}
	}

	return fmt.Sprintf("%v", occurrences), nil
}

func (*FourthDay) RunSecondPart(input string) (string, error) {
	return "", nil
}

func searchWord(grid [][]string, rowsCount, colsCount, curRow, curCol, dirRow, dirCol int, word string) bool {
	found := true
	foundCells := [][]int{}

	for i, w := range word {
		r := curRow + i*dirRow
		c := curCol + i*dirCol

		foundCells = append(foundCells, []int{r, c})

		if r < 0 || r >= rowsCount || c < 0 || c >= colsCount || grid[r][c] != string(w) {
			found = false
			break
		}
	}

	if found {
		if foundWordCells[foundCells[0][0]][foundCells[0][1]] == 1 &&
			foundWordCells[foundCells[1][0]][foundCells[1][1]] == 1 &&
			foundWordCells[foundCells[2][0]][foundCells[2][1]] == 1 &&
			foundWordCells[foundCells[3][0]][foundCells[3][1]] == 1 {
			return false
		}

		for _, cell := range foundCells {
			foundWordCells[cell[0]][cell[1]] = 1
		}
	}

	return found
}

func buildGrid(input string) [][]string {
	var grid [][]string

	for _, rows := range strings.Split(input, "\n") {
		var row []string
		for _, char := range rows {
			row = append(row, string(char))
		}
		grid = append(grid, row)
	}

	return grid
}
