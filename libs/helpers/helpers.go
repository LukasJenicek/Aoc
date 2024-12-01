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
