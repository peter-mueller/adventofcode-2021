package day10

type tokenType int

const (
	TokenUnknown tokenType = iota
	TokenOpen
	TokenClose
)

type token struct {
	typ   tokenType
	value rune
}

func Lex(input string) chan token {
	l := &lexer{
		input:  input,
		tokens: make(chan token, 1000),
	}
	go l.lex()
	return l.tokens
}

type lexer struct {
	input  string
	tokens chan token
}

func (l *lexer) lex() {
	defer close(l.tokens)
	for _, r := range l.input {
		typ := TokenUnknown
		switch {
		case isAnyOf(r, "([{<"):
			typ = TokenOpen
		case isAnyOf(r, ")]}>"):
			typ = TokenClose
		}
		l.tokens <- token{
			typ:   typ,
			value: r,
		}
	}
}

func isAnyOf(r rune, s string) bool {
	for _, sr := range s {
		if r == sr {
			return true
		}
	}
	return false
}
