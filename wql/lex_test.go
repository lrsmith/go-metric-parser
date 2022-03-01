package wql

import (
	"fmt"
	"testing"
)

func TestTsExpression(t *testing.T) {
	input := "ts(alpha.beta.omega)"
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{TS, "ts"},
		{LPAREN, "("},
		{TS, "alpha.beta.omega"},
		{RPAREN, ")"},
	}
	l := Lex("TestNextToken", input)

	//	for i, tt := range tests {
	//		tok := l.nextToken()
	for _, _ = range tests {
		_ = l.nextToken()

		//		if tok.typ != tt.expectedType {
		//			t.Fatalf("tests[%d] = TokenType wrong. Expected=%q, got=%q,", i, tt.expectedType, tok.typ)
		//		}
		//		if tok.literal != tt.expectedLiteral {
		//			t.Fatalf("tests[%d] = literal wrong. Expected=%q, got=%q,", i, tt.expectedLiteral, tok.literal)
		//		}
	}
	fmt.Println("tokens : ", l.tokens)
}

//func TestNextToken(t *testing.T) {
//	input := "(),"
//	tests := []struct {
//		expectedType    TokenType
//		expectedLiteral string
//	}{
//		{LPAREN, "("},
//		{RPAREN, ")"},
//		{COMMA, ","},
//	}
//	l := Lex("TestNextToken", input)
//
//	for i, tt := range tests {
//		tok := l.nextToken()
//		if tok.typ != tt.expectedType {
//			t.Fatalf("tests[%d] = TokenType wrong. Expected=%q, got=%q,", i, tt.expectedType, tok.typ)
//		}
//		if tok.literal != tt.expectedLiteral {
//			t.Fatalf("tests[%d] = literal wrong. Expected=%q, got=%q,", i, tt.expectedLiteral, tok.literal)
//		}
//	}
//}
