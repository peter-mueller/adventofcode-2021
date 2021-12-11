package day10

import (
	"fmt"
	"sort"

	"github.com/peter-mueller/adventofcode-2021/puzzle"
)

var Puzzles = []puzzle.Puzzle{
	Puzzle1{},
	Puzzle2{},
}

type Puzzle1 struct{}

func (Puzzle1) PrintQuestion() {
	fmt.Println("> Q1")
}

func (Puzzle1) PrintAnswer() {
	inputs := ReadNavigationSubsystem()

	score := 0
	for input := range inputs {
		errs := Parse(input)
		if len(errs) > 0 {
			score += ScoreForErrUnexpected(errs[0])
		}
	}

	fmt.Println(score)
}

func ScoreForErrUnexpected(err error) int {
	switch e := err.(type) {
	case ErrUnexpected:
		switch e.Found {
		case ')':
			return 3
		case ']':
			return 57
		case '}':
			return 1197
		case '>':
			return 25137
		}
	}
	return 0
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> Q2")
}

func (Puzzle2) PrintAnswer() {
	inputs := ReadNavigationSubsystem()

	scores := make([]int, 0)

nextinput:
	for input := range inputs {
		errs := Parse(input)

		score := 0
		for _, err := range errs {
			if _, ok := err.(ErrUnexpected); ok {
				// ignore
				goto nextinput
			}
			score = (score * 5) + ScoreForErrExpected(err)
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	middleScore := scores[len(scores)/2]
	fmt.Println(middleScore)
}

func ScoreForErrExpected(err error) int {
	switch e := err.(type) {
	case ErrExpected:
		switch e.Expected {
		case ')':
			return 1
		case ']':
			return 2
		case '}':
			return 3
		case '>':
			return 4
		}
	}
	return 0
}
