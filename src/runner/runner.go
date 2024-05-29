package runner

import (
	"io"
	"os"
	"reboot/evaluator"
	"reboot/lexer"
	"reboot/object"
	"reboot/parser"
)

func Run(pathToFile string) {
	file, err := os.OpenFile(pathToFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	env := object.NewEnvironment()

	input, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(input))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParseErrors(os.Stdout, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
