package main

import (
	. "./logic"
	"fmt"
)

func main() {
	n, parseError := Parser("2 * +3 - 6 / 4 + 5 + 2 ^ 3 + 2 << 3")
	if parseError != nil {
		panic(parseError)
	}

	v, evalError := Evaluator(n)
	if evalError != nil {
		panic(evalError)
	}
	fmt.Printf("Result: %s = %f", n.ToString(), v)
}
