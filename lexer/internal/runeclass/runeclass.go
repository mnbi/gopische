// lexer/internal/runeclass/runeclass.go

package runeclass

import (
	"unicode"
)

type RuneClass string

const (
	// end of string
	EOS = "EOS"
	// word scan
	WHITE_SPACE = "WHITE_SPACE"
	LEFT_PAREN  = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
	DOUBLE_QUOT = "DOUBLE_QUOT"
	ESCAPE_CHAR = "ESCAPE_CHAR"
	// number scan
	SIGN          = "SIGN"
	DIGIT_ZERO    = "DIGIT_ZERO"
	DIGIT_NONZERO = "DIGIT_NONZERO"
	POINT         = "POINT"
	// general scan
	ANY_OTHER = "ANY_OTHER"
)

func IsWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

// Extended identifier characters, see 2.1 in R7RS (or R5RS).
var extendedIdChars = [...]rune{'!', '$', '%', '&', '*', '+', '-', '.', '/', ':', '<', '=', '>', '?', '@', '^', '_', '~'}

var digits = [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func IsDigit(r rune) (found bool) {
	for _, digit := range digits {
		if r == digit {
			found = true
			break
		}
	}
	return
}

func IsZeroDigit(r rune) bool {
	return r == '0'
}

func IsNonZeroDigit(r rune) (found bool) {
	for _, digit := range digits[1:] {
		if r == digit {
			found = true
			break
		}
	}
	return
}
