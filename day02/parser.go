package day02

import (
	"log"
	"strconv"
	"strings"
)

type parser struct {
	input string
	pos   int
}

const (
	tokenNone    = ""
	tokenForward = "forward"
	tokenUp      = "up"
	tokenDown    = "down"
	tokenSpace   = " "
)

var commandTokens = []string{tokenForward, tokenUp, tokenDown}

func ParseCommand(s string) Command {
	p := &parser{input: s}
	return p.parse()
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

func (p *parser) parse() Command {
	invalidCommand := func() {
		log.Fatalf("invalid command: %s", p.input)
	}

	token := tokenNone
	for _, t := range commandTokens {
		ok := p.accept(t)
		if ok {
			token = t
			continue
		}
	}
	if token == tokenNone {
		invalidCommand()
	}
	if !p.accept(tokenSpace) {
		invalidCommand()
	}

	switch token {
	case tokenForward:
		amount, err := strconv.Atoi(p.peek())
		if err != nil {
			invalidCommand()
		}
		return CommandForward{Amount: amount}
	case tokenUp:
		amount, err := strconv.Atoi(p.peek())
		if err != nil {
			invalidCommand()
		}
		return CommandUp{Amount: amount}
	case tokenDown:
		amount, err := strconv.Atoi(p.peek())
		if err != nil {
			invalidCommand()
		}
		return CommandDown{Amount: amount}
	default:
		invalidCommand()
		return nil
	}
}
