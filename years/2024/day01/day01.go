package day202401

import (
	_ "embed"
	"fmt"
	"math"
	"sort"

	"github.com/LukasJenicek/aoc/internal/helpers"
)

type FirstDay struct{}

func (*FirstDay) RunFirstPart(input string) (string, error) {
	numbers, err := helpers.LoadNumbers(input)
	if err != nil {
		return "", fmt.Errorf("load numbers: %w", err)
	}

	firstCol := make([]int, len(numbers))
	secondCol := make([]int, len(numbers))
	for i, rows := range numbers {
		firstCol[i] = rows[0]
		secondCol[i] = rows[1]
	}

	sort.Ints(firstCol)
	sort.Ints(secondCol)

	distance := 0
	for i, smallestNumber := range firstCol {
		distance += int(math.Abs(float64(smallestNumber - secondCol[i])))
	}

	return fmt.Sprintf("%v", distance), nil
}

func (*FirstDay) RunSecondPart(input string) (string, error) {
	numbers, err := helpers.LoadNumbers(input)
	if err != nil {
		return "", fmt.Errorf("load numbers: %w", err)
	}

	firstCol := make([]int, len(numbers))
	secondCol := make([]int, len(numbers))
	for i, rows := range numbers {
		firstCol[i] = rows[0]
		secondCol[i] = rows[1]
	}

	sort.Ints(firstCol)
	sort.Ints(secondCol)

	distance := 0
	score := 0
	for _, firstColNumber := range firstCol {
		score = 0

		for _, secondColNumber := range secondCol {
			if secondColNumber > firstColNumber {
				break
			}

			if firstColNumber == secondColNumber {
				score++
			}
		}

		distance += firstColNumber * score
	}

	return fmt.Sprintf("%v", distance), nil
}
