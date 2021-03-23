package src

import (
	"fmt"
	"math"
)

func Evaluator(n Node) (float64, error) {
	switch n.(type) {
	case AtomicNode:
		atomicNode := n.(AtomicNode)
		return atomicNode.Value, nil
	case AddNode:
		addNode := n.(AddNode)
		return InfixEvaluator(addNode.Left, addNode.Right, func(vl float64, vr float64) float64 {
			return vl + vr
		})
	case SubtractNode:
		subtractNode := n.(SubtractNode)
		return InfixEvaluator(subtractNode.Left, subtractNode.Right, func(vl float64, vr float64) float64 {
			return vl - vr
		})
	case MultiplyNode:
		multiplyNode := n.(MultiplyNode)
		return InfixEvaluator(multiplyNode.Left, multiplyNode.Right, func(vl float64, vr float64) float64 {
			return vl * vr
		})
	case DivideNode:
		divideNode := n.(DivideNode)
		return InfixEvaluator(divideNode.Left, divideNode.Right, func(vl float64, vr float64) float64 {
			return vl / vr
		})
	case ExponentialNode:
		exponentialNode := n.(ExponentialNode)
		return InfixEvaluator(exponentialNode.Left, exponentialNode.Right, func(vl float64, vr float64) float64 {
			return math.Pow(vl, vr)
		})
	case ShiftLeftNode:
		shiftLeftNode := n.(ShiftLeftNode)
		return InfixEvaluator(shiftLeftNode.Left, shiftLeftNode.Right, func(vl float64, vr float64) float64 {
			return float64(int(vl) << int(vr))
		})
	case ShiftRightNode:
		shiftRightNode := n.(ShiftRightNode)
		return InfixEvaluator(shiftRightNode.Left, shiftRightNode.Right, func(vl float64, vr float64) float64 {
			return float64(int(vl) >> int(vr))
		})
	default:
		err := fmt.Errorf("failed to evaluate node %s", n)
		return 0, err
	}
}

func InfixEvaluator(leftN Node, rightN Node, op func(vl float64, vr float64) float64) (float64, error) {
	leftR, errL := Evaluator(leftN)
	rightR, errR := Evaluator(rightN)
	if errL != nil || errR != nil {
		err := fmt.Errorf("failed to evaluate infix nodes %s %s", errL, errR)
		return 0, err
	} else {
		return op(leftR, rightR), nil
	}
}
