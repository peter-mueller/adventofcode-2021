package day10

import (
	"fmt"
)

type ErrUnexpected struct {
	Expected rune
	Found    rune
}

func (e ErrUnexpected) Error() string {
	return fmt.Sprintf("Expected %s, but found %s instead", string(e.Expected), string(e.Found))
}

type ErrExpected struct {
	Expected rune
}

func (e ErrExpected) Error() string {
	return fmt.Sprintf("Expected %s", string(e.Expected))
}

func Parse(input string) []error {
	tokens := Lex(input)
	errors := make([]error, 0)

	var (
		openchunks = make([]token, 0)
	)
	for t := range tokens {
		switch t.typ {
		case TokenOpen:
			openchunks = append(openchunks, t)
		case TokenClose:
			last := len(openchunks) - 1
			openToken := openchunks[last]
			if closingValueFor(openToken.value) != t.value {
				err := ErrUnexpected{
					Expected: openToken.value,
					Found:    t.value,
				}
				errors = append(errors, err)
			}
			openchunks = openchunks[:last]
		}
	}

	for i := range openchunks {
		t := openchunks[len(openchunks)-1-i]
		err := ErrExpected{Expected: closingValueFor(t.value)}
		errors = append(errors, err)
	}
	return errors
}

func closingValueFor(r rune) rune {
	switch r {
	case '(':
		return ')'
	case '{':
		return '}'
	case '[':
		return ']'
	case '<':
		return '>'
	}
	panic("there is no closing value for rune " + string(r))
}
