package day06

import (
	_ "embed"
	"log"
	"strconv"
)

type Submarine interface {
	ScanNearbyLanternfish() <-chan Lanternfish
}

func NewSubmarine() Submarine {
	return &mockSubmarine{}
}

type mockSubmarine struct{}

//go:embed input.txt
var input string

func (mockSubmarine) ScanNearbyLanternfish() <-chan Lanternfish {
	c := make(chan Lanternfish, 10000)

	go func() {
		defer close(c)
		pos := 0
		for i, r := range input {
			if r == ',' {
				c <- parseFish(input[pos:i])
				pos = i + 1
			}
		}
		c <- parseFish(input[pos:])
	}()

	return c
}

func parseFish(s string) Lanternfish {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return Lanternfish{BirthIn: Duration(number)}
}
