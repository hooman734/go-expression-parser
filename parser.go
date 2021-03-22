package go_expression_parser

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

import (
	"fmt"
	. "github.com/amir734jj/go-lexer"
	underscore "github.com/ahl5esoft/golang-underscore"
)

type Node interface {
	ToString() string
}

type AtomicNode struct {
	Node
	Value float64
}

func (receiver AtomicNode) ToString() string {
	return fmt.Sprintf("%f", receiver.Value)
}

type AddNode struct {
	Node
	Left  Node
	Right Node
}

func (receiver AddNode) ToString() string {
	return fmt.Sprintf(" (%s + %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type SubtractNode struct {
	Node
	Left  Node
	Right Node
}

func (receiver SubtractNode) ToString() string {
	return fmt.Sprintf(" (%s - %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type MultiplyNode struct {
	Node
	Left  Node
	Right Node
}

func (receiver MultiplyNode) ToString() string {
	return fmt.Sprintf(" (%s * %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type DivideNode struct {
	Node
	Left  Node
	Right Node
}

func (receiver DivideNode) ToString() string {
	return fmt.Sprintf(" (%s / %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type ExponentialNode struct {
	Node
	Left  Node
	Right Node
}

func (receiver ExponentialNode) ToString() string {
	return fmt.Sprintf(" (%s ^ %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type ShiftLeftNode struct {
	Node
	Left  Node
	Right Node
}

func (receiver ShiftLeftNode) ToString() string {
	return fmt.Sprintf(" (%s << %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type ShiftRightNode struct {
	Node
	Left  Node
	Right Node
}

func (receiver ShiftRightNode) ToString() string {
	return fmt.Sprintf(" (%s >> %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

func Evaluator(n Node) (float64, error) {
	switch n.(type) {
	case AtomicNode:
		atomicNode := n.(AtomicNode)
		return atomicNode.Value, nil
	case AddNode:
		addNode := n.(AddNode)
		return infixEvaluator(addNode.Left, addNode.Right, func(vl float64, vr float64) float64 {
			return vl + vr
		})
	case SubtractNode:
		subtractNode := n.(SubtractNode)
		return infixEvaluator(subtractNode.Left, subtractNode.Right, func(vl float64, vr float64) float64 {
			return vl - vr
		})
	case MultiplyNode:
		multiplyNode := n.(MultiplyNode)
		return infixEvaluator(multiplyNode.Left, multiplyNode.Right, func(vl float64, vr float64) float64 {
			return vl * vr
		})
	case DivideNode:
		divideNode := n.(DivideNode)
		return infixEvaluator(divideNode.Left, divideNode.Right, func(vl float64, vr float64) float64 {
			return vl / vr
		})
	case ExponentialNode:
		exponentialNode := n.(ExponentialNode)
		return infixEvaluator(exponentialNode.Left, exponentialNode.Right, func(vl float64, vr float64) float64 {
			return math.Pow(vl, vr)
		})
	case ShiftLeftNode:
		shiftLeftNode := n.(ShiftLeftNode)
		return infixEvaluator(shiftLeftNode.Left, shiftLeftNode.Right, func(vl float64, vr float64) float64 {
			return float64(int(vl) << int(vr))
		})
	case ShiftRightNode:
		shiftRightNode := n.(ShiftRightNode)
		return infixEvaluator(shiftRightNode.Left, shiftRightNode.Right, func(vl float64, vr float64) float64 {
			return float64(int(vl) >> int(vr))
		})
	default:
		err := fmt.Errorf("failed to evaluate node %s", n)
		return 0, err
	}
}

func infixEvaluator(leftN Node, rightN Node, op func(vl float64, vr float64) float64) (float64, error) {
	leftR, errL := Evaluator(leftN)
	rightR, errR := Evaluator(rightN)
	if errL != nil || errR != nil {
		err := fmt.Errorf("failed to evaluate infix nodes %s %s", errL, errR)
		return 0, err
	} else {
		return op(leftR, rightR), nil
	}
}

func generateCombinations(str []Token, sep string) [][2]string {
	result := strings.Split(str, sep)
	combinations := make([][2]string, len(result)-1)

	for i := 0; i < len(result)-1; i++ {
		combinations[i] = [2]string{strings.Join(result[:i+1], sep), strings.Join(result[i+1:], sep)}
	}

	return combinations
}

func infixParser(tokens [2]string, operator string) (Node, Node, error) {
	leftN, errL := Parser(tokens[0])
	rightN, errR := Parser(tokens[1])

	if errL != nil || errR != nil {
		err := fmt.Errorf("unable to parse '%s'", strings.Join(tokens[:], operator))
		return leftN, rightN, err
	}

	return leftN, rightN, nil
}

func Parser(tokens []Token) (Node, error) {
	const (
		ShiftLeft  = "<<"
		ShiftRight = ">>"
		PLUS       = "+"
		MINUS      = "-"
		MULTIPLY   = "*"
		DIVIDE     = "/"
		EXPONENT   = "^"
	)
	operators := [7]string{ShiftLeft, ShiftRight, PLUS, MINUS, MULTIPLY, DIVIDE, EXPONENT}

	if len(tokens) == 0 {
		tokenValues := make([]string, 0)
		underscore.Chain(tokens).Select(func(token Token, _ int) string {
			return token.Name
		}).Value(&tokenValues)

		err := fmt.Errorf("unable to parse '%s'", strings.Join(tokenValues, ", "))
		return nil, err
	}

	for i := 0; i < len(operators); i++ {
		operator := operators[i]
		if underscore.Chain(tokens).Any(func(token Token) bool { return token.Name == operator }) {
			for _, result := range generateCombinations(tokens, operator) {
				leftN, rightN, err := infixParser(result, operator)

				if err != nil {
					continue
				}

				switch operator {
				case ShiftLeft:
					return ShiftLeftNode{Left: leftN, Right: rightN}, nil
				case ShiftRight:
					return ShiftRightNode{Left: leftN, Right: rightN}, nil
				case PLUS:
					return AddNode{Left: leftN, Right: rightN}, nil
				case MINUS:
					return SubtractNode{Left: leftN, Right: rightN}, nil
				case MULTIPLY:
					return MultiplyNode{Left: leftN, Right: rightN}, nil
				case DIVIDE:
					return DivideNode{Left: leftN, Right: rightN}, nil
				case EXPONENT:
					return ExponentialNode{Left: leftN, Right: rightN}, nil
				}

			}
		}
	}

	value, err := strconv.ParseFloat(tokens, 64)
	return AtomicNode{Value: value}, err
}
