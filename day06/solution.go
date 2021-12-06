package day06

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
	fmt.Println("> Find a way to simulate lanternfish. How many lanternfish would there be after 80 days?")
}

func (Puzzle1) PrintAnswer() {
	submarine := NewSubmarine()
	nearbyFishs := submarine.ScanNearbyLanternfish()

	deepSea := NewDeepSea()
	for f := range nearbyFishs {
		deepSea.AddFish(f)
	}

	deepSea.SimulateDuration(80 * Day)
	fmt.Println(deepSea.TotalCount())
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> How many lanternfish would there be after 256 days?")
}

func (Puzzle2) PrintAnswer() {
	submarine := NewSubmarine()
	nearbyFishs := submarine.ScanNearbyLanternfish()

	deepSea := NewDeepSea()
	for f := range nearbyFishs {
		deepSea.AddFish(f)
	}

	deepSea.SimulateDuration(256 * Day)
	fmt.Println(deepSea.TotalCount())
}
