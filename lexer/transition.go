package lexer

import (
	"unicode"
)

type RuneClass string

const (
	RC_EOS         = "EOS"
	RC_WHITE_SPACE = "WHITE_SPACE"
	RC_LEFT_PAREN  = "LEFT_PAREN"
	RC_RIGHT_PAREN = "RIGHT_PAREN"
	RC_DOUBLE_QUOT = "DOUBLE_QUOT"
	RC_ESCAPE_CHAR = "ESCAPE_CHAR"
	RC_ANY_OTHER   = "ANY_OTHER"
)

func runeClass(r rune) (class RuneClass) {
	switch r {
	case 0:
		class = RC_EOS
	case '(':
		class = RC_LEFT_PAREN
	case ')':
		class = RC_RIGHT_PAREN
	case '"':
		class = RC_DOUBLE_QUOT
	case '\\':
		class = RC_ESCAPE_CHAR
	default:
		if IsWhitespace(r) {
			class = RC_WHITE_SPACE
		} else {
			class = RC_ANY_OTHER
		}
	}
	return
}

func IsWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

var digits = [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func IsDigit(r rune) bool {
	found := false
	for _, digit := range digits {
		if r == digit {
			found = true
			break
		}
	}
	return found
}

type state int

const (
	s0   state = iota // start
	s1                // read quotation
	s2                // read escape in quotation
	s3                // read any characters other than space, delimiter, or quotation mark
	s4                // read '(', ')', or ' ' after any other character
	s99  = 99         // illegal; read a wrong rune in a wrong state
	s100 = 100        // accept
)

const (
	start   = s0
	illegal = s99
	accept  = s100
)

type edge struct {
	state state
	input RuneClass
}

var transition = map[edge]state{
	// s0: start
	{state: s0, input: RC_EOS}:         s100,
	{state: s0, input: RC_WHITE_SPACE}: s0,
	{state: s0, input: RC_LEFT_PAREN}:  s100,
	{state: s0, input: RC_RIGHT_PAREN}: s100,
	{state: s0, input: RC_DOUBLE_QUOT}: s1,
	{state: s0, input: RC_ESCAPE_CHAR}: s99,
	{state: s0, input: RC_ANY_OTHER}:   s3,
	// s1: in quotation mark
	{state: s1, input: RC_EOS}:         s99,
	{state: s1, input: RC_WHITE_SPACE}: s1,
	{state: s1, input: RC_LEFT_PAREN}:  s1,
	{state: s1, input: RC_RIGHT_PAREN}: s1,
	{state: s1, input: RC_DOUBLE_QUOT}: s100,
	{state: s1, input: RC_ESCAPE_CHAR}: s2,
	{state: s1, input: RC_ANY_OTHER}:   s1,
	// s2: read an escapce character in quotation mark
	{state: s2, input: RC_EOS}:         s99,
	{state: s2, input: RC_WHITE_SPACE}: s1,
	{state: s2, input: RC_LEFT_PAREN}:  s1,
	{state: s2, input: RC_RIGHT_PAREN}: s1,
	{state: s2, input: RC_DOUBLE_QUOT}: s1,
	{state: s2, input: RC_ESCAPE_CHAR}: s1,
	{state: s2, input: RC_ANY_OTHER}:   s1,
	// s3: read any other character (not in qoutation mark)
	{state: s3, input: RC_EOS}:         s100,
	{state: s3, input: RC_WHITE_SPACE}: s4,
	{state: s3, input: RC_LEFT_PAREN}:  s4,
	{state: s3, input: RC_RIGHT_PAREN}: s4,
	{state: s3, input: RC_DOUBLE_QUOT}: s99,
	{state: s3, input: RC_ESCAPE_CHAR}: s99,
	{state: s3, input: RC_ANY_OTHER}:   s3,
	// s4: read '(' or ')' after any other character
	{state: s4, input: RC_EOS}:         s100,
	{state: s4, input: RC_WHITE_SPACE}: s99,
	{state: s4, input: RC_LEFT_PAREN}:  s100,
	{state: s4, input: RC_RIGHT_PAREN}: s100,
	{state: s4, input: RC_DOUBLE_QUOT}: s99,
	{state: s4, input: RC_ESCAPE_CHAR}: s99,
	{state: s4, input: RC_ANY_OTHER}:   s99,
}
