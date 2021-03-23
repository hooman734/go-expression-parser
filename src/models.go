package src

import "fmt"

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
