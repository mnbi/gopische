// lexer/internal/wscanner/wordscanner.go

package wscanner

import (
	"fmt"

	"github.com/mnbi/gopische/lexer/internal/runeclass"
)

var debug = false

type WordScanner struct {
	runes  []rune
	cursor int
	length int
}

func NewWordScanner(input string) *WordScanner {
	runes := []rune(input)
	return &WordScanner{runes: runes, cursor: 0, length: len(runes)}
}

func (ws *WordScanner) Cursor() int {
	return ws.cursor
}

func (ws *WordScanner) PeekRune(distance int) rune {
	peekPos := ws.cursor + distance
	if peekPos < 0 || peekPos >= ws.length {
		return 0
	}
	return ws.runes[peekPos]
}

func (ws *WordScanner) ReadRune() rune {
	r := ws.PeekRune(0)
	if r != 0 {
		ws.cursor++
	}
	return r
}

func (ws *WordScanner) Unread(n int) {
	pos := ws.cursor - n
	if pos < 0 {
		pos = 0
	}
	ws.cursor = pos
}

func (ws *WordScanner) SubRunes(left int, right int) []rune {
	if left < 0 || right > ws.length {
		return nil
	}
	return ws.runes[left:right]
}

func (ws *WordScanner) NextWord() string {
	leftPos := ws.Cursor()

	var r rune
	q := Start
Loop:
	for {
		r = ws.ReadRune()

		if debug {
			fmt.Printf("%c: ", r)
			fmt.Printf("%d -> ", q)
		}

		q = transition[Edge{State: q, Input: runeClassify(r)}]

		if debug {
			fmt.Printf("%d\n", q)
		}

		switch q {
		case Start: // skip whitespaces
			leftPos++
		case s1, s2, s3:
			// nothing to do
		case s4:
			ws.Unread(1)
			break Loop
		case Illegal: // read a character illegally since
			// Something goes wrong, returns `false` to indicate such
			// condition and also returns the last word which already
			// read.
			return string(ws.SubRunes(leftPos, ws.Cursor()))
		case Accept: // accept
			break Loop
		}
	}
	return string(ws.SubRunes(leftPos, ws.Cursor()))
}

// rune to RuneClass mapping for a word
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
