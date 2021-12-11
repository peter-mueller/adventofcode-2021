package day10

import (
	"bufio"
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func ReadNavigationSubsystem() chan string {
	c := make(chan string, 1000)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			c <- scanner.Text()
		}
	}()
	return c
}
