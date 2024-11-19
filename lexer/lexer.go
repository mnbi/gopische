package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	TT_ILLEGAL = "ILLEGAL"

	TT_LPAREN = "LPAREN"
	TT_RPAREN = "RPAREN"

	TT_SYMBOL = "SYMBOL"
	TT_NUMBER = "NUMBER"

	TT_STRING = "STRING"
)

type Lexer struct {
	Tokens []Token
	Cursor int
	Input  string
}

func NewLexer(input string) *Lexer {
	lexer := Lexer{Input: input}
	if ok := lexer.analyze(); !ok {
		return nil
	}
	return &lexer
}

func (l *Lexer) analyze() bool {
	scanner := newScanner(l.Input)

	var word string
	var ok bool

	for {
		word = scanner.nextWord()
		if word == "" {
			ok = true
			break
		}
		l.Tokens = append(l.Tokens, Token{Type: tokenType(word), Literal: word})
	}
	return ok
}

func tokenType(literal string) (tt TokenType) {
	runes := []rune(literal)
	switch runes[0] {
	case '(':
		tt = TT_LPAREN
	case ')':
		tt = TT_RPAREN
	case '"':
		tt = TT_STRING
	case '+', '-':
		if len(runes) > 1 && IsDigit(runes[1]) && readNumber(literal[1:]) {
			tt = TT_NUMBER
		} else {
			tt = TT_SYMBOL
		}
	default:
		if IsDigit(runes[0]) && readNumber(literal) {
			tt = TT_NUMBER
		} else {
			tt = TT_SYMBOL
		}
	}
	return
}

func readNumber(literal string) bool {
	runes := []rune(literal)
	for _, r := range runes {
		if !IsDigit(r) {
			return false
		}
	}
	return true
}
