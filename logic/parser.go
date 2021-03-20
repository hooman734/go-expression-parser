package logic

import (
	"fmt"
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
	Left  Node
	Right Node
}

func (receiver AddNode) ToString() string {
	return fmt.Sprintf(" (%s + %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type SubtractNode struct {
	Left  Node
	Right Node
}

func (receiver SubtractNode) ToString() string {
	return fmt.Sprintf(" (%s - %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type MultiplyNode struct {
	Left  Node
	Right Node
}

func (receiver MultiplyNode) ToString() string {
	return fmt.Sprintf(" (%s * %s) ", receiver.Left.ToString(), receiver.Right.ToString())
}

type DivideNode struct {
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

func combinations(str string, sep string) [][2]string {
	result := strings.Split(str, sep)
	combinations := make([][2]string, len(result)-1)

	for i := 0; i < len(result)-1; i++ {
		combinations[i] = [2]string{strings.Join(result[:i+1], sep), strings.Join(result[i+1:], sep)}
	}

	return combinations
}

func infix(tokens [2]string, operator string) (Node, Node, error) {
	leftN, errL := Parser(tokens[0])
	rightN, errR := Parser(tokens[1])

	if errL != nil || errR != nil {
		err := fmt.Errorf("unable to parse '%s'", strings.Join(tokens[:], operator))
		return leftN, rightN, err
	}

	return leftN, rightN, nil
}

func Parser(str string) (Node, error) {
	operators := [5]int{'+', '-', '*', '/'}
	str = strings.TrimSpace(str)

	if str == "" {
		err := fmt.Errorf("unable to parse '%s'", str)
		return nil, err
	}

	for i := 0; i < len(operators); i++ {
		operator := string(rune(operators[i]))
		if strings.Contains(str, operator) {
			for _, result := range combinations(str, operator) {
				leftN, rightN, err := infix(result, operator)

				if err != nil {
					continue
				}

				if operator == "+" {
					return AddNode{Left: leftN, Right: rightN}, nil
				} else if operator == "-" {
					return SubtractNode{Left: leftN, Right: rightN}, nil
				} else if operator == "*" {
					return MultiplyNode{Left: leftN, Right: rightN}, nil
				} else if operator == "/" {
					return DivideNode{Left: leftN, Right: rightN}, nil
				}
			}
		}
	}

	value, err := strconv.ParseFloat(str, 64)
	return AtomicNode{Value: value}, err
}
