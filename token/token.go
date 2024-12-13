// token/token.go

package token

import (
	"fmt"

	"github.com/mnbi/gopische/scheme"
)

type TokenType string

const (
	LPAREN     = "LPAREN"
	RPAREN     = "RPAREN"
	QUOTE      = "QUOTE"
	EMPTY_LIST = "EMPTY_LIST"
	BOOLEAN    = "BOOLEAN"
	NUMBER     = "NUMBER"
	STRING     = "STRING"
	SYMBOL     = "SYMBOL"
	ILLEGAL    = "ILLEGAL"
)

type Token struct {
	TokenType TokenType
	Literal   string
	Value     scheme.Object
}

func NewIllegalToken(lit string) *Token {
	return NewToken(ILLEGAL, lit, scheme.EmptyList)
}

func NewToken(tt TokenType, lit string, val scheme.Object) *Token {
	return &Token{TokenType: tt, Literal: lit, Value: val}
}

// Stringer interface for Token. The main purpose is to print token
// content in debugging.
func (t *Token) String() string {
	return fmt.Sprintf("[type:%s, literal:%s, value:%s]", t.TokenType, t.Literal, t.Value)
}
