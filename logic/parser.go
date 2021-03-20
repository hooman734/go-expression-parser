package logic

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func generateCombinations(str string, sep string) [][2]string {
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

func Parser(str string) (Node, error) {
	const (
		PLUS     = "+"
		MINUS    = "-"
		MULTIPLY = "*"
		DIVIDE   = "/"
		EXPONENT = "^"
	)
	operators := [5]string{PLUS, MINUS, MULTIPLY, DIVIDE, EXPONENT}

	str = strings.TrimSpace(str)

	if str == "" {
		err := fmt.Errorf("unable to parse '%s'", str)
		return nil, err
	}

	for i := 0; i < len(operators); i++ {
		operator := operators[i]
		if strings.Contains(str, operator) {
			for _, result := range generateCombinations(str, operator) {
				leftN, rightN, err := infixParser(result, operator)

				if err != nil {
					continue
				}

				if operator == PLUS {
					return AddNode{Left: leftN, Right: rightN}, nil
				} else if operator == MINUS {
					return SubtractNode{Left: leftN, Right: rightN}, nil
				} else if operator == MULTIPLY {
					return MultiplyNode{Left: leftN, Right: rightN}, nil
				} else if operator == DIVIDE {
					return DivideNode{Left: leftN, Right: rightN}, nil
				} else if operator == EXPONENT {
					return ExponentialNode{Left: leftN, Right: rightN}, nil
				}
			}
		}
	}

	value, err := strconv.ParseFloat(str, 64)
	return AtomicNode{Value: value}, err
}
