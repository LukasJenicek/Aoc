package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/LukasJenicek/aoc/days/day01"
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

	var result string
	switch *day {
	case 1:
		result, err = day01.CalculateFirstAssignment(input)
	default:
		log.Fatalf("day %d not implemented", *day)
	}

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
		filename = fmt.Sprintf("%s/days/day%s/example.txt", workingDir, dayFormatted)
	} else {
		filename = fmt.Sprintf("%s/days/day%s/input%s.txt", workingDir, dayFormatted, partFormatted)
	}

	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return string(contents), nil
}
