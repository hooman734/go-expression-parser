package logic

import (
	"fmt"
	"strconv"
	"strings"
)

type Node interface{
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

func Evaluator(n Node) (float64, error) {
	switch n.(type) {
	case AtomicNode:
		atomicNode := n.(AtomicNode)
		return atomicNode.Value, nil
	case AddNode:
		addNode := n.(AddNode)
		leftR, errL := Evaluator(addNode.Left)
		rightR, errR := Evaluator(addNode.Right)
		if errL != nil || errR != nil {
			err := fmt.Errorf("failed to add node %s %s %s", n, errL, errR)
			return 0, err
		} else {
			return leftR + rightR, nil
		}
	case SubtractNode:
		subtractNode := n.(SubtractNode)
		leftR, errL := Evaluator(subtractNode.Left)
		rightR, errR := Evaluator(subtractNode.Right)
		if errL != nil || errR != nil {
			err := fmt.Errorf("failed to subtract node %s %s %s", n, errL, errR)
			return 0, err
		} else {
			return leftR - rightR, nil
		}
	case MultiplyNode:
		multiplyNode := n.(MultiplyNode)
		leftR, errL := Evaluator(multiplyNode.Left)
		rightR, errR := Evaluator(multiplyNode.Right)
		if errL != nil || errR != nil {
			err := fmt.Errorf("failed to multiply node %s %s %s", n, errL, errR)
			return 0, err
		} else {
			return leftR * rightR, nil
		}
	case DivideNode:
		divideNode := n.(DivideNode)
		leftR, errL := Evaluator(divideNode.Left)
		rightR, errR := Evaluator(divideNode.Right)
		if errL != nil || errR != nil {
			err := fmt.Errorf("failed to divide node %s %s %s", n, errL, errR)
			return 0, err
		} else {
			return leftR / rightR, nil
		}
	default:
		err := fmt.Errorf("failed to evaluate node %s", n)
		return 0, err
	}
}

func Parser(str string) (Node, error) {
	operators := [5]int{'+', '-', '*', '/'}

	str = strings.TrimSpace(str)

	if str == "" {
		err := fmt.Errorf("unable to parse %s", str)
		return nil, err
	}

	for _, operator := range operators {
		if strings.Contains(str, string(rune(operator))) {
			result := strings.SplitN(str, string(rune(operator)), 2)
			if operator == '+' {
				leftN, errL := Parser(result[0])
				rightN, errR := Parser(result[1])

				if errL != nil || errR != nil {
					continue
				}

				return AddNode{Left: leftN, Right: rightN}, nil
			} else if operator == '-' {
				leftN, errL := Parser(result[0])
				rightN, errR := Parser(result[1])

				if errL != nil || errR != nil {
					continue
				}

				return SubtractNode{Left: leftN, Right: rightN}, nil
			} else if operator == '*' {
				leftN, errL := Parser(result[0])
				rightN, errR := Parser(result[1])

				if errL != nil || errR != nil {
					continue
				}

				return MultiplyNode{Left: leftN, Right: rightN}, nil
			} else if operator == '/' {
				leftN, errL := Parser(result[0])
				rightN, errR := Parser(result[1])

				if errL != nil || errR != nil {
					continue
				}

				return DivideNode{Left: leftN, Right: rightN}, nil
			}
		}
	}

	value, err := strconv.ParseFloat(str, 64)
	return AtomicNode{Value: value}, err
}
