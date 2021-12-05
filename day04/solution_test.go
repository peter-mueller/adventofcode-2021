package day04_test

import (
	"testing"

	"github.com/peter-mueller/adventofcode-2021/day04"
)

func BenchmarkPuzzle1(b *testing.B) {
	p := day04.Puzzle1{}
	for n := 0; n < b.N; n++ {
		p.PrintAnswer()
	}
}
