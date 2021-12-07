package day07

import (
	"fmt"
	"math"
	"sort"

	"github.com/peter-mueller/adventofcode-2021/puzzle"
)

var Puzzles = []puzzle.Puzzle{
	Puzzle1{},
	Puzzle2{},
}

type Puzzle1 struct{}

func (Puzzle1) PrintQuestion() {
	fmt.Println("> How much fuel must they spend to align to that position?")
}

func (Puzzle1) PrintAnswer() {
	list := NewCrabList()
	sort.Sort(ByPos(list))
	medianPos := list[len(list)/2].HorizontalPosition

	totalfuel := 0
	for _, crab := range list {
		totalfuel += crab.FuelTo(medianPos)
	}
	fmt.Println(totalfuel)
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> Determine the horizontal position that the crabs can align to using the least fuel possible so they can make you an escape route!")
	fmt.Println("> How much fuel must they spend to align to that position?")
}

func (Puzzle2) PrintAnswer() {
	list := NewCrabList()

	var (
		meanPos      = meanPosition(list)
		meanPosRight = int(math.Ceil(meanPos))
		meanPosLeft  = int(math.Floor(meanPos))
	)
	var (
		totalFuelToRight = 0
		totalFuelToLeft  = 0
	)
	for _, crab := range list {
		totalFuelToRight += crab.FuelVariantTo(meanPosLeft)
		totalFuelToLeft += crab.FuelVariantTo(meanPosRight)
	}

	bestTotalFuel := min(totalFuelToLeft, totalFuelToRight)
	fmt.Println(bestTotalFuel)
}

func min(n1, n2 int) int {
	if n2 < n1 {
		return n2
	}
	return n1
}

func meanPosition(list []CrabSubmarine) float64 {
	mean := float64(0)
	for _, c := range list {
		mean += float64(c.HorizontalPosition)
	}
	mean = mean / float64(len(list))
	return mean
}
