package day03

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
	fmt.Println("> Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together.")
	fmt.Println("> What is the power consumption of the submarine?")
}

func (Puzzle1) PrintAnswer() {
	report := NewDiagnosticReport()
	// not efficient:
	calc := NewPowerCalc(report.Read())
	result := calc.PowerConsumption()
	fmt.Println(result)
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> Use the binary numbers in your diagnostic report to calculate the oxygen generator rating and CO2 scrubber rating, then multiply them together.")
	fmt.Println("> What is the life support rating of the submarine?")
}

func (Puzzle2) PrintAnswer() {
	report := NewDiagnosticReport()
	// not efficient:
	calc := NewLifeSupportRatingCalc(report.Read())
	result := calc.LifeSupportRating()
	fmt.Println(result)
}
