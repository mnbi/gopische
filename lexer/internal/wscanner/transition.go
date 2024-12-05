// lexer/internal/wscanner/transition.go

package wscanner

import (
	"github.com/mnbi/gopische/lexer/internal/runeclass"
)

type State int

const (
	Start   State = 0
	Illegal       = 99
	Accept        = 100
)

type Edge struct {
	State State
	Input runeclass.RuneClass
}

const (
	s1 State = iota + 1 // read quotation
	s2                  // read escape in quotation
	s3                  // read any characters other than space, delimiter, or quotation mark
	s4                  // read '(', ')', or ' ' after any other character
)

var transition = map[Edge]State{
	// start
	{State: Start, Input: runeclass.EOS}:         Accept,
	{State: Start, Input: runeclass.WHITE_SPACE}: Start,
	{State: Start, Input: runeclass.LEFT_PAREN}:  Accept,
	{State: Start, Input: runeclass.RIGHT_PAREN}: Accept,
	{State: Start, Input: runeclass.DOUBLE_QUOT}: s1,
	{State: Start, Input: runeclass.ESCAPE_CHAR}: Illegal,
	{State: Start, Input: runeclass.ANY_OTHER}:   s3,
	// s1: in quotation mark
	{State: s1, Input: runeclass.EOS}:         Illegal,
	{State: s1, Input: runeclass.WHITE_SPACE}: s1,
	{State: s1, Input: runeclass.LEFT_PAREN}:  s1,
	{State: s1, Input: runeclass.RIGHT_PAREN}: s1,
	{State: s1, Input: runeclass.DOUBLE_QUOT}: Accept,
	{State: s1, Input: runeclass.ESCAPE_CHAR}: s2,
	{State: s1, Input: runeclass.ANY_OTHER}:   s1,
	// s2: read an escapce character in quotation mark
	{State: s2, Input: runeclass.EOS}:         Illegal,
	{State: s2, Input: runeclass.WHITE_SPACE}: s1,
	{State: s2, Input: runeclass.LEFT_PAREN}:  s1,
	{State: s2, Input: runeclass.RIGHT_PAREN}: s1,
	{State: s2, Input: runeclass.DOUBLE_QUOT}: s1,
	{State: s2, Input: runeclass.ESCAPE_CHAR}: s1,
	{State: s2, Input: runeclass.ANY_OTHER}:   s1,
	// s3: read any other character (not in qoutation mark)
	{State: s3, Input: runeclass.EOS}:         Accept,
	{State: s3, Input: runeclass.WHITE_SPACE}: s4,
	{State: s3, Input: runeclass.LEFT_PAREN}:  s4,
	{State: s3, Input: runeclass.RIGHT_PAREN}: s4,
	{State: s3, Input: runeclass.DOUBLE_QUOT}: Illegal,
	{State: s3, Input: runeclass.ESCAPE_CHAR}: Illegal,
	{State: s3, Input: runeclass.ANY_OTHER}:   s3,
	// s4: read '(' or ')' after any other character
	{State: s4, Input: runeclass.EOS}:         Accept,
	{State: s4, Input: runeclass.WHITE_SPACE}: Illegal,
	{State: s4, Input: runeclass.LEFT_PAREN}:  Accept,
	{State: s4, Input: runeclass.RIGHT_PAREN}: Accept,
	{State: s4, Input: runeclass.DOUBLE_QUOT}: Illegal,
	{State: s4, Input: runeclass.ESCAPE_CHAR}: Illegal,
	{State: s4, Input: runeclass.ANY_OTHER}:   Illegal,
}
