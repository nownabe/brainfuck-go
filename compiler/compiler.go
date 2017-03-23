package compiler

import (
	"strings"

	"github.com/nownabe/brainfuck-go/config"
	"github.com/nownabe/brainfuck-go/lexer"
	"github.com/nownabe/brainfuck-go/token"
)

type Compiler struct {
	input string
	lexer *lexer.Lexer
}

func New(c config.Config, i string) *Compiler {
	return &Compiler{
		input: i,
		lexer: lexer.New(c, strings.Split(i, "")),
	}
}

func (c *Compiler) Compile() (binary []token.Token) {
	for {
		t := c.lexer.Next()
		switch t {
		case token.INVALID:
			panic("Invalid token")
		case token.EOF:
			binary = append(binary, t)
			goto FINISH
		default:
			binary = append(binary, t)
		}
	}
FINISH:
	return
}
