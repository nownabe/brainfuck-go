package lexer

import (
	//"fmt"
	"reflect"
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

	v := reflect.ValueOf(c)

	for i := 0; i < v.NumField(); i++ {
		if l := length(v.Field(i).String()); max < l {
			max = l
		}
	}

	return max
}

func (l *Lexer) Next() token.Token {
	l.skipSpace()
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

func (l *Lexer) skipSpace() {
	for {
		if l.isEOF() {
			return
		}
		switch l.input[l.position] {
		case " ", "\t", "\n", "\r":
			l.position++
		default:
			return
		}
	}
}

func (l *Lexer) readChar() string {
	pos := l.position
	l.position++
	return l.input[pos]
}

func (l *Lexer) isEOF() bool {
	return l.position >= l.length
}
