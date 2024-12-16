package gopische

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mnbi/gopische/lexer"
	"github.com/mnbi/gopische/scheme"
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

	var sexp scheme.Object
	var err error

	for {
		prompt(writer)

		if !scanner.Scan() {
			break
		}
		firstLine := scanner.Text()
		if sexp, err = read(firstLine); err != nil {
			log.Print(err)
			continue
		}
		print(writer, eval(sexp))
	}

	farewell(writer)

	return 0
}

func read(input string) (sexp scheme.Object, err error) {
	l := lexer.NewLexer(input)
	if l == nil {
		emsg := fmt.Sprintf("fail to analyze lexically: \"%s\"", input)
		return scheme.EmptyList, errors.New(emsg)
	}
	sexp, err = parse(l)
	return
}

func eval(sexp scheme.Object) scheme.Object {
	return sexp
}

func print(writer *bufio.Writer, value scheme.Object) {
	answerLine := fmt.Sprintf("%s\n", value.String())
	writeString(writer, answerLine)
}

func parse(l *lexer.Lexer) (sexp scheme.Object, err error) {
	for tk, ok := l.NextToken(); ok; tk, ok = l.NextToken() {
		switch tk.TokenType {
		case token.NUMBER:
			sexp = tk.Value
		case token.STRING:
			sexp = tk.Value
		case token.EMPTY_LIST:
			sexp = tk.Value
		case token.BOOLEAN:
			sexp = tk.Value
		default:
			emsg := fmt.Sprintf("fail to parse: %s", tk)
			err = errors.New(emsg)
		}
	}
	return
}
