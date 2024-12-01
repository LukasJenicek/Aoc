package helpers

import (
	"fmt"
	"strconv"
	"unicode"
)

func LoadNumbers(line string) ([]int, error) {
	numbers := make([]int, 0)
	number := ""

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

			numbers = append(numbers, int(n))
			number = ""
			continue
		}
	}

	if number != "" {
		n, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parsing number: %s", err)
		}

		numbers = append(numbers, int(n))
	}

	return numbers, nil
}
