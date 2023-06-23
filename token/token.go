package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	// priority constants
	_ uint8 = iota
	LOWEST
	EQUALS      // ==
	LessGreater // > or <
	SUM         //+
	PRODUCT     //*
	PREFIX      //-Xor!X
	CALL        // myFunction(X)
	INDEX       // array[index]
	// base
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"
	STRING = "STRING"
	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	// equal and not equal
	EQ    = "=="
	NotEq = "!="
	// array
	LBRACKET = "["
	RBRACKET = "]"
	// hash map
	COLON = ":"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
