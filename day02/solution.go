package day02

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
	fmt.Println("Calculate the horizontal position and depth you would have after following the planned course.")
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?")
}

func (Puzzle1) PrintAnswer() {
	submarine := NewSubmarine()
	missionControl := NewMissionControl()

	for sentence := range missionControl.Listen() {
		command := ParseCommand(sentence)
		submarine.Execute(command)
	}

	fmt.Printf("Submarine at: %+v\n", submarine.Pos)
	fmt.Println(submarine.Pos.X * submarine.Pos.Depth)
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("Using this new interpretation of the commands, calculate the horizontal position and depth you would have after following the planned course.")
	fmt.Println("What do you get if you multiply your final horizontal position by your final depth?")
}

func (Puzzle2) PrintAnswer() {
	submarine := NewSubmarine()
	missionControl := NewMissionControl()

	for sentence := range missionControl.Listen() {
		command := ParseCommand(sentence)
		submarine.ExecuteAimVariant(command)
	}

	fmt.Printf("Submarine at: %+v\n", submarine.Pos)
	fmt.Println(submarine.Pos.X * submarine.Pos.Depth)
}
