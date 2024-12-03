package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func LoadNumbers(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	linesCount := len(lines)
	numbers := make([][]int, linesCount)

	number := ""
	for i, line := range lines {
		for _, char := range line {
			if number == "" && unicode.IsSpace(char) {
				continue
			}

			if unicode.IsDigit(char) {
				number += string(char)
				continue
			}

			if number != "" {
				n, err := strconv.ParseInt(number, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("parsing number: %s", err)
				}

				numbers[i] = append(numbers[i], int(n))
				number = ""
				continue
			}
		}

		if number != "" {
			n, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("parsing number: %s", err)
			}

			numbers[i] = append(numbers[i], int(n))
			number = ""
		}
	}

	return numbers, nil
}

func AbsoluteDiff(x, y int) int {
	if x < y {
		return y - x
	}

	return x - y
}

func RemoveAtIndex(slice []int, index int) []int {
	// no bound check just dont call a oob index
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:index]...)
	newSlice = append(newSlice, slice[index+1:]...)

	return newSlice
}
