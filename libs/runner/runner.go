package runner

import (
	"fmt"
)

type AssignmentRunnable interface {
	RunFirstPart(input string) (string, error)
	RunSecondPart(input string) (string, error)
}

type Runner struct {
	assignments map[int]AssignmentRunnable
}

func NewRunner(assignments map[int]AssignmentRunnable) *Runner {
	return &Runner{assignments: assignments}
}

func (r *Runner) Run(day int, part int, input string) (string, error) {
	assignment, ok := r.assignments[day]
	if !ok {
		return "", fmt.Errorf("assignment not found for day %d", day)
	}

	if part == 1 {
		return assignment.RunFirstPart(input)
	}

	return assignment.RunSecondPart(input)
}
