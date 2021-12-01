package day01

import (
	"fmt"

	"github.com/peter-mueller/adventofcode-2021/puzzle"
)

var Puzzles = []puzzle.Puzzle{
	Puzzle1{},
	Puzzle2{},
}

type Sonar interface {
	Sweep() <-chan Measurement
}

type Measurement int

const (
	NoMeasurement Measurement = -1
)

type Puzzle1 struct{}

func (Puzzle1) PrintQuestion() {
	fmt.Println("> How many measurements are larger than the previous measurement?")
}

func (Puzzle1) PrintAnswer() {
	sonar := MockSonar{}

	var (
		amountIncreasing = 0
		lastSum          = NoMeasurement
		window           = NewWindow(1)
	)
	for measurement := range sonar.Sweep() {
		window.Put(measurement)
		newSum := window.Sum()
		if lastSum != NoMeasurement && newSum > lastSum {
			amountIncreasing += 1
		}
		lastSum = newSum
	}
	fmt.Println(amountIncreasing)
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> Consider sums of a three-measurement sliding window.")
	fmt.Println("> How many sums are larger than the previous sum?")
}

func (Puzzle2) PrintAnswer() {
	sonar := MockSonar{}

	var (
		amountIncreasing = 0
		lastSum          = NoMeasurement
		window           = NewWindow(3)
	)
	for measurement := range sonar.Sweep() {
		window.Put(measurement)
		newSum := window.Sum()
		if lastSum != NoMeasurement && newSum > lastSum {
			amountIncreasing += 1
		}
		lastSum = newSum
	}
	fmt.Println(amountIncreasing)
}
