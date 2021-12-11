package day08

import (
	"fmt"

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
	notebook := TakeNote()
	total := 0
	for _, line := range notebook.Lines {
		result := Solve(line)
		total += count1478(line, result)
	}
	fmt.Println(total)
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> Q2")
}

func (Puzzle2) PrintAnswer() {
	notebook := TakeNote()
	total := 0
	for _, line := range notebook.Lines {
		result := Solve(line)
		total += number(line, result)
	}
	fmt.Println(total)
}

func count1478(line NotebookLine, result SegmentPatterns) (count int) {
	for _, finalPattern := range line.Final {
		switch finalPattern.String() {
		case result[1].String():
			count++
		case result[4].String():
			count++
		case result[7].String():
			count++
		case result[8].String():
			count++
		}
	}
	return count
}

func number(line NotebookLine, result SegmentPatterns) (number int) {
	for i, finalPattern := range line.Final {
		s := finalPattern.String()
		for ri := range result {
			if result[ri].String() == s {
				number += pow10(len(line.Final)-1-i) * ri
				break
			}
		}
	}
	return number
}

func pow10(n int) (res int) {
	res = 1
	for i := 0; i < n; i++ {
		res *= 10
	}
	return res
}
