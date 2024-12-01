package day01

import (
	_ "embed"
	"fmt"
	"math"
	"sort"

	"github.com/LukasJenicek/aoc/libs/helpers"
)

//go:embed input01.txt
var input01 string

func CalculateFirstAssignment() (string, error) {
	numbers, err := helpers.LoadNumbers(input01)
	if err != nil {
		return "", fmt.Errorf("load numbers: %w", err)
	}

	firstCol := make([]int, len(numbers))
	secondCol := make([]int, len(numbers))
	for i, rows := range numbers {
		firstCol[i] = rows[0]
		secondCol[i] = rows[1]
	}

	sort.Slice(firstCol, func(i, j int) bool {
		return firstCol[i] < firstCol[j]
	})

	sort.Slice(secondCol, func(i, j int) bool {
		return secondCol[i] < secondCol[j]
	})

	distance := 0
	for i, smallestNumber := range firstCol {
		distance += int(math.Abs(float64(smallestNumber - secondCol[i])))
	}

	return fmt.Sprintf("%v", distance), nil
}
