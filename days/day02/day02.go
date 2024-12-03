package day02

import (
	"fmt"

	"github.com/LukasJenicek/aoc/libs/helpers"
)

type SecondDay struct{}

func (*SecondDay) RunFirstPart(input string) (string, error) {
	reports, err := helpers.LoadNumbers(input)
	if err != nil {
		return "", fmt.Errorf("load numbers: %w", err)
	}

	safeReports := 0
	for _, report := range reports {
		safe := true
		increasing := false
		for j, level := range report {
			if j == len(report)-1 {
				continue
			}

			val := level - report[j+1]
			if val == 0 {
				safe = false
				break
			}

			if j == 0 && val < 0 {
				increasing = true
			}

			if increasing && (val > 0 || val < -3) {
				safe = false
				break
			}

			if !increasing && (val < 0 || val > 3) {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}
	}

	return fmt.Sprintf("%v", safeReports), nil
}

func (*SecondDay) RunSecondPart(input string) (string, error) {
	return "", nil
}
