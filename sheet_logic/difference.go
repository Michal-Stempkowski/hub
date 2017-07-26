package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type IntDifference struct {
	GrammarElement
	BinaryOperationInt
}

func (i *IntDifference) CalculateInt() (result int64, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt()
	rightVal, errR := i.GetRightArg().CalculateInt()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal - rightVal
	}

	return
}

func NewIntDifference(name string) *IntDifference {
	return &IntDifference{
		&grammarElementImpl{name, sheet_logic_types.IntDifference},
		DefaultBinaryOperationIntImpl()}
}

type FloatDifference struct {
	GrammarElement
	BinaryOperationFloat
}

func (f *FloatDifference) CalculateFloat() (result float64, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat()
	rightVal, errR := f.GetRightArg().CalculateFloat()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal - rightVal
	}

	return
}

func NewFloatDifference(name string) *FloatDifference {
	return &FloatDifference{
		&grammarElementImpl{name, sheet_logic_types.FloatDifference},
		DefaultBinaryOperationFloatImpl()}
}
