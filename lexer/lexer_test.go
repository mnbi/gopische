// lexer/lexer_test.go

package lexer

import (
	"testing"
)

func TestTokenType(t *testing.T) {
	tests := []struct {
		id           int
		testcase     string
		expectedType TokenType
	}{
		{1, "(", TT_LPAREN},
		{2, ")", TT_RPAREN},
		{3, "+", TT_SYMBOL},
		{4, "car", TT_SYMBOL},
		{5, "1", TT_NUMBER},
		{6, "23", TT_NUMBER},
		/*
			{7, "-4567", TT_NUMBER},
			{8, "3.14", TT_NUMBER},
			{9, "2.71828182845904523", TT_NUMBER},
			{10, "1/2", TT_NUMBER},
			{11, "1+2i", TT_NUMBER},
			{12, "1-2i", TT_NUMBER},
		*/
		{13, "\"Go\"", TT_STRING}, // qouted string
		{14, "\"Scheme is a programming language.\"", TT_STRING},
		{15, "++", TT_SYMBOL},
	}

	for _, tc := range tests {
		tt := tokenType(tc.testcase)
		if tt != tc.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q",
				tc.id, tc.expectedType, tt)
		}
	}
}
