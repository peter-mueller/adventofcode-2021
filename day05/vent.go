package day05

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Vent struct {
	P1, P2 Point
}

type Point struct {
	X, Y int
}

func (v Vent) IsHorizontal() bool {
	return v.P1.Y == v.P2.Y
}
func (v Vent) IsVertical() bool {
	return v.P1.X == v.P2.X
}

func (p Point) Subtract(o Point) (r Point) {
	r.X = p.X - o.X
	r.Y = p.Y - o.Y
	return r
}

func (p Point) Add(o Point) (r Point) {
	r.X = p.X + o.X
	r.Y = p.Y + o.Y
	return r
}

func (v Vent) Points() (points []Point) {
	p1, p2 := v.P1, v.P2
	delta := p2.Subtract(p1)
	delta.X = normalize(delta.X)
	delta.Y = normalize(delta.Y)
	pencil := p1
	for {
		points = append(points, pencil)
		if pencil == p2 {
			return points
		}
		pencil = pencil.Add(delta)
	}
}

func normalize(value int) int {
	switch {
	case value < 0:
		return -1
	case value > 0:
		return 1
	}
	return 0
}

func ParseVent(s string) (v Vent) {
	p := parser{input: s}
	v.P1.X = p.int()
	p.must(",")
	v.P1.Y = p.int()
	p.must(" -> ")
	v.P2.X = p.int()
	p.must(",")
	v.P2.Y = p.int()
	return v
}

type parser struct {
	input string
	pos   int
}

func (p *parser) int() int {
	number := p.peek()
	for i, r := range p.peek() {
		if !unicode.IsDigit(r) {
			number = p.input[p.pos : p.pos+i]
			break
		}
	}

	p.advance(number)
	value, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func (p *parser) peek() string {
	return p.input[p.pos:]
}
func (p *parser) advance(s string) {
	p.pos += len(s)
}
func (p *parser) must(s string) {
	if strings.HasPrefix(s, p.peek()) {
		log.Fatalf("expected %s", s)
	}
	p.advance(s)
}
