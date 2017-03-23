package main

import (
	"fmt"

	"github.com/nownabe/brainfuck-go/config"
	"github.com/nownabe/brainfuck-go/evaluator"
)

var conf = config.Config{
	NEXT:  ">",
	PREV:  "<",
	INC:   "+",
	DEC:   "-",
	READ:  ",",
	WRITE: ".",
	OPEN:  "[",
	CLOSE: "]",
}

func main() {
	var input string
	fmt.Scan(&input)
	fmt.Println("Input:", input, "\n")

	e := evaluator.New(conf, input)
	e.Eval()
}
