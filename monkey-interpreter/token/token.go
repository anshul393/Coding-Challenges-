package token

// token literal (how is it visible)
// type

type TokenType int

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENTIFIER"
	case INT:
		return "INT"
	case ASSIGN:
		return "="
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case BANG:
		return "!"
	case ASTERISK:
		return "*"
	case SLASH:
		return "/"
	case LT:
		return "<"
	case GT:
		return ">"

	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	default:
		return ""
	}
}

type Token struct {
	// + return PLUS tokentype
	// "" return EOF tokentype
	Literal string
	Type    TokenType
}

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers and literals
	IDENT // x,y,foobar
	INT   // 1,2,3,111212

	// OPERATORS
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH
	LT
	GT

	// DELIMETERS
	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// KEYWORDS
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
)

func NewToken(literal string, tokenType TokenType) Token {
	return Token{Literal: literal, Type: tokenType}
}

func GetToken(literal string) Token {

	// words --> KeyWords, Identifier
	keywordMap := map[string]TokenType{
		"fn":     FUNCTION,
		"let":    LET,
		"true":   TRUE,
		"false":  FALSE,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
	}

	if tokentype, ok := keywordMap[literal]; ok {
		return NewToken(literal, tokentype)
	}

	return NewToken(literal, IDENT)
}
