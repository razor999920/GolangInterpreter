package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) { 
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT) 
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() { 
			fmt.Printf("%+v\n", tok)
		} 
	}
}
