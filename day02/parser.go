package day02

import (
	"log"
	"strconv"
	"strings"
)

func ParseCommand(s string) Command {
	p := &parser{input: s}
	return p.parse()
}

type parser struct {
	input string
	pos   int
}

const (
	tokenForward = "forward"
	tokenUp      = "up"
	tokenDown    = "down"
	tokenSpace   = " "
)

func (p *parser) parse() Command {

	switch {
	case p.accept(tokenForward):
		p.mustAccept(tokenSpace)
		amount := p.mustParseInt(p.peek())
		return CommandForward{Amount: amount}
	case p.accept(tokenUp):
		p.mustAccept(tokenSpace)
		amount := p.mustParseInt(p.peek())
		return CommandUp{Amount: amount}
	case p.accept(tokenDown):
		p.mustAccept(tokenSpace)
		amount := p.mustParseInt(p.peek())
		return CommandDown{Amount: amount}
	default:
		p.fail()
	}
	return nil
}

func (p *parser) advance(token string) {
	p.pos += len(token)
}

func (p *parser) peek() string {
	return p.input[p.pos:]
}

func (p *parser) accept(token string) bool {
	if strings.HasPrefix(p.peek(), token) {
		p.advance(token)
		return true
	}
	return false
}

func (p *parser) fail() {
	log.Fatalf("invalid command: %s", p.input)
}

func (p *parser) mustAccept(token string) {
	ok := p.accept(token)
	if !ok {
		p.fail()
	}
}
func (p *parser) mustParseInt(token string) int {
	value, err := strconv.Atoi(p.peek())
	if err != nil {
		p.fail()
	}
	return value
}
