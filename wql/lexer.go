package wql

import (
	"fmt"
	"unicode/utf8"
)

type TokenType int

type Token struct {
	typ     TokenType
	literal string
	pos     int
}

// lexer holds the state of the scanner.
type lexer struct {
	name  string // used only for error reports.
	input string // the string being scanned.
	start int    // start position of this item.
	pos   int    // current position in the input.
	width int    // width of last rune read from input.

	parendepth int
	quotedepth int

	tokens []Token
}

const eof = -1
const (
	//Special Tokens
	ERROR TokenType = iota
	EOF

	//Identifiers and Literals
	IDENT
	LITERAL
	METRIC

	// Keywords
	AND
	OR
	NOT
	SOURCE
	TAG

	// Delimiters
	COMMA
	LPAREN
	RPAREN
	DQUOTE
	EQUAL

	// Functions
	TS
)

var TokenTypeStr = map[TokenType]string{
	IDENT:   "IDENT",
	LITERAL: "LITERAL",
	LPAREN:  "LPAREN",
	RPAREN:  "RPAREN",
	COMMA:   "COMMA",
	DQUOTE:  "DQUOTE",
	AND:     "AND",
	OR:      "OR",
	NOT:     "NOT",
	EQUAL:   "EQUAL",
	SOURCE:  "SOURCE",
	TAG:     "TAG",
	TS:      "TS",
}

var keywords = map[string]TokenType{
	"and":    AND,
	"or":     OR,
	"not":    NOT,
	"source": SOURCE,
	"tag":    TAG,
}

var functions = map[string]TokenType{
	"ts": TS,
}

func (t Token) String() string {
	switch t.typ {
	case EOF:
		return "EOF"
	case ERROR:
		return fmt.Sprintf("%d:%s:%s", t.pos, TokenTypeStr[t.typ], t.literal)
		//	return fmt.Sprintf("%d:%s", t.typ, t.literal)
		//	return t.literal
	}
	return fmt.Sprintf("%d:%s:%s", t.pos, TokenTypeStr[t.typ], t.literal)
}

func Lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
	}

	return l
}

//Valid characters are: a-z, A-Z, 0-9, hyphen ("-"), underscore ("_"), dot (".").
func isAlpha(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || r == '.' || r == '*' || r == '~'
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func (l *lexer) appendToken(tokenType TokenType, value string, pos int) {
	// Add test to fail if illegal token is found rather then continue
	tmpToken := newToken(tokenType, value, pos)
	//fmt.Println("Appending : ", l.start, pos)
	l.start = pos + 1
	l.tokens = append(l.tokens, tmpToken)
}

func isKeywordOrIdentifier(l *lexer) {
Loop:
	for {
		//if l.pos < len(l.input) {
		//	fmt.Println("2 :", l.pos, string(l.input[l.pos]), l.width)
		//}
		switch r := l.next(); {
		case isAlpha(r):
			// absorb.
		default:
			typ := LITERAL

			l.backup()
			word := l.input[l.start:l.pos]
			if kw, found := keywords[word]; found {
				typ = kw
			} else if fn, found := functions[word]; found {
				typ = fn
			}

			l.appendToken(typ, word, l.start)
			break Loop
		}
	}
}

func (l *lexer) tokenize() {
Loop:
	for {
		//if l.pos < len(l.input) {
		//	fmt.Println("1 :", l.pos, string(l.input[l.pos]), l.width)
		//}
		switch r := l.next(); {
		case isAlpha(r):
			l.backup()
			//			fmt.Println(l.pos, string(l.input[l.pos]), l.width)
			isKeywordOrIdentifier(l)
		case isSpace(r):
			l.start = l.pos
		case r == '=':
			l.appendToken(EQUAL, string(r), l.pos-1)
		case r == '"':
			l.appendToken(DQUOTE, string(r), l.pos-1)
			if l.quotedepth > 0 {
				l.quotedepth--
			} else {
				l.quotedepth++
			}
		case r == '(':
			l.appendToken(LPAREN, string(r), l.pos-1)
			l.parendepth++
		case r == ')':
			l.appendToken(RPAREN, string(r), l.pos-1)
			l.parendepth--
		case r == ',':
			l.appendToken(COMMA, string(r), l.pos-1)
		default:
			break Loop
		}

	}
}

func newToken(tokenType TokenType, value string, pos int) Token {
	return Token{typ: tokenType, literal: value, pos: pos}
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
