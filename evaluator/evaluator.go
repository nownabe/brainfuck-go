package evaluator

import (
	"fmt"
	"os"

	"github.com/nownabe/brainfuck-go/compiler"
	"github.com/nownabe/brainfuck-go/config"
	"github.com/nownabe/brainfuck-go/token"
)

const MAX_BUFFER_SIZE = 4096

type Evaluator struct {
	binary []token.Token
	buffer []byte
	p      int
	pc     int
}

func New(c config.Config, i string) *Evaluator {
	compiler := compiler.New(c, i)
	binary := compiler.Compile()
	return &Evaluator{
		binary: binary,
		buffer: make([]byte, MAX_BUFFER_SIZE, MAX_BUFFER_SIZE),
	}
}

func (e *Evaluator) op() token.Token {
	return e.binary[e.pc]
}

func (e *Evaluator) tick() {
	e.pc++
}

func (e *Evaluator) next() {
	e.p++
}

func (e *Evaluator) prev() {
	e.p--
}

func (e *Evaluator) inc() {
	e.buffer[e.p]++
}

func (e *Evaluator) dec() {
	e.buffer[e.p]--
}

func (e *Evaluator) read() {
	b := make([]byte, 1)
	os.Stdin.Read(b)
	e.buffer[e.p] = b[0]
}

func (e *Evaluator) write() {
	b := make([]byte, 1)
	b[0] = e.buffer[e.p]
	os.Stdout.Write(b)
}

func (e *Evaluator) open() {
	if e.buffer[e.p] == 0 {
		n := 0
		for {
			e.pc++
			if e.op() == token.OPEN {
				n++
			} else if e.op() == token.CLOSE {
				if n--; n < 0 {
					break
				}
			}
		}
	}
}

func (e *Evaluator) close() {
	if e.buffer[e.p] != 0 {
		n := 0
		for {
			e.pc--
			if e.op() == token.CLOSE {
				n++
			} else if e.op() == token.OPEN {
				if n--; n < 0 {
					break
				}
			}
		}
	}
}

func (e *Evaluator) Eval() {
	var t token.Token

	for {
		switch e.op() {
		case token.NEXT:
			e.next()
		case token.PREV:
			e.prev()
		case token.INC:
			e.inc()
		case token.DEC:
			e.dec()
		case token.READ:
			e.read()
		case token.WRITE:
			e.write()
		case token.OPEN:
			e.open()
		case token.CLOSE:
			e.close()
		case token.EOF:
			fmt.Fprintln(os.Stderr, "\n\n(END)")
			goto FINISH
		default:
			fmt.Println(t)
		}
		e.tick()
	}

FINISH:
}
