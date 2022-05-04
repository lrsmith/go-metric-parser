package wql

import (
	"fmt"
	"testing"
)

func TestTsExpressionWithTag(t *testing.T) {
	input := "ts(alpha, tag=mytag)"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestTsExpressionUnderscore(t *testing.T) {
	input := "(ts(\"zookeeper.zk_open_file_descriptor_count.gauge\", source=\"on58-hi.aue1t.internal\")"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestTsExpressionWildUnder(t *testing.T) {
	input := "ts(\"interface.*.if_packets.rx\" ,source=\"monapp03b.aue1d.internal\")"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestTsExpressionWithSourcet(t *testing.T) {
	input := "ts(alpha, source=mysource)"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestTsExpressionWithAndNot(t *testing.T) {
	input := "ts(alpha and not omega)"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}
func TestTsExpressionWithAnd(t *testing.T) {
	input := "ts(alpha and omega)"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestTsExpressionWithTilda(t *testing.T) {
	input := "ts(\"~alpha.beta.omega\")"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestTsExpressionWithWildcard(t *testing.T) {
	input := "ts(\"alpha.*.omega\")"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)

}

func TestTsExpressionWithQuotes(t *testing.T) {
	input := "ts(\"alpha.beta.omega\")"
	l := Lex("TestNextToken", input)

	l.tokenize()
	fmt.Println("tokens : ", l.tokens)
}

func TestBasicTsExpressionWithQuotes(t *testing.T) {
	input := "ts(\"alpha\")"
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
