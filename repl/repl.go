package repl

import (
	"bufio"
	"fmt"
	"io"

	"pjserol/simple-interpreter/lexer"
	"pjserol/simple-interpreter/token"
)

const prompt = ">> "

// Start print the tokens
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, prompt)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
