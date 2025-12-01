package day202405

import (
	"fmt"
	"strconv"
	"strings"
)

type FifthDay struct{}

func (*FifthDay) RunFirstPart(input string) (string, error) {
	pagesOrdering, pages, err := loadPagesOrdering(input)
	if err != nil {
		return "", fmt.Errorf("loading pagesOrdering: %w", err)
	}

	fmt.Println(pagesOrdering)
	fmt.Println(pages)

	return "", nil
}

func (*FifthDay) RunSecondPart(input string) (string, error) {
	return "", nil
}

func loadPagesOrdering(input string) ([][]int64, [][]int64, error) {
	var pagesOrderingRules [][]int64
	var pages [][]int64

	rows := strings.Split(input, "\n")
	lastPageOrdering := 0

	for i, row := range rows {
		if row == "" {
			lastPageOrdering = i + 1
			break
		}

		p := strings.Split(row, "|")

		p1, err := strconv.ParseInt(p[0], 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing page rule %d: %v", i, err)
		}

		p2, err := strconv.ParseInt(p[1], 10, 64)
		if err != nil {
			return nil, nil, fmt.Errorf("parsing page rule %d: %v", i, err)
		}

		pagesOrderingRules = append(pagesOrderingRules, []int64{p1, p2})
	}

	for _, row := range rows[lastPageOrdering:] {
		pagesOrder := strings.Split(row, ",")

		ps := []int64{}
		for _, page := range pagesOrder {
			p1, err := strconv.ParseInt(page, 10, 64)
			if err != nil {
				return nil, nil, fmt.Errorf("parsing page %w", err)
			}

			ps = append(ps, p1)
		}

		pages = append(pages, ps)
	}

	return pagesOrderingRules, pages, nil
}
