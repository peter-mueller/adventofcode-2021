package day07

import (
	"log"
	"strconv"
	"strings"

	_ "embed"
)

//go:embed input.txt
var input string

func NewCrabList() (list []CrabSubmarine) {
	numbers := strings.Split(input, ",")
	list = make([]CrabSubmarine, len(numbers))
	for i, n := range numbers {
		list[i] = CrabSubmarine{
			HorizontalPosition: atoi(n),
		}
	}
	return list
}

func atoi(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
