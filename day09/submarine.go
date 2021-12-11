package day09

import (
	"bufio"
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func GenerateHeightMap() *Heightmap {
	hm := Heightmap{
		data: make(map[Location]Height),
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "\n" {
			hm.bounds.Max.Y++
			hm.bounds.Max.X = 0
			continue
		}
		number := atoi(text)
		hm.data[hm.bounds.Max] = Height(number)
		hm.bounds.Max.X += 1
	}

	hm.bounds.Max.X -= 1
	return &hm
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return n
}
