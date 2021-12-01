package day01

import (
	"bufio"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type Sonar interface {
	Sweep() <-chan Measurement
}

type Measurement int

const (
	NoMeasurement Measurement = -1
)

func (m Measurement) IsValid() bool {
	if m == NoMeasurement {
		return false
	}
	return true
}

func NewSonar() Sonar {
	return &mockSonar{}
}

type mockSonar struct{}

func (s *mockSonar) Sweep() <-chan Measurement {
	c := make(chan Measurement)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			number, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			c <- Measurement(number)
		}
	}()
	return c
}
