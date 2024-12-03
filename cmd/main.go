package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/LukasJenicek/aoc/days/day01"
	"github.com/LukasJenicek/aoc/days/day02"
	"github.com/LukasJenicek/aoc/days/day03"
	"github.com/LukasJenicek/aoc/libs/runner"
)

var (
	flags   = flag.NewFlagSet("aoc", flag.ExitOnError)
	day     = flags.Int("day", 1, "select which day to run")
	part    = flags.Int("part", 1, "select which part to run")
	example = flags.Bool("example", false, "example input")
)

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	input, err := loadInput(*day, *part, *example)
	if err != nil {
		log.Fatal(err)
	}

	r := runner.NewRunner(map[int]runner.AssignmentRunnable{
		1: &day01.FirstDay{},
		2: &day02.SecondDay{},
		3: &day03.ThirdDay{},
	})

	result, err := r.Run(*day, *part, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func loadInput(day int, part int, example bool) (string, error) {
	dayFormatted := fmt.Sprintf("%02d", day)
	partFormatted := fmt.Sprintf("%02d", part)

	workingDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("get working directory: %w", err)
	}

	var filename string
	if example {
		filename = fmt.Sprintf("%s/days/day%s/example%s.txt", workingDir, dayFormatted, partFormatted)
	} else {
		filename = fmt.Sprintf("%s/days/day%s/input%s.txt", workingDir, dayFormatted, partFormatted)
	}

	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return string(contents), nil
}
