package day04

import (
	"fmt"
	"sync"

	"github.com/peter-mueller/adventofcode-2021/puzzle"
)

var Puzzles = []puzzle.Puzzle{
	Puzzle1{},
	Puzzle2{},
}

type Puzzle1 struct{}

func (Puzzle1) PrintQuestion() {
	fmt.Println("> To guarantee victory against the giant squid, figure out which board will win first.")
	fmt.Println("> What will your final score be if you choose that board?")
}

func (Puzzle1) PrintAnswer() {
	subsystem := NewBingoSubsystem()

	results := runGame(subsystem.Boards(), subsystem.Numbers())

	firstWinner := <-results
	fmt.Println(firstWinner.score)
}

func runGame(boards <-chan Board, input <-chan int) chan result {
	inputsForPlayers := make([]chan int, 0)
	results := make(chan result)
	var wg sync.WaitGroup

	for board := range boards {
		c := make(chan int)
		inputsForPlayers = append(inputsForPlayers, c)
		p := player{board: board}

		wg.Add(1)
		go func() {
			defer wg.Done()
			res, ok := <-p.play(c)
			if ok {
				results <- res
			}
		}()
	}

	go announce(input, inputsForPlayers)

	go func() {
		defer close(results)
		wg.Wait()
	}()
	return results
}

func announce(numbers <-chan int, to []chan int) {
	for n := range numbers {
		for _, c := range to {
			c <- n
		}
	}
	for _, c := range to {
		close(c)
	}
}

type Puzzle2 struct{}

func (Puzzle2) PrintQuestion() {
	fmt.Println("> Figure out which board will win last. Once it wins, what would its final score be?")
}

func (Puzzle2) PrintAnswer() {

	subsystem := NewBingoSubsystem()

	results := runGame(subsystem.Boards(), subsystem.Numbers())

	var lastWinner result
	for r := range results {
		lastWinner = r
	}

	fmt.Printf("score: %d with %d moves", lastWinner.score, lastWinner.moves)
}
