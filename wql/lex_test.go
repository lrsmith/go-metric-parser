package wql

import (
	"fmt"
	"testing"
)

func TestTsExpression(t *testing.T) {
	input := "ts(alpha.beta)"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestBasicTsExpression(t *testing.T) {
	input := "ts(alpha)"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestNextToken(t *testing.T) {
	input := "(),"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
	//if tok.typ != tt.expectedType {
	//	t.Fatalf("tests[%d] = TokenType wrong. Expected=%q, got=%q,", i, tt.expectedType, tok.typ)
	//}
	//		if tok.literal != tt.expectedLiteral {
	//			t.Fatalf("tests[%d] = literal wrong. Expected=%q, got=%q,", i, tt.expectedLiteral, tok.literal)
	//		}
}
