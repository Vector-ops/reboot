package repl

import (
	"bufio"
	"fmt"
	"io"
	"reboot/evaluator"
	"reboot/lexer"
	"reboot/object"
	"reboot/parser"
)

const PROMPT = ">> "

const REBOOT = `
██████  ███████ ██████   ██████   ██████  ████████ 
██   ██ ██      ██   ██ ██    ██ ██    ██    ██    
██████  █████   ██████  ██    ██ ██    ██    ██    
██   ██ ██      ██   ██ ██    ██ ██    ██    ██    
██   ██ ███████ ██████   ██████   ██████     ██    

`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	io.WriteString(out, REBOOT)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
