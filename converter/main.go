package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nownabe/brainfuck-go/config"
)

var bfConf = config.Config{
	NEXT:        ">",
	PREV:        "<",
	INC:         "+",
	DEC:         "-",
	READ:        ",",
	WRITE:       ".",
	OPEN:        "[",
	CLOSE:       "]",
	WHITESPACES: " \t\n\r",
}

func main() {
	var confPath string
	flag.StringVar(&confPath, "conf", "", "Config JSON file for your xxxxfuck.")
	flag.Parse()

	var conf config.Config
	confData, err := ioutil.ReadFile(confPath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read config file.")
		help()
	}

	if err = json.Unmarshal(confData, &conf); err != nil {
		fmt.Fprintln(os.Stderr, "cannot decode config file.")
		help()
	}

	sourceBlob, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read source file.")
		help()
	}

	source := string(sourceBlob)

	bfOps := bfConf.Ops()
	ops := conf.Ops()

	for i, op := range ops {
		source = strings.Replace(source, bfOps[i], op, -1)
	}

	fmt.Println(source)
}

func help() {
	fmt.Fprintln(os.Stderr, "Usage: converter [-conf config] source")
	os.Exit(1)
}
