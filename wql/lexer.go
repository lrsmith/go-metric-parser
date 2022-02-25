package wql

import (
	"fmt"
	"unicode/utf8"
)

// Rob Pike

type TokenType int

type Token struct {
	typ     TokenType
	literal string
}

// lexer holds the state of the scanner.
type lexer struct {
	name   string // used only for error reports.
	input  string // the string being scanned.
	start  int    // start position of this item.
	pos    int    // current position in the input.
	width  int    // width of last rune read from input.
	tokens []Token
}

type stateFn func(*lexer) stateFn

const eof = -1
const (
	//Special Tokens
	ERROR TokenType = iota
	EOF

	//Identifiers and Literals
	IDENT
	METRIC

	// Delimiters
	COMMA
	LPAREN
	RPAREN

	//KeyWords
	TS
)

func (t Token) String() string {
	switch t.typ {
	case EOF:
		return "EOF"
	case ERROR:
		return t.literal
	}
	return fmt.Sprintf("%q", t.literal)
}

func Lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
	}

	return l
}

// isAlpha reports whether r is an alphabetic
func isAlpha(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

func lexIdentifier(l *lexer) Token {

	var tok Token
	start := (l.pos) - 1
	for isAlpha(l.next()) {
		// absorb
	}
	l.backup()
	tok.typ = TS
	tok.literal = l.input[start:l.pos]
	return tok
}

func (l *lexer) nextToken() Token {
	var tok Token

	switch r := l.next(); {
	case isAlpha(r):
		return lexIdentifier(l)
	case r == '(':
		tok = newToken(LPAREN, string(r))
	case r == ')':
		tok = newToken(RPAREN, string(r))
	case r == ',':
		tok = newToken(COMMA, string(r))
	}

	return tok
}

func newToken(tokenType TokenType, value string) Token {
	return Token{typ: tokenType, literal: value}
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = w
	l.pos += l.width
	return r
}

func lexState(l *lexer) stateFn {

	switch r := l.next(); {
	case r == eof:
		return nil
	default:
		fmt.Println(string(r))
		return lexState
	}

	return nil
}

func (l *lexer) run() {
	for state := lexState; state != nil; {
		state = state(l)
	}
}
