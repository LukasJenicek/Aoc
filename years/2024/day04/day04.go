package day202404

import (
	"fmt"
	"strings"
)

type FourthDay struct{}

func (*FourthDay) RunFirstPart(input string) (string, error) {
	grid := buildGrid(input)

	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // down
		{1, 1},  // diagonal down-right
		{1, -1}, // diagonal down-left
	}

	occurrences := 0
	word := "XMAS"
	reversedWord := "SAMX"

	rowsCount := len(grid)
	colsCount := len(grid[0])

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
	grid := buildGrid(input)

	directions := [][]int{
		{1, 1}, // diagonal down-right
	}

	occurrences := 0
	word := "MAS"
	reversedWord := "SAM"

	rowsCount := len(grid)
	colsCount := len(grid[0])

	for i, row := range grid {
		for j, _ := range row {
			for _, dc := range directions {
				if searchWord(grid, rowsCount, colsCount, i, j, dc[0], dc[1], word) {
					if j+2 < colsCount && i+2 < rowsCount {
						if searchWord(grid, rowsCount, colsCount, i, j+2, 1, -1, word) ||
							searchWord(grid, rowsCount, colsCount, i, j+2, 1, -1, reversedWord) {
							occurrences++
							continue
						}
					}
				}

				if searchWord(grid, rowsCount, colsCount, i, j, dc[0], dc[1], reversedWord) {
					if j+2 < colsCount && i+2 < rowsCount {
						if searchWord(grid, rowsCount, colsCount, i, j+2, 1, -1, reversedWord) ||
							searchWord(grid, rowsCount, colsCount, i, j+2, 1, -1, word) {
							occurrences++
						}
					}
				}
			}
		}
	}

	return fmt.Sprintf("%v", occurrences), nil
}

func searchWord(grid [][]string, rowsCount, colsCount, curRow, curCol, dirRow, dirCol int, word string) bool {
	found := true

	for i, w := range word {
		r := curRow + i*dirRow
		c := curCol + i*dirCol

		if r < 0 || r >= rowsCount || c < 0 || c >= colsCount || grid[r][c] != string(w) {
			found = false
			break
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
