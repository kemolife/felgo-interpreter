package repl

import (
	"fmt"
	"github.com/kemolife/felgo-interpreter/evaluator"
	"github.com/kemolife/felgo-interpreter/lexer"
	"github.com/kemolife/felgo-interpreter/object"
	"github.com/kemolife/felgo-interpreter/parser"
	"log"

	"github.com/openengineer/go-repl"
)

var helpMessage = `
Welcome to Flexible Executed Language base on Golang(FELGo) 0.0.1
Need more information please visit code examples folder in project 
exit              quit this program
`

var env = object.NewEnvironment()

type MyHandler struct {
	r *repl.Repl
}

func Start() {
	fmt.Println("REPL of FEL")
	fmt.Println("Type help for more information.")

	h := &MyHandler{}
	h.r = repl.NewRepl(h)

	if err := h.r.Loop(); err != nil {
		log.Fatal(err)
	}
}

func (h *MyHandler) Prompt() string {
	return "> "
}

func (h *MyHandler) Tab(_ string) string {
	return " "
}

// first return value is for stdout, second return value is for history
func (h *MyHandler) Eval(buffer string) string {
	if buffer == "" {
		return ""
	}

	switch buffer {
	case "help":
		return helpMessage
	case "exit":
		h.r.Quit()
		return ""
	default:
		return processProgram(buffer)
	}
}

func processProgram(buffer string) string {
	l := lexer.New(buffer)
	p := parser.New(l)
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			println(msg)
		}
		return ""
	}
	evaluated := evaluator.Eval(program, env)
	if evaluated == nil {
		return ""
	}

	return evaluated.Inspect()
}
