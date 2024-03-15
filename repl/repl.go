package repl

// Read Eval Print Loop

import (
	"MonkeyInterpreter/lexer"
	"MonkeyInterpreter/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

// read from the input source until encountering a newline, take
// the just read line and pass it to an instance of our lexer and finally print all the tokens the lexer
// gives us until we encounter EOF
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
