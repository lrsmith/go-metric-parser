package wql

import (
	"testing"
)

//func TestEmptyInput(t *testing.T) {
//
//	l := lex("Test - Empty String", "")
//	l.run()
//	fmt.Println(l)
//}

//func TestTSExpression(t *testing.T) {
//	l := lex("Test - TS Expression", "ts(alpha.beta.omega)")
//	l.run()
//	fmt.Println(l)
//
//}

func TestTsExpression(t *testing.T) {
	input := "ts(alpha)"
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{TS, "ts"},
		{LPAREN, "("},
		{TS, "alpha"},
		{RPAREN, ")"},
	}
	l := Lex("TestNextToken", input)

	for i, tt := range tests {
		tok := l.nextToken()
		if tok.typ != tt.expectedType {
			t.Fatalf("tests[%d] = TokenType wrong. Expected=%q, got=%q,", i, tt.expectedType, tok.typ)
		}
		if tok.literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] = literal wrong. Expected=%q, got=%q,", i, tt.expectedLiteral, tok.literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := "(),"
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{LPAREN, "("},
		{RPAREN, ")"},
		{COMMA, ","},
	}
	l := Lex("TestNextToken", input)

	for i, tt := range tests {
		tok := l.nextToken()
		if tok.typ != tt.expectedType {
			t.Fatalf("tests[%d] = TokenType wrong. Expected=%q, got=%q,", i, tt.expectedType, tok.typ)
		}
		if tok.literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] = literal wrong. Expected=%q, got=%q,", i, tt.expectedLiteral, tok.literal)
		}
	}
}
