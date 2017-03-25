package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nownabe/brainfuck-go/config"
	"github.com/nownabe/brainfuck-go/evaluator"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "conf", "", "Config JSON file for your xxxxfuck.")
	flag.Parse()

	var conf config.Config
	if confPath == "" {
		// Standard brainfuck
		conf = config.Config{
			NEXT:  ">",
			PREV:  "<",
			INC:   "+",
			DEC:   "-",
			READ:  ",",
			WRITE: ".",
			OPEN:  "[",
			CLOSE: "]",
		}
	} else {
		// Customize brainfuck
		confData, err := ioutil.ReadFile(confPath)

		if err != nil {
			fmt.Fprintln(os.Stderr, "cannot read config file.")
			help()
		}

		if err = json.Unmarshal(confData, &conf); err != nil {
			fmt.Fprintln(os.Stderr, "cannot decode config file.")
			help()
		}
	}

	source, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read source file.")
		help()
	}
	fmt.Fprintln(os.Stderr, "Input:", string(source), "\n")

	e := evaluator.New(conf, string(source))
	e.Eval()
}

func help() {
	fmt.Fprintln(os.Stderr, "Usage: brainfuck-go [-conf config] source")
	os.Exit(1)
}
