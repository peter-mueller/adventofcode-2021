package day01

import (
	"bufio"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

type MockSonar struct {
}

func (s *MockSonar) Sweep() <-chan Measurement {
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
