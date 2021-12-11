package day08

import (
	"bufio"
	_ "embed"
	"log"
	"strings"
)

type NotebookLine struct {
	Patterns [10]Pattern
	Final    [4]Pattern
}

type Notebook struct {
	Lines []NotebookLine
}

//go:embed input.txt
var input string

func TakeNote() (notebook Notebook) {
	tokens := lex(input)

	for {
		t, ok := <-tokens
		if !ok {
			return notebook
		}
		if t.typ != TokenNext {
			log.Fatal("expected next line")
		}
		line := NotebookLine{}
		for i := 0; i < 10; i++ {
			t := <-tokens
			if t.typ != TokenPattern {
				log.Fatal("expected pattern")
			}
			line.Patterns[i] = Pattern(t.value).Normalize()
		}
		if (<-tokens).typ != TokenDivider {
			log.Fatal("expected divider")
		}
		for i := 0; i < 4; i++ {
			t := <-tokens
			if t.typ != TokenPattern {
				log.Fatal("expected pattern")
			}
			line.Final[i] = Pattern(t.value).Normalize()
		}
		notebook.Lines = append(notebook.Lines, line)
	}
}

type tokenType int

const (
	TokenUnknown tokenType = iota
	TokenPattern
	TokenDivider
	TokenNext
)

type token struct {
	typ   tokenType
	value string
}

func lex(input string) <-chan token {
	c := make(chan token, 1000)
	go func() {
		defer close(c)
		scanner := bufio.NewScanner(strings.NewReader(input))
		for scanner.Scan() {
			c <- token{
				typ: TokenNext,
			}
			lexLine(scanner.Text(), c)
		}

	}()
	return c
}

func lexLine(line string, c chan token) {
	scanner := bufio.NewScanner(strings.NewReader(line))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		typ := TokenUnknown
		switch text {
		case "|":
			typ = TokenDivider
		default:
			typ = TokenPattern
		}
		c <- token{
			typ:   typ,
			value: scanner.Text(),
		}
	}
}
