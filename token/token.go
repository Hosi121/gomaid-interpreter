package token

type TokenType string

type Token struct {
	Type     TokenType
	Liteeral string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // End of fileの略

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// 演算子
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNTION = "FUNCTION"
	LET     = "LET"
)
