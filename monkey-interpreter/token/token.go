package token

// token literal (how is it visible)
// type

type TokenType int

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "illegal"
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
)
