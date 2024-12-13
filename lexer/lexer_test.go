// lexer/lexer_test.go

package lexer

import (
	"testing"

	"github.com/mnbi/gopische/token"
)

func TestCreateToken(t *testing.T) {
	tests := []struct {
		id           int
		testcase     string
		expectedType token.TokenType
	}{
		{1, "(", token.LPAREN},
		{2, ")", token.RPAREN},
		{3, "+", token.SYMBOL},
		{4, "car", token.SYMBOL},
		{5, "1", token.NUMBER},
		{6, "23", token.NUMBER},
		{7, "-4567", token.NUMBER},
		/*
			{8, "3.14", token.NUMBER},
			{9, "2.71828182845904523", token.NUMBER},
			{10, "1/2", token.NUMBER},
			{11, "1+2i", token.NUMBER},
			{12, "1-2i", token.NUMBER},
		*/
		{13, "\"Go\"", token.STRING}, // qouted string
		{14, "\"Scheme is a programming language.\"", token.STRING},
		{15, "++", token.SYMBOL},
		{16, "()", token.EMPTY_LIST},
		{17, "#t", token.BOOLEAN},
		{18, "#f", token.BOOLEAN},
		{19, "#true", token.BOOLEAN},
		{20, "#false", token.BOOLEAN},
	}

	for _, tc := range tests {
		l := NewLexer(tc.testcase)
		tk, err := l.createToken(0, len([]rune(tc.testcase)))
		if err != nil {
			t.Fatalf("tests[%d] - fail to create a token for %s\n", tc.id, tc.testcase)
		}

		if tk.TokenType != tc.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q",
				tc.id, tc.expectedType, tk.TokenType)
		}
	}
}
