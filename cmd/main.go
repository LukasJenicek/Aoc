package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/LukasJenicek/aoc/days/day01"
)

var (
	flags = flag.NewFlagSet("aoc", flag.ExitOnError)
	day   = flags.Int("day", 1, "select which day to run")
)

func main() {
	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	var result string
	var err error

	switch *day {
	case 1:
		result, err = day01.RunFirst()
	default:
		log.Fatalf("day %d not implemented", *day)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
