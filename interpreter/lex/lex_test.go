package lex

import (
	"reflect"
	"testing"
)

func TestLex_ParseAll(t *testing.T) {
	tests := []struct {
		name string
		code string
		want []Token
	}{
		{name: "code1", code: "let five = 5;", want: []Token{
			LET,
			NewToken(IDENT, "five"),
			ASSIGN,
			NewToken(INT, "5"),
			SEMICOLON,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLex(tt.code)
			if got := l.ParseAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
