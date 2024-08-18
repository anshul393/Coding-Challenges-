package lexer

import "github.com/anshul393/coding-challenges/monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	return &Lexer{input: input, readPosition: 1}
}

func (l *Lexer) NextToken() token.Token {
	if l.position >= len(l.input) {
		return token.Token{Type: token.EOF, Literal: ""}
	}

	return token.Token{}
}
