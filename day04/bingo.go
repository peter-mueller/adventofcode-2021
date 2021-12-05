package day04

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type BingoSubsystem interface {
	Numbers() <-chan int
	Boards() <-chan Board
}

type BoardNumber struct {
	Value  int
	Marked bool
}

type Board struct {
	Numbers [5][5]BoardNumber
}

func (b *Board) IsFinished() bool {
	for row := range b.Numbers {
		for column := range b.Numbers[row] {
			if !b.Numbers[row][column].Marked {
				goto nextrow
			}
		}
		// found row
		return true
	nextrow:
	}
	for column := range b.Numbers {
		for row := range b.Numbers[column] {
			if !b.Numbers[row][column].Marked {
				goto nextcolumn
			}
		}
		// found column
		return true
	nextcolumn:
	}
	return false
}

func (b *Board) Mark(number int) (marked bool) {
	for row := range b.Numbers {
		for col := range b.Numbers[row] {
			if b.Numbers[row][col].Value == number {
				b.Numbers[row][col].Marked = true
				return true
			}
		}
	}
	return false
}

func (b *Board) Score(lastCalledNumber int) (score int) {
	for row := range b.Numbers {
		for _, n := range b.Numbers[row] {
			if !n.Marked {
				score += n.Value
			}
		}
	}
	return score * lastCalledNumber
}

func NewBingoSubsystem() BingoSubsystem {
	return &mockSubsystem{}
}

type mockSubsystem struct{}

//go:embed input.txt
var input string

func (mockSubsystem) Boards() <-chan Board {
	c := make(chan Board)

	go func() {
		defer close(c)

		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Scan() // skip numbers line

		for scanner.Scan() {
			emptyLine := scanner.Text()
			if emptyLine != "" {
				log.Fatalf("expected empty line for block start, but was: %s", emptyLine)
			}

			var board Board
			for row := range board.Numbers {
				if !scanner.Scan() {
					log.Fatal("expectet row of numbers")
				}
				text := scanner.Text()
				fmt.Sscanf(text, "%d %d %d %d %d\n",
					&board.Numbers[row][0].Value,
					&board.Numbers[row][1].Value,
					&board.Numbers[row][2].Value,
					&board.Numbers[row][3].Value,
					&board.Numbers[row][4].Value,
				)
			}
			c <- board
		}

	}()
	return c
}

func (mockSubsystem) Numbers() <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		var (
			pos = 0
		)
		for i, r := range input {
			switch r {
			case ',':
				text := input[pos:i]
				value, err := strconv.Atoi(text)
				if err != nil {
					log.Fatal(err)
				}
				c <- value
				pos = i + 1
				continue
			case '\n':
				return
			}
		}
	}()
	return c
}
