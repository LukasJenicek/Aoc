package day202403

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

type ThirdDay struct{}

func (*ThirdDay) RunFirstPart(input string) (string, error) {
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	multiplications := re.FindAllStringSubmatch(input, -1)

	var result int64
	for _, multiplication := range multiplications {
		mul := multiplication[0]

		parsingSecondNumber := false
		firstNumber := ""
		secondNumber := ""
		for _, m := range mul {
			if unicode.IsDigit(m) {
				if parsingSecondNumber {
					secondNumber += string(m)
					continue
				}

				firstNumber += string(m)
			}

			if m == ',' {
				parsingSecondNumber = true
			}
		}

		fNumber, err := strconv.ParseInt(firstNumber, 10, 64)
		if err != nil {
			return "", fmt.Errorf("could not parse first number: %w", err)
		}

		sNumber, err := strconv.ParseInt(secondNumber, 10, 64)
		if err != nil {
			return "", fmt.Errorf("could not parse second number: %w", err)
		}

		result += fNumber * sNumber
	}

	return fmt.Sprintf("%v", result), nil
}

func (*ThirdDay) RunSecondPart(input string) (string, error) {
	re := regexp.MustCompile(`do\(\)|don't\(\)|mul\([0-9]+,[0-9]+\)`)
	multiplications := re.FindAllStringSubmatch(input, -1)

	var result int64
	do := true
	for _, multiplication := range multiplications {
		mul := multiplication[0]

		if mul == "don't()" {
			do = false
			continue
		}

		if mul == "do()" {
			do = true
			continue
		}

		if !do {
			continue
		}

		parsingSecondNumber := false
		firstNumber := ""
		secondNumber := ""
		for _, m := range mul {
			if unicode.IsDigit(m) {
				if parsingSecondNumber {
					secondNumber += string(m)
					continue
				}

				firstNumber += string(m)
			}

			if m == ',' {
				parsingSecondNumber = true
			}
		}

		fNumber, err := strconv.ParseInt(firstNumber, 10, 64)
		if err != nil {
			return "", fmt.Errorf("could not parse first number: %w", err)
		}

		sNumber, err := strconv.ParseInt(secondNumber, 10, 64)
		if err != nil {
			return "", fmt.Errorf("could not parse second number: %w", err)
		}

		result += fNumber * sNumber
	}

	return fmt.Sprintf("%v", result), nil
}
