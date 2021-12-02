package day02

import (
	"bufio"
	"strings"

	_ "embed"
)

type MissionControl interface {
	Listen() <-chan string
}

//go:embed input.txt
var input string

func NewMissionControl() MissionControl {
	return &mockMissionControl{}
}

type mockMissionControl struct{}

func (mockMissionControl) Listen() <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()
	return c
}
