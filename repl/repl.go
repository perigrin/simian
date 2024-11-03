package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/perigrin/simian/lexer"
	"github.com/perigrin/simian/token"
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
		l := lexer.New([]byte(line))

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Printf("%v\n", t)
		}
	}
}
