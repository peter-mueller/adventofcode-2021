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
	var (
		moderator = moderator{}
		results   = make(chan result)
	)

	for board := range boards {
		p := player{board: board}
		c := moderator.subscribe()

		go func() {
			defer moderator.markDone()
			res, ok := <-p.play(c)
			if ok {
				results <- res
			}
		}()
	}

	go moderator.announce(input)

	go func() {
		defer close(results)
		moderator.waitAllDone()
	}()
	return results
}

type moderator struct {
	chanToPlayers []chan int
	wg            sync.WaitGroup
}

func (m *moderator) subscribe() <-chan int {
	c := make(chan int)
	m.chanToPlayers = append(m.chanToPlayers, c)
	m.wg.Add(1)
	return c
}

func (m *moderator) markDone() {
	m.wg.Done()
}

func (m *moderator) waitAllDone() {
	m.wg.Wait()
}

func (m *moderator) announce(numbers <-chan int) {
	for n := range numbers {
		for _, c := range m.chanToPlayers {
			c <- n
		}
	}
	for _, c := range m.chanToPlayers {
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
