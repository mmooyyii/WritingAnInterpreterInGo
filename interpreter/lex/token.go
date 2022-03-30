package lex

type TokenType string

const (
	INT   = "INT"   // integer
	IDENT = "IDENT" // 变量
)

var (
	EOF       = NewUniqToken("eof")
	LET       = NewUniqToken("let") // 声明变量
	SEMICOLON = NewUniqToken(";")
	ASSIGN    = NewUniqToken("=") // 赋值
)

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(tp TokenType, value string) Token { return Token{tp, value} }
func NewUniqToken(tp TokenType) Token           { return Token{tp, ""} }
