package lex

type Lex struct {
	code   string
	index  int
	buffer []byte
}

func NewLex(code string) *Lex {
	return &Lex{code, 0, []byte{}}
}

func (l *Lex) NextToken() Token {
	return Token{}
}

func (l *Lex) ParseAll() []Token {
	ans := make([]Token, 0)
	var token Token
	for token != EOF {
		token = l.NextToken()
		ans = append(ans, token)
	}
	return ans
}
