package day09

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
	hm := GenerateHeightMap()

	locations := FilterKernel(IsLowPoint).Apply(hm)

	total := 0
	for _, l := range locations {
		h, _ := hm.At(l)
		risklevel := int(h) + 1
		total += risklevel
	}

	fmt.Println(total)
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> Q2")
}

func (Puzzle2) PrintAnswer() {

	hm := GenerateHeightMap()

	lowPoints := FilterKernel(IsLowPoint).Apply(hm)

	threelargest := [3]int{1, 1, 1}
	for _, lowpoint := range lowPoints {
		kernel := SelectKernel(NearBasinLocations)
		locations := kernel.SelectFrom(lowpoint, hm)
		size := len(locations)

		for i := range threelargest {
			if size > threelargest[i] {
				threelargest[i] = size
				break
			}
		}
	}

	total := 1
	for _, n := range threelargest {
		total *= n
	}
	fmt.Println(total)
}

func IsLowPoint(l Location, hm *Heightmap) bool {
	h, _ := hm.At(l)
	neighbors := [4]Location{
		l.North(), l.South(), l.East(), l.West(),
	}
	for _, n := range neighbors {
		nh, ok := hm.At(n)
		if !ok {
			continue
		}
		if h >= nh {
			// is no lowpoint
			return false
		}
	}
	return true
}

func NearBasinLocations(lowpoint Location, hm *Heightmap) (selected []Location) {
	neighbors := [4]Location{
		lowpoint.North(), lowpoint.South(), lowpoint.East(), lowpoint.West(),
	}
	for _, n := range neighbors {
		nh, ok := hm.At(n)
		if !ok {
			continue
		}
		if nh < 9 {
			// flows down if l was a lowpoint
			selected = append(selected, n)
		}
	}
	return selected
}
