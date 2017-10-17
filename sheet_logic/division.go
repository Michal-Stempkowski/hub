package sheet_logic

import (
	"fmt"
	"hub/sheet_logic/sheet_logic_types"
)

type IntDivision struct {
	GrammarElement
	BinaryOperationInt
}

func (i *IntDivision) CalculateInt(g GrammarContext) (result int64, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt(g)
	rightVal, errR := i.GetRightArg().CalculateInt(g)
	var errDivisionByZero error
	if rightVal == 0 {
		errDivisionByZero = fmt.Errorf("IntDivision: Division by zero!")
	}
	if err = getFirstError(errL, errR, errDivisionByZero); err == nil {
		result = leftVal / rightVal
	}

	return
}

func NewIntDivision(name string) *IntDivision {
	return &IntDivision{
		&grammarElementImpl{name, sheet_logic_types.IntDivision},
		DefaultBinaryOperationIntImpl()}
}

type FloatDivision struct {
	GrammarElement
	BinaryOperationFloat
}

func (f *FloatDivision) CalculateFloat(g GrammarContext) (result float64, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat(g)
	rightVal, errR := f.GetRightArg().CalculateFloat(g)
	var errDivisionByZero error
	if rightVal == 0 {
		errDivisionByZero = fmt.Errorf("FloatDivision: Division by zero!")
	}
	if err = getFirstError(errL, errR, errDivisionByZero); err == nil {
		result = leftVal / rightVal
	}

	return
}

func NewFloatDivision(name string) *FloatDivision {
	return &FloatDivision{
		&grammarElementImpl{name, sheet_logic_types.FloatDivision},
		DefaultBinaryOperationFloatImpl()}
}
