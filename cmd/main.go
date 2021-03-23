package main

import (
	"fmt"
	"go-expression-parser/src"
)

func main() {
	str := "2 + 3"

	lexerTokens, lexerError := src.Lexer()(str)
	if lexerError != nil {
		panic(lexerError)
	}

	n, parseError := src.Parse(lexerTokens)
	if parseError != nil {
		panic(parseError)
	}

	v, evalError := src.Evaluator(n)
	if evalError != nil {
		panic(evalError)
	}
	fmt.Printf("Result: %s = %f", n.ToString(), v)
}
