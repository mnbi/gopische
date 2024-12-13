package lexer

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/mnbi/gopische/lexer/internal/runeclass"
	"github.com/mnbi/gopische/lexer/internal/wscanner"
	"github.com/mnbi/gopische/scheme"
	"github.com/mnbi/gopische/token"
)

type Lexer struct {
	tokens []*token.Token
	cursor int
	input  []rune
}

// NewLexer accepts a string as a Scheme expression.  It analyzes the
// input and converts it into a sequence of tokens.  If any error
// ocurrs in the analysis, NewLexer returns nil.
func NewLexer(input string) *Lexer {
	runes := []rune(input)
	// The number of tokens is less than the number of runes in input.
	cap := len(runes)
	lexer := Lexer{tokens: make([]*token.Token, 0, cap), input: runes}
	if ok := lexer.analyze(); !ok {
		return nil
	}
	return &lexer
}

func (l *Lexer) Length() int {
	return len(l.tokens)
}

// Returns a next token to be read and true when tokens stil
// remain. When all tokens have been already read, returns 0 valuen
// and false.
func (l *Lexer) NextToken() (tk *token.Token, ok bool) {
	if l.cursor < len(l.tokens) {
		tk = l.tokens[l.cursor]
		ok = true
		l.cursor++
	}
	return tk, ok
}

func (l *Lexer) analyze() bool {
	wordScanner := wscanner.NewWordScanner(l.input)

	var leftPos, rightPos int
	var ok bool

	for {
		leftPos, rightPos = wordScanner.NextWord()
		if leftPos == rightPos { // eos
			ok = true
			break
		}
		if tk, err := l.createToken(leftPos, rightPos); err == nil {
			l.tokens = append(l.tokens, tk)
		} else {
			log.Printf("fail to create token: %s\n", err)
			return false
		}
	}
	return ok
}

func (l *Lexer) createToken(left int, right int) (tk *token.Token, err error) {
	var tt token.TokenType = token.ILLEGAL
	var lit string = string(l.input[left:right])
	var sobj scheme.Object = scheme.EmptyList

	length := right - left

	if length < 1 {
		tk = token.NewIllegalToken(lit)
		err = errors.New("empty literal")
		return
	}

	if length == 1 {
		switch l.input[left] {
		case '(':
			tt = token.LPAREN
		case ')':
			tt = token.RPAREN
		default:
			if runeclass.IsDigit(l.input[left]) {
				sobj, err = parseNumber(lit)
				if err == nil {
					tt = token.NUMBER
				}
			} else {
				sobj, err = scheme.NewSchemeObject(scheme.SYMBOL, lit)
				if err == nil {
					tt = token.SYMBOL
				}
			}
		}
		if err != nil {
			return &token.Token{}, err
		}

		tk = token.NewToken(tt, lit, sobj)
		return
	}

	var currRune, nextRune = l.input[left], l.input[left+1]
	switch currRune {
	case '(':
		if nextRune == ')' {
			tt = token.EMPTY_LIST
		} else {
			tt = token.ILLEGAL
			err = errors.New("weird literal")
		}
	case '"':
		lit := string(l.input[left+1 : right-1]) // eliminate quotation marks
		if sobj, err = scheme.NewSchemeObject(scheme.STRING, lit); err == nil {
			tt = token.STRING
		}
	case '#':
		if nextRune == 't' || nextRune == 'f' {
			if sobj, err = parseBoolean(lit); err == nil {
				tt = token.BOOLEAN
			}
		} else {
			tt = token.SYMBOL
		}
	case '+', '-':
		if runeclass.IsDigit(nextRune) || nextRune == '.' {
			sobj, err = parseNumber(lit)
			if err == nil {
				tt = token.NUMBER
			}
		} else {
			sobj, err = scheme.NewSchemeObject(scheme.SYMBOL, lit)
			if err == nil {
				tt = token.SYMBOL
			}
		}
	case '.':
		if runeclass.IsDigit(nextRune) {
			sobj, err = parseNumber(lit)
			if err == nil {
				tt = token.NUMBER
			}
		} else {
			sobj, err = scheme.NewSchemeObject(scheme.SYMBOL, lit)
			if err == nil {
				tt = token.SYMBOL
			}
		}
	default:
		if runeclass.IsDigit(currRune) {
			sobj, err = parseNumber(lit)
			if err == nil {
				tt = token.NUMBER
			}
		} else {
			sobj, err = scheme.NewSchemeObject(scheme.SYMBOL, lit)
			if err == nil {
				tt = token.SYMBOL
			}
		}
	}

	if err != nil {
		return token.NewIllegalToken(lit), err
	}

	tk = token.NewToken(tt, lit, sobj)
	return
}

func parseBoolean(lit string) (sobj scheme.Object, err error) {
	var bv bool

	switch lit {
	case "#t", "#true":
		bv = true
	case "#f", "#false":
		bv = false
	default:
		emsg := fmt.Sprintf("illegal boolean literal, %s", lit)
		err = errors.New(emsg)
		return
	}

	sobj, err = scheme.NewSchemeObject(scheme.BOOLEAN, bv)
	return
}

func parseNumber(lit string) (sobj scheme.Object, err error) {
	var iv64 int64
	var fv64 float64
	var cv128 complex128

	if iv64, err = strconv.ParseInt(lit, 0, 64); err == nil {
		if sobj, err = scheme.NewSchemeObject(scheme.NUMBER, iv64); err == nil {
			return
		}
	}

	if fv64, err = strconv.ParseFloat(lit, 64); err == nil {
		if sobj, err = scheme.NewSchemeObject(scheme.NUMBER, fv64); err == nil {
			return
		}
	}

	if cv128, err = strconv.ParseComplex(lit, 128); err == nil {
		if sobj, err = scheme.NewSchemeObject(scheme.NUMBER, cv128); err == nil {
			return
		}
	}

	return
}
