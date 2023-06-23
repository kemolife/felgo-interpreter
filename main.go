package main

import (
	"fmt"
	"github.com/kemolife/felgo-interpreter/evaluator"
	"github.com/kemolife/felgo-interpreter/lexer"
	"github.com/kemolife/felgo-interpreter/object"
	"github.com/kemolife/felgo-interpreter/parser"
	"github.com/kemolife/felgo-interpreter/repl"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		repl.Start()
	}

	env := object.NewEnvironment()
	data, err := ioutil.ReadFile(os.Args[1])
	path, _ := os.Getwd()
	if err != nil {
		fmt.Printf("can't open file: %s/%s . No such file or directory", path, os.Args[1])
		return
	}

	fileExtension := filepath.Ext(os.Args[1])
	if fileExtension != ".fel" {
		fmt.Println("File extension isn't equal to .fel")
		return
	}

	l := lexer.New(string(data))
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}

func printParserErrors(errors []string) {
	fmt.Println("Parser errors:")
	for _, msg := range errors {
		fmt.Println("\t" + msg)
	}
}
