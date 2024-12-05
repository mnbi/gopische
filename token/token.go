// token/token.go

package token

import (
	"fmt"
)

type TokenType string

const (
	LPAREN  = "LPAREN"
	RPAREN  = "RPAREN"
	QUOTE   = "QUOTE"
	NUMBER  = "NUMBER"
	SYMBOL  = "SYMBOL"
	STRING  = "STRING"
	ILLEGAL = "ILLEGAL"
)

type Token struct {
	TokenType TokenType
	Literal   string
}

func NewToken(t TokenType, l string) (tk Token, err error) {
	tk = Token{TokenType: t, Literal: l}
	return
}

// Stringer interface for Token. The main purpose is to print token
// content in debugging.
func (t *Token) String() string {
	return fmt.Sprintf("[%s (%s)]", t.TokenType, t.Literal)
}
