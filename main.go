package main

import (
	"fmt"

	"github.com/peter-mueller/adventofcode-2021/day01"
	"github.com/peter-mueller/adventofcode-2021/day02"
	"github.com/peter-mueller/adventofcode-2021/day03"
	"github.com/peter-mueller/adventofcode-2021/day04"
	"github.com/peter-mueller/adventofcode-2021/day05"
	"github.com/peter-mueller/adventofcode-2021/day06"
	"github.com/peter-mueller/adventofcode-2021/day07"
	"github.com/peter-mueller/adventofcode-2021/day08"
	"github.com/peter-mueller/adventofcode-2021/day09"
	"github.com/peter-mueller/adventofcode-2021/day10"
	"github.com/peter-mueller/adventofcode-2021/puzzle"
)

type Day int

var puzzlesPerDay = map[Day][]puzzle.Puzzle{
	1:  day01.Puzzles,
	2:  day02.Puzzles,
	3:  day03.Puzzles,
	4:  day04.Puzzles,
	5:  day05.Puzzles,
	6:  day06.Puzzles,
	7:  day07.Puzzles,
	8:  day08.Puzzles,
	9:  day09.Puzzles,
	10: day10.Puzzles,
}

func main() {
	fmt.Println("==============")
	fmt.Println("Advent of Code")
	fmt.Println("  {:year 2021}")
	fmt.Println("==============")
	fmt.Println()

	for day := Day(1); int(day) <= 25; day++ {
		puzzles, ok := puzzlesPerDay[day]
		if !ok {
			continue
		}

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
