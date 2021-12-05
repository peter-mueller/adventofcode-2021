package day05

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
	fmt.Println("> Consider only horizontal and vertical lines. At how many points do at least two lines overlap?")
}

func (Puzzle1) PrintAnswer() {
	vents := NewSubmarine().ProduceVents()

	image := make(map[Point]int)
	for v := range vents {
		if !(v.IsHorizontal() || v.IsVertical()) {
			continue
		}
		for _, p := range v.Points() {
			image[Point{X: p.X, Y: p.Y}] += 1
		}
	}

	amount := 0
	for _, hits := range image {
		if hits > 1 {
			amount += 1
		}
	}
	fmt.Println(amount)
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> Consider all of the lines. At how many points do at least two lines overlap?")
}

func (Puzzle2) PrintAnswer() {
	vents := NewSubmarine().ProduceVents()

	image := make(map[Point]int)
	for v := range vents {
		for _, p := range v.Points() {
			image[Point{X: p.X, Y: p.Y}] += 1
		}
	}

	amount := 0
	for _, hits := range image {
		if hits > 1 {
			amount += 1
		}
	}
	fmt.Println(amount)
}
