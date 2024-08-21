package lexer

import "github.com/anshul393/coding-challenges/monkey-interpreter/token"

type Lexer struct {
	input string
	curr  int
	prev  int
	ch    byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token // Literal tokenType

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		tok = token.NewToken("=", token.ASSIGN)
	case '+':
		tok = token.NewToken("+", token.PLUS)
	case '-':
		tok = token.NewToken("-", token.MINUS)
	case '!':
		tok = token.NewToken("!", token.BANG)
	case '*':
		tok = token.NewToken("-", token.ASTERISK)
	case '/':
		tok = token.NewToken("/", token.SLASH)
	case '<':
		tok = token.NewToken("<", token.LT)
	case '>':
		tok = token.NewToken(">", token.GT)
	case ',':
		tok = token.NewToken(",", token.COMMA)
	case ';':
		tok = token.NewToken(";", token.SEMICOLON)
	case '(':
		tok = token.NewToken("(", token.LPAREN)
	case ')':
		tok = token.NewToken(")", token.RPAREN)
	case '{':
		tok = token.NewToken("{", token.LBRACE)
	case '}':
		tok = token.NewToken("}", token.RBRACE)
	case 0:
		tok = token.NewToken("", token.EOF)
	default:
		if isLetter(l.ch) {
			// Read the identifier
			literal := l.readIdentifier()
			tok = token.GetToken(literal)

		} else if isDigit(l.ch) {
			// Read the Number
			tok = token.NewToken(l.readNumber(), token.INT)
		} else {
			tok = token.NewToken("ILLEGAL", token.ILLEGAL)
		}
		return tok
	}

	l.readChar()

	return tok
}

func (l *Lexer) readChar() {
	if l.curr >= len(l.input) {
		l.ch = 0
		return
	} else {
		l.ch = l.input[l.curr]
	}

	l.prev = l.curr
	l.curr += 1
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	prev := l.prev
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[prev:l.prev]
}

func (l *Lexer) readNumber() string {
	prev := l.prev
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[prev:l.prev]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}
