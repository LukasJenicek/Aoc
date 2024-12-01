package day01

import (
	_ "embed"
	"fmt"

	"github.com/LukasJenicek/aoc/libs/helpers"
)

//go:embed input01.txt
var input01 string

func RunFirst() (string, error) {
	numbers, err := helpers.LoadNumbers(input01)
	if err != nil {
		return "", fmt.Errorf("load numbers: %w", err)
	}

	fmt.Println(numbers)

	return "", nil
}
