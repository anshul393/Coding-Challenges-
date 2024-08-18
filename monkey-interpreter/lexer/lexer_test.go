package lexer

import (
	"testing"

	"github.com/anshul393/coding-challenges/monkey-interpreter/token"
)

func TestLexer(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, test := range tests {
		tkn := l.NextToken()
		if test.expectedLiteral != tkn.Literal {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected=%q, got=%q",
				i, test.expectedLiteral, tkn.Literal)
		}

		if test.expectedType != tkn.Type {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test.expectedType, tkn.Type)
		}
	}
}
