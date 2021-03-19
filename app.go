package main

import (
	"fmt"
	. "go-expression-parser/logic"
)

func main() {
	n, _ := Parser(" 2 * +3  - 6 / 4 + 5")
	v, _ := Evaluator(n)
	fmt.Printf("Result: %s = %f", n.ToString(), v)
}