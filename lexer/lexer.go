package lexer

import (
	"strings"
	"unicode/utf8"

	"github.com/nownabe/brainfuck-go/config"
	"github.com/nownabe/brainfuck-go/token"
)

type Lexer struct {
	config   config.Config
	input    []string
	position int
	max      int
	length   int
}

func New(c config.Config, input []string) *Lexer {
	return &Lexer{
		config:   c,
		input:    input,
		position: 0,
		max:      getMaxLength(c),
		length:   len(input),
	}
}

func length(s string) int {
	return utf8.RuneCountInString(s)
}

func getMaxLength(c config.Config) int {
	max := 0

	ops := []string{c.NEXT, c.PREV, c.INC, c.DEC, c.READ, c.WRITE, c.OPEN, c.CLOSE}

	for _, op := range ops {
		if l := length(op); max < l {
			max = l
		}
	}

	return max
}

func (l *Lexer) Next() token.Token {
	l.skipWhitespace()
	buf := ""
	for {
		if l.isEOF() {
			return token.EOF
		}
		if length(buf) >= l.max {
			return token.INVALID
		}
		buf = buf + l.readChar()
		switch buf {
		case l.config.NEXT:
			return token.NEXT
		case l.config.PREV:
			return token.PREV
		case l.config.INC:
			return token.INC
		case l.config.DEC:
			return token.DEC
		case l.config.READ:
			return token.READ
		case l.config.WRITE:
			return token.WRITE
		case l.config.OPEN:
			return token.OPEN
		case l.config.CLOSE:
			return token.CLOSE
		}
	}
}

func (l *Lexer) skipWhitespace() {
	for {
		if l.isEOF() {
			return
		}

		if l.isMatchWhitespace(l.input[l.position]) {
			l.position++
		} else {
			return
		}
	}
}

func (l *Lexer) isMatchWhitespace(c string) (matched bool) {
	return strings.Contains(l.config.WHITESPACES, c)
}

func (l *Lexer) readChar() string {
	pos := l.position
	l.position++
	return l.input[pos]
}

func (l *Lexer) isEOF() bool {
	return l.position >= l.length
}
