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
	s1 State = iota + 1 // read '(' at start
	s2                  // read ')' at start
	s3                  // read a string
	s4                  // read a symbol
	s5                  // read an escape character in a string
	s6                  // read the empty list
)

var transition = map[Edge]State{
	// start
	{State: Start, Input: runeclass.EOS}:         Accept,
	{State: Start, Input: runeclass.WHITE_SPACE}: Start,
	{State: Start, Input: runeclass.LEFT_PAREN}:  s1,
	{State: Start, Input: runeclass.RIGHT_PAREN}: s2,
	{State: Start, Input: runeclass.DOUBLE_QUOT}: s3,
	{State: Start, Input: runeclass.ESCAPE_CHAR}: Illegal,
	{State: Start, Input: runeclass.ANY_OTHER}:   s4,
	// s1: read '(' at start
	{State: s1, Input: runeclass.EOS}:         Accept,
	{State: s1, Input: runeclass.WHITE_SPACE}: s1,
	{State: s1, Input: runeclass.LEFT_PAREN}:  Accept,
	{State: s1, Input: runeclass.RIGHT_PAREN}: s6,
	{State: s1, Input: runeclass.DOUBLE_QUOT}: Accept,
	{State: s1, Input: runeclass.ESCAPE_CHAR}: Illegal,
	{State: s1, Input: runeclass.ANY_OTHER}:   Accept,
	// s2: read ')' at start
	{State: s2, Input: runeclass.EOS}:         Accept,
	{State: s2, Input: runeclass.WHITE_SPACE}: Accept,
	{State: s2, Input: runeclass.LEFT_PAREN}:  Accept,
	{State: s2, Input: runeclass.RIGHT_PAREN}: Accept,
	{State: s2, Input: runeclass.DOUBLE_QUOT}: Accept,
	{State: s2, Input: runeclass.ESCAPE_CHAR}: Illegal,
	{State: s2, Input: runeclass.ANY_OTHER}:   Accept,
	// s3: read a string
	{State: s3, Input: runeclass.EOS}:         Illegal,
	{State: s3, Input: runeclass.WHITE_SPACE}: s3,
	{State: s3, Input: runeclass.LEFT_PAREN}:  s3,
	{State: s3, Input: runeclass.RIGHT_PAREN}: s3,
	{State: s3, Input: runeclass.DOUBLE_QUOT}: Accept,
	{State: s3, Input: runeclass.ESCAPE_CHAR}: s5,
	{State: s3, Input: runeclass.ANY_OTHER}:   s3,
	// s4: read a symbol
	{State: s4, Input: runeclass.EOS}:         Accept,
	{State: s4, Input: runeclass.WHITE_SPACE}: Accept,
	{State: s4, Input: runeclass.LEFT_PAREN}:  Accept,
	{State: s4, Input: runeclass.RIGHT_PAREN}: Accept,
	{State: s4, Input: runeclass.DOUBLE_QUOT}: Accept,
	{State: s4, Input: runeclass.ESCAPE_CHAR}: Illegal,
	{State: s4, Input: runeclass.ANY_OTHER}:   s4,
	// s5: read an escapce character in a string
	{State: s5, Input: runeclass.EOS}:         Illegal,
	{State: s5, Input: runeclass.WHITE_SPACE}: s3,
	{State: s5, Input: runeclass.LEFT_PAREN}:  s3,
	{State: s5, Input: runeclass.RIGHT_PAREN}: s3,
	{State: s5, Input: runeclass.DOUBLE_QUOT}: s3,
	{State: s5, Input: runeclass.ESCAPE_CHAR}: s3,
	{State: s5, Input: runeclass.ANY_OTHER}:   s3,
	// s6: read the empyt list
	{State: s6, Input: runeclass.EOS}:         Accept,
	{State: s6, Input: runeclass.WHITE_SPACE}: Accept,
	{State: s6, Input: runeclass.LEFT_PAREN}:  Accept,
	{State: s6, Input: runeclass.RIGHT_PAREN}: Accept,
	{State: s6, Input: runeclass.DOUBLE_QUOT}: Accept,
	{State: s6, Input: runeclass.ESCAPE_CHAR}: Illegal,
	{State: s6, Input: runeclass.ANY_OTHER}:   Accept,
}
