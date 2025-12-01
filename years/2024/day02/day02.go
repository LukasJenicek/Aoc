package day202402

import (
	"fmt"

	"github.com/LukasJenicek/aoc/internal/helpers"
)

type SecondDay struct{}

func (*SecondDay) RunFirstPart(input string) (string, error) {
	reports, err := helpers.LoadNumbers(input)
	if err != nil {
		return "", fmt.Errorf("load numbers: %w", err)
	}

	safeReports := 0
	for _, report := range reports {
		safe := isLvlSafe(report)
		if safe {
			safeReports++
		}
	}

	return fmt.Sprintf("%v", safeReports), nil
}

func (*SecondDay) RunSecondPart(input string) (string, error) {
	reports, err := helpers.LoadNumbers(input)
	if err != nil {
		return "", fmt.Errorf("load numbers: %w", err)
	}

	safeReports := 0
	for _, report := range reports {
		if len(report) == 0 {
			continue
		}

		valid := isLvlSafe(report)
		if valid {
			safeReports++
			continue
		}

		for i := range len(report) {
			if isLvlSafe(helpers.RemoveAtIndex(report, i)) {
				safeReports++
				break
			}
		}
	}

	return fmt.Sprintf("%v", safeReports), nil
}

func isLvlSafe(report []int) bool {
	prev := 0
	increasing := false

	for i, level := range report {
		if i == 0 {
			prev = level
			continue
		}

		diff := helpers.AbsoluteDiff(prev, level)
		if diff == 0 || diff > 3 {
			return false
		}

		if i == 1 {
			increasing = level > prev
			prev = level
			continue
		}

		if (level > prev) != increasing {
			return false
		}

		prev = level
	}

	return true
}
