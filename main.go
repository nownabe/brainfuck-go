package main

import (
	"flag"
	"fmt"
	"io/ioutil"

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
	flag.Parse()
	source, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		panic(err)
	}
	fmt.Println("Input:", string(source), "\n")

	e := evaluator.New(conf, string(source))
	e.Eval()
}
