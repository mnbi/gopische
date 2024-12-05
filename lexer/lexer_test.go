// lexer/lexer_test.go

package lexer

import (
	"testing"

	"github.com/mnbi/gopische/token"
)

func TestTokenType(t *testing.T) {
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
		/*
			{7, "-4567", token.NUMBER},
			{8, "3.14", token.NUMBER},
			{9, "2.71828182845904523", token.NUMBER},
			{10, "1/2", token.NUMBER},
			{11, "1+2i", token.NUMBER},
			{12, "1-2i", token.NUMBER},
		*/
		{13, "\"Go\"", token.STRING}, // qouted string
		{14, "\"Scheme is a programming language.\"", token.STRING},
		{15, "++", token.SYMBOL},
	}

	for _, tc := range tests {
		tt := tokenType(tc.testcase)
		if tt != tc.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q",
				tc.id, tc.expectedType, tt)
		}
	}
}
