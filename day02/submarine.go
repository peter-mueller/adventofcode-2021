package day02

import "log"

type Submarine struct {
	Pos Position
	Aim int
}

type Position struct {
	X     int
	Depth int
}

func NewSubmarine() *Submarine {
	return &Submarine{}
}

type Command interface{}

type CommandForward struct {
	Amount int
}
type CommandUp struct {
	Amount int
}
type CommandDown struct {
	Amount int
}

func (s *Submarine) Execute(command interface{}) {
	switch c := command.(type) {
	case CommandForward:
		s.Pos.X += c.Amount
	case CommandUp:
		s.Pos.Depth -= c.Amount
	case CommandDown:
		s.Pos.Depth += c.Amount
	default:
		log.Fatalf("unkown command: %T", command)
	}
}

func (s *Submarine) ExecuteAimVariant(command interface{}) {
	switch c := command.(type) {
	case CommandForward:
		s.Pos.X += c.Amount
		s.Pos.Depth += s.Aim * c.Amount
	case CommandUp:
		s.Aim -= c.Amount
	case CommandDown:
		s.Aim += c.Amount
	default:
		log.Fatalf("unkown command: %T", command)
	}
}
