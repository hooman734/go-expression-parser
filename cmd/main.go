package main

import (
	"fmt"
	. "github.com/amir734jj/go-lexer"
	. "go-expression-parser"
)

func main() {
	lexer := NewLexer().
		Add(Token{Name: "NUMBER", Pattern: "^[0-9]+$"}).
		Add(Token{Name: "PLUS", Pattern: "^\\+$"}).
		Add(Token{Name: "MINUS", Pattern: "^\\-$"}).
		Add(Token{Name: "MULTIPLY", Pattern: "^\\*$"}).
		Add(Token{Name: "DIVIDE", Pattern: "^\\/$"}).
		Add(Token{Name: "EXPONENT", Pattern: "^\\^$"}).
		Add(Token{Name: "SHIFT_LEFT", Pattern: "^<<$"}).
		Add(Token{Name: "SHIFT_RIGHT", Pattern: "^>>$"}).
		Add(Token{Name: "SPACE", Pattern: "^\\s+$", Ignore: true}).
		Build()

	str := "2 * +3 - 6 / 4 + 5 + 2 ^ 3 + 2 << 3"

	lexerTokens, lexerError := lexer(str)
	if lexerError != nil {
		panic(lexerError)
	}

	n, parseError := Parser(lexerTokens)
	if parseError != nil {
		panic(parseError)
	}

	v, evalError := Evaluator(n)
	if evalError != nil {
		panic(evalError)
	}
	fmt.Printf("Result: %s = %f", n.ToString(), v)
}
