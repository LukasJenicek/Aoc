package day01

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/LukasJenicek/aoc/libs/helpers"
)

//go:embed input01.txt
var input01 string

func RunFirst() (string, error) {
	for _, line := range strings.Split(input01, "\n") {
		numbers, err := helpers.LoadNumbers(line)
		if err != nil {
			return "", fmt.Errorf("load numbers: %w", err)
		}

		fmt.Println(numbers)
	}

	return "", nil
}
