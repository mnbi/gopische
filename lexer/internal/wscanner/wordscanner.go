// lexer/internal/wscanner/wordscanner.go

package wscanner

import (
	"fmt"

	"github.com/mnbi/gopische/lexer/internal/runeclass"
)

var debug = false

// WordScanner scans the input by one rune at a time.
type WordScanner struct {
	runes  []rune
	cursor int
	length int
}

func NewWordScanner(input []rune) *WordScanner {
	return &WordScanner{runes: input, cursor: 0, length: len(input)}
}

// The position to be read at the next ReadRune call.
func (ws *WordScanner) Cursor() int {
	return ws.cursor
}

// PeekRune returns a rune at the distance position from the cursor.
// Unlike ReadRune, the cursor won't be moved.
func (ws *WordScanner) PeekRune(distance int) rune {
	peekPos := ws.cursor + distance
	if peekPos < 0 || peekPos >= ws.length {
		return 0
	}
	return ws.runes[peekPos]
}

// ReadRune returns a rune at the cursor position and move the cursor
// forward.  It returns 0 at the end of input (as EOS).
func (ws *WordScanner) ReadRune() rune {
	r := ws.PeekRune(0)
	if r != 0 {
		ws.cursor++
	}
	return r
}

// Unread moves the cursor backword.
func (ws *WordScanner) Unread(n int) {
	pos := ws.cursor - n
	if pos < 0 {
		pos = 0
	}
	ws.cursor = pos
}

func (ws *WordScanner) NextWord() (leftPos int, rightPos int) {
	leftPos = ws.Cursor()
	rightPos = ws.Cursor()

	var r rune
	var c runeclass.RuneClass

	q := Start
Loop:
	for {
		r = ws.ReadRune()
		c = runeClassify(r)

		if debug {
			fmt.Printf("%v (%c): ", c, r)
			fmt.Printf("%d -> ", q)
		}

		q = transition[Edge{State: q, Input: c}]

		if debug {
			fmt.Printf("%d\n", q)
		}

		switch q {
		case Start: // skip whitespaces
			leftPos++
		case s1:
		case s2:
		case s3:
		case s4:
		case s5:
		case s6:
		case Illegal: // read a character illegally since
			// Something goes wrong, returns `false` to indicate such
			// condition and also returns the last word which already
			// read.
			rightPos = ws.Cursor()
			return
		case Accept:
			if c != runeclass.EOS {
				ws.Unread(1)
			}
			rightPos = ws.Cursor()
			break Loop
		}
	}
	return
}

func (ws *WordScanner) SubRunes(left int, right int) []rune {
	if left < 0 || right > ws.length {
		return nil
	}
	return ws.runes[left:right]
}

// Rune to RuneClass mapping for a word
func runeClassify(r rune) (class runeclass.RuneClass) {
	switch r {
	case 0:
		class = runeclass.EOS
	case '(':
		class = runeclass.LEFT_PAREN
	case ')':
		class = runeclass.RIGHT_PAREN
	case '"':
		class = runeclass.DOUBLE_QUOT
	case '\\':
		class = runeclass.ESCAPE_CHAR
	default:
		if runeclass.IsWhitespace(r) {
			class = runeclass.WHITE_SPACE
		} else {
			class = runeclass.ANY_OTHER
		}
	}
	return
}
