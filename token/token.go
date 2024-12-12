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
	NUMBER     = "NUMBER"
	SYMBOL     = "SYMBOL"
	STRING     = "STRING"
	EMPTY_LIST = "EMPTY_LIST"
	ILLEGAL    = "ILLEGAL"
)

type Token struct {
	TokenType TokenType
	Literal   string
	Value     scheme.Object
}

func NewToken(t TokenType, l string, v scheme.Object) (tk *Token, err error) {
	tk = &Token{TokenType: t, Literal: l, Value: v}
	return
}

// Stringer interface for Token. The main purpose is to print token
// content in debugging.
func (t *Token) String() string {
	return fmt.Sprintf("[type:%s, literal:%s, value:%s]", t.TokenType, t.Literal, t.Value)
}
