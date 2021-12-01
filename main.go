package main

import (
	"fmt"

	"github.com/peter-mueller/adventofcode-2021/day01"
	"github.com/peter-mueller/adventofcode-2021/puzzle"
)

type Day int

var puzzlesPerDay = map[Day][]puzzle.Puzzle{
	1: day01.Puzzles,
}

func main() {
	fmt.Println("==============")
	fmt.Println("Advent of Code")
	fmt.Println("  {:year 2021}")
	fmt.Println("==============")
	fmt.Println()

	for day, puzzles := range puzzlesPerDay {
		fmt.Printf("# Day %02d, see %s\n", day, linkToDay(day))
		fmt.Println()
		for _, p := range puzzles {
			p.PrintQuestion()
			p.PrintAnswer()
		}
		fmt.Println()
	}
}

func linkToDay(day Day) string {
	return fmt.Sprintf("https://adventofcode.com/2021/day/%d/", day)
}
