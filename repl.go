package gopische

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mnbi/gopische/lexer"
	"github.com/mnbi/gopische/token"
)

func writeString(writer *bufio.Writer, str string) {
	_, _ = writer.WriteString(str)
	writer.Flush()
}

func welcome(writer *bufio.Writer) {
	msg := fmt.Sprintf("Welcome to %s - %s (%s)\n", name, version, revision)
	writeString(writer, msg)
}

func prompt(writer *bufio.Writer) {
	header := fmt.Sprintf("%s > ", name)
	writeString(writer, header)
}

func farewell(writer *bufio.Writer) {
	msg := fmt.Sprintf("\nBye!\n")
	writeString(writer, msg)
}

func Repl() int {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	welcome(writer)

	var exp string
	var ok bool

	for {
		prompt(writer)

		if !scanner.Scan() {
			break
		}
		firstLine := scanner.Text()
		if exp, ok = read(firstLine); !ok {
			continue
		}

		print(writer, eval(exp))
	}

	farewell(writer)

	return 0
}

func read(input string) (string, bool) {
	l := lexer.NewLexer(input)
	if l == nil {
		log.Printf("fail to analyze lexically: \"%s\"", input)
		return "", false
	}
	return parse(l), true
}

func eval(exp string) string {
	return exp
}

func print(writer *bufio.Writer, value string) {
	answerLine := fmt.Sprintf("%s\n", value)
	writeString(writer, answerLine)
}

func parse(l *lexer.Lexer) string {
	var ok bool
	var tk *token.Token

	tokenStrings := make([]string, 0, l.Length())
	for {
		if tk, ok = l.NextToken(); !ok { // no more tokens
			break
		}
		tokenStrings = append(tokenStrings, tk.String())
	}

	return strings.Join(tokenStrings, ",\n")
}
