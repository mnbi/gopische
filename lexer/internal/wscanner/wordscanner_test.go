// lexer/internal/wscanner/wordscanner_test.go

package wscanner

import (
	"testing"
)

func cmpWords(expected []string, actual []string) bool {
	if len(expected) != len(actual) {
		return false
	}

	for pos, word := range expected {
		if word != actual[pos] {
			return false
		}
	}
	return true
}

func readAllWords(s *WordScanner) (words []string) {
	for {
		l, r := s.NextWord()
		word := string(s.SubRunes(l, r))
		if word == "" {
			break
		}
		words = append(words, word)
	}
	return
}

func TestNextWord(t *testing.T) {
	tests := []struct {
		id       int
		testcase string
		expected []string
	}{
		{1, "(", []string{"("}},
		{2, ")", []string{")"}},
		{3, "()", []string{"()"}},
		{4, "1", []string{"1"}},
		{5, "12", []string{"12"}},
		{6, "123", []string{"123"}},
		{10, "(+ 1 2)", []string{"(", "+", "1", "2", ")"}},
		{11, "(+ 10 234 (- 56 7) (* 8 9))",
			[]string{"(", "+", "10", "234", "(", "-", "56", "7", ")", "(", "*", "8", "9", ")", ")"}},
		{12, "( )", []string{"( )"}},
	}

	for _, tc := range tests {
		s := NewWordScanner([]rune(tc.testcase))
		if s == nil {
			t.Fatalf("tests[%d] - fail to instantiate a scanner for \"%s\"",
				tc.id, tc.testcase)
		}
		actual := readAllWords(s)
		if !cmpWords(tc.expected, actual) {
			t.Fatalf("tests[%d] - expected=%q, got=%q", tc.id,
				tc.expected, actual)
		}
	}
}
