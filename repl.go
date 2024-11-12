package gopische

import (
	"bufio"
	"fmt"
	"os"
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
	prompt(writer)

	for scanner.Scan() {
		firstLine := scanner.Text()

		print(eval(read(firstLine)))

		prompt(writer)
	}

	farewell(writer)

	return 0
}

func read(input string) string {
	return input
}

func eval(exp string) string {
	return exp
}

func print(value string) {
	fmt.Printf("%s\n", value)
}
