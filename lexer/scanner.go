package lexer

import "fmt"

var debug = false

type scanner struct {
	runes  []rune
	cursor int
	length int
}

func newScanner(input string) *scanner {
	runes := []rune(input)
	return &scanner{runes: runes, cursor: 0, length: len(runes)}
}

func (s *scanner) peekRune(distance int) rune {
	peekPos := s.cursor + distance
	if peekPos < 0 || peekPos >= s.length {
		return 0
	}
	return s.runes[peekPos]
}

func (s *scanner) readRune() rune {
	r := s.peekRune(0)
	if r != 0 {
		s.cursor++
	}
	return r
}

func (s *scanner) unread(n int) {
	pos := s.cursor - n
	if pos < 0 {
		pos = 0
	}
	s.cursor = pos
}

func (s *scanner) nextWord() string {
	leftPos := s.cursor

	var r rune
	q := start
Loop:
	for {
		r = s.readRune()

		if debug {
			fmt.Printf("%c: ", r)
			fmt.Printf("%d -> ", q)
		}

		q = transition[edge{state: q, input: runeClass(r)}]

		if debug {
			fmt.Printf("%d\n", q)
		}

		switch q {
		case s0: // skip whitespaces
			leftPos++
		case s1, s2, s3:
			// nothing to do
		case s4:
			s.unread(1)
			break Loop
		case s99: // read a character illegally since
			// Something goes wrong, returns `false` to indicate such
			// condition and also returns the last word which already
			// read.
			return string(s.runes[leftPos:s.cursor])
		case s100: // accept
			break Loop
		}
	}
	return string(s.runes[leftPos:s.cursor])
}
