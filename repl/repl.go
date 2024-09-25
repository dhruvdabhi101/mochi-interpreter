package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dhruvdabhi101/interpreter/evaluator"
	"github.com/dhruvdabhi101/interpreter/lexer"
	"github.com/dhruvdabhi101/interpreter/parser"
)

const PROMPT = ">> "

const MOCHI_ASCII = `__,__
  ⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠔⠶⠒⠉⠈⠸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠪⣦⢄⣀⡠⠁⠀⠀⠀⠀⠀⠀⠀⢀⣀⣠⣤⣤⣤⣤⣤⣄⣀⣀⣀⣀⣀⣀⣀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠈⠉⠀⠀⠀⣰⣶⣶⣦⠶⠛⠋⠉⠀⠀⠀⠀⠀⠀⠀⠉⠉⢷⡔⠒⠚⢽⠃⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣰⣿⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⢅⢰⣾⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣀⡴⠞⠛⠉⣿⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣧⠀⠀⠀⠀⠀
⠀⣀⣀⣤⣤⡞⠋⠀⠀⠀⢠⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⡇⠀⠀⠀⠀
⢸⡏⠉⣴⠏⠀⠀⠀⠀⠀⢸⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠀
⠈⣧⢰⠏⠀⠀⠀⠀⠀⠀⢸⡆⠀⠀⠀⠀⠀⠀⠀⠀⠰⠯⠥⠠⠒⠄⠀⠀⠀⠀⠀⠀⢠⠀⣿⠀⠀⠀⠀
⠀⠈⣿⠀⠀⠀⠀⠀⠀⠀⠈⡧⢀⢻⠿⠀⠲⡟⣞⠀⠀⠀⠀⠈⠀⠁⠀⠀⠀⠀⠀⢀⠆⣰⠇⠀⠀⠀⠀
⠀⠀⣿⠀⠀⠀⠀⠀⠀⠀⠀⣧⡀⠃⠀⠀⠀⠱⣼⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠠⣂⡴⠋⠀⣀⡀⠀⠀
⠀⠀⢹⡄⠀⠀⠀⠀⠀⠀⠀⠹⣜⢄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠒⠒⠿⡻⢦⣄⣰⠏⣿⠀⠀
⠀⠀⠀⢿⡢⡀⠀⠀⠀⠀⠀⠀⠙⠳⢮⣥⣤⣤⠶⠖⠒⠛⠓⠀⠀⠀⠀⠀⠀⠀⠀⠀⠑⢌⢻⣴⠏⠀⠀
⠀⠀⠀⠀⠻⣮⣒⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣧⣤⣀⣀⣀⣤⡴⠖⠛⢻⡆⠀⠀⠀⠀⠀⠀⢣⢻⡄⠀⠀
⠀⠀⠀⠀⠀⠀⠉⠛⠒⠶⠶⡶⢶⠛⠛⠁⠀⠀⠀⠀⠀⠀⠀⢀⣀⣤⠞⠁⠀⠀⠀⠀⠀⠀⠈⢜⢧⣄⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠃⠇⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⠉⢻⠀⠀⠀⠀⠀⠀⠀⢀⣀⠀⠀⠉⠈⣷
⠀⠀⠀⠀⠀⠀⠀⣼⠟⠷⣿⣸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠲⠶⢶⣶⠶⠶⢛⣻⠏⠙⠛⠛⠛⠁
⠀⠀⠀⠀⠀⠀⠀⠈⠷⣤⣀⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀⠀⠀⠉⠛⠓⠚⠋⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⣟⡂⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢹⡟⡟⢻⡟⠛⢻⡄⠀⠀⣸⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⡄⠀⠀⠀⠈⠷⠧⠾⠀⠀⠀⠻⣦⡴⠏⠀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠁⠀⠀⠀⠀⠈⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
  `

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MOCHI_ASCII)
	io.WriteString(out, "Woops! We ran into some mochi errors here!\n")
	io.WriteString(out, "  parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+"\x1B[31m"+msg+"\x1B[0m"+"\n")
	}
}
