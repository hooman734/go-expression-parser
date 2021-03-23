package src

import (
	. "github.com/amir734jj/go-lexer"
)

func Lexer() func(text string) ([]Token, error) {
	lexer := NewLexer().
		Add(Token{Name: "NUMBER", Pattern: "^[0-9]+$"}).
		Add(Token{Name: "PLUS", Pattern: "^\\+$"}).
		Add(Token{Name: "MINUS", Pattern: "^\\-$"}).
		Add(Token{Name: "MULTIPLY", Pattern: "^\\*$"}).
		Add(Token{Name: "DIVIDE", Pattern: "^\\/$"}).
		Add(Token{Name: "L_PRN", Pattern: "^\\($"}).
		Add(Token{Name: "R_PRN", Pattern: "^\\)$"}).
		Add(Token{Name: "SPACE", Pattern: "^\\s+$", Ignore: true}).
		Build()

	return lexer
}
