package day05

import (
	"bufio"
	_ "embed"
	"strings"
)

type Submarine interface {
	ProduceVents() chan Vent
}

func NewSubmarine() Submarine {
	return &mockSubmarine{}
}

type mockSubmarine struct{}

//go:embed input.txt
var input string

func (mockSubmarine) ProduceVents() chan Vent {
	c := make(chan Vent)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			vent := ParseVent(scanner.Text())
			c <- vent
		}
	}()
	return c
}
