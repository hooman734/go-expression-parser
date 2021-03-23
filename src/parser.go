package src

import (
	"fmt"
	. "github.com/amir734jj/go-lexer"
	"strconv"
)

const (
	ShiftLeft  = "SHIFT_LEFT"
	ShiftRight = "SHIFT_RIGHT"
	PLUS       = "PLUS"
	MINUS      = "MINUS"
	MULTIPLY   = "MULTIPLY"
	DIVIDE     = "DIVIDE"
	EXPONENT   = "EXPONENT"
	NUMBER     = "NUMBER"
	LPrn       = "L_PRN"
	RPrn       = "R_PRN"
)

//   E -> T
//      | T + E
//
//   T -> a
//      | a * T
//      | (E)
type RecursiveDescentParserState struct {
	Cursor     int
	PrevCursor int
	Tokens     []Token
}

func (ctx *RecursiveDescentParserState) EOF() bool {
	return ctx.Cursor == len(ctx.Tokens)
}

func (ctx *RecursiveDescentParserState) Current() Token {
	return ctx.Tokens[ctx.Cursor-1]
}

func (ctx *RecursiveDescentParserState) BackTrack() {
	ctx.Cursor = ctx.PrevCursor
}

func (ctx *RecursiveDescentParserState) SaveCursor() {
	ctx.PrevCursor = ctx.Cursor
}

func (ctx *RecursiveDescentParserState) GetNextToken() Token {
	nextToken := ctx.Tokens[ctx.Cursor]
	ctx.Cursor++
	return nextToken
}

func (ctx *RecursiveDescentParserState) Term(name string) bool {
	return ctx.GetNextToken().Name == name
}

//   E -> T
//      | T + E
//		| T - E
func ParsePlusMinus(ctx *RecursiveDescentParserState) (Node, error) {
	ctx.SaveCursor()
	r0, err0 := ParsePlusMinus0(ctx)
	if err0 == nil && ctx.EOF() {
		return r0, nil
	}

	ctx.BackTrack()
	r1, err1 := ParsePlusMinus1(ctx)
	if err1 == nil && ctx.EOF() {
		return r1, nil
	}

	ctx.BackTrack()
	r2, err2 := ParsePlusMinus2(ctx)
	if err2 == nil && ctx.EOF() {
		return r2, nil
	}

	return nil, fmt.Errorf("ParsePlusMinus failed")
}

func ParsePlusMinus0(ctx *RecursiveDescentParserState) (Node, error) {
	return ParseMultiplyDivide(ctx)
}

func ParsePlusMinus1(ctx *RecursiveDescentParserState) (Node, error) {
	resultL, errL := ParseMultiplyDivide(ctx)

	if errL == nil && ctx.Term(PLUS) {
		resultR, errR := ParsePlusMinus(ctx)

		if errR == nil {
			return AddNode{Left: resultL, Right: resultR}, nil
		}
	}

	return nil, fmt.Errorf("ParsePlusMinus1 failed")
}

func ParsePlusMinus2(ctx *RecursiveDescentParserState) (Node, error) {
	resultL, errL := ParseMultiplyDivide(ctx)

	if errL == nil && ctx.Term(MINUS) {
		resultR, errR := ParsePlusMinus(ctx)

		if errR == nil {
			return SubtractNode{Left: resultL, Right: resultR}, nil
		}
	}

	return nil, fmt.Errorf("ParsePlusMinus1 failed")
}

//   T -> a
//      | a * T
//      | a / T
//      | (E)
func ParseMultiplyDivide(ctx *RecursiveDescentParserState) (Node, error) {
	ctx.SaveCursor()
	r0, err0 := ParseMultiplyDivide0(ctx)
	if err0 == nil {
		return r0, nil
	}

	ctx.BackTrack()
	r1, err1 := ParseMultiplyDivide1(ctx)
	if err1 == nil {
		return r1, nil
	}

	ctx.BackTrack()
	r2, err2 := ParseMultiplyDivide2(ctx)
	if err2 == nil {
		return r2, nil
	}

	ctx.BackTrack()
	r3, err3 := ParseMultiplyDivide3(ctx)
	if err3 == nil {
		return r3, nil
	}

	return nil, fmt.Errorf("ParseMultiplyDivide failed")
}

func ParseMultiplyDivide0(ctx *RecursiveDescentParserState) (Node, error) {
	if ctx.Term(NUMBER) {
		value, err := strconv.ParseFloat(ctx.Current().Value, 64)
		return AtomicNode{Value: value}, err
	}
	return nil, fmt.Errorf("ParseMultiplyDivide0 failed")
}

func ParseMultiplyDivide1(ctx *RecursiveDescentParserState) (Node, error) {
	if ctx.Term(NUMBER) {
		value, errL := strconv.ParseFloat(ctx.Current().Value, 64)
		resultL := AtomicNode{Value: value}
		if errL == nil && ctx.Term(MULTIPLY) {
			resultR, errR := ParseMultiplyDivide(ctx)
			if errR == nil {
				return MultiplyNode{Left: resultL, Right: resultR}, nil
			}
		}
	}

	return nil, fmt.Errorf("ParseMultiplyDivide1 failed")
}

func ParseMultiplyDivide2(ctx *RecursiveDescentParserState) (Node, error) {
	if ctx.Term(NUMBER) {
		value, errL := strconv.ParseFloat(ctx.Current().Value, 64)
		resultL := AtomicNode{Value: value}
		if errL == nil && ctx.Term(DIVIDE) {
			resultR, errR := ParseMultiplyDivide(ctx)
			if errR == nil {
				return DivideNode{Left: resultL, Right: resultR}, nil
			}
		}
	}

	return nil, fmt.Errorf("ParseMultiplyDivide2 failed")
}

func ParseMultiplyDivide3(ctx *RecursiveDescentParserState) (Node, error) {
	if ctx.Term(LPrn) {
		result, err := ParsePlusMinus(ctx)
		if err == nil && ctx.Term(RPrn) {
			return result, nil
		}
	}

	return nil, fmt.Errorf("ParseMultiplyDivide3 failed")
}

func Parse(tokens []Token) (Node, error) {
	ctx := RecursiveDescentParserState{Cursor: 0, PrevCursor: 0, Tokens: tokens}

	result, err := ParsePlusMinus(&ctx)

	if err == nil && ctx.Cursor == len(ctx.Tokens) {
		return result, nil
	}

	return nil, fmt.Errorf("parse failed")
}
