package lexer

import (
	"github.com/mnbi/gopische/lexer/internal/runeclass"
	"github.com/mnbi/gopische/lexer/internal/wscanner"
	"github.com/mnbi/gopische/token"
)

type Lexer struct {
	tokens []token.Token
	cursor int
	input  string
}

func NewLexer(input string) *Lexer {
	// The number of tokens is less than the number of runes in input.
	cap := len([]rune(input))
	lexer := Lexer{tokens: make([]token.Token, 0, cap), input: input}
	if ok := lexer.analyze(); !ok {
		return nil
	}
	return &lexer
}

func (l *Lexer) Length() int {
	return len(l.tokens)
}

// Returns a next token to be read and true when tokens stil
// remain. When all tokens have been already read, returns 0 valuen
// and false.
func (l *Lexer) NextToken() (tk token.Token, ok bool) {
	if l.cursor < len(l.tokens) {
		tk = l.tokens[l.cursor]
		ok = true
		l.cursor++
	}
	return tk, ok
}

func (l *Lexer) analyze() bool {
	wordScanner := wscanner.NewWordScanner(l.input)

	var word string
	var ok bool

	for {
		word = wordScanner.NextWord()
		if word == "" { // eos
			ok = true
			break
		}
		if tk, err := token.NewToken(tokenType(word), word); err == nil {
			l.tokens = append(l.tokens, tk)
		}
	}
	return ok
}

func tokenType(literal string) (tt token.TokenType) {
	runes := []rune(literal)
	if len(runes) < 1 {
		tt = token.ILLEGAL
		return
	}

	switch runes[0] {
	case '(':
		tt = token.LPAREN
	case ')':
		tt = token.RPAREN
	case '"':
		tt = token.STRING
	case '+', '-':
		tt = token.SYMBOL
		if len(runes) > 1 && (runeclass.IsDigit(runes[1]) || runes[1] == '.') {
			tt = token.NUMBER
		}
	case '.':
		if len(runes) == 1 {
			tt = token.ILLEGAL
		}
		if runeclass.IsDigit(runes[1]) {
			tt = token.NUMBER
		}
	default:
		if runeclass.IsDigit(runes[0]) {
			tt = token.NUMBER
		} else {
			tt = token.SYMBOL
		}
	}
	return
}
