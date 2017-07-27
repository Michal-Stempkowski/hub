package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type IntMultiplication struct {
	GrammarElement
	BinaryOperationInt
}

func (i *IntMultiplication) CalculateInt() (result int64, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt()
	rightVal, errR := i.GetRightArg().CalculateInt()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal * rightVal
	}

	return
}

func NewIntMultiplication(name string) *IntMultiplication {
	return &IntMultiplication{
		&grammarElementImpl{name, sheet_logic_types.IntMultiplication},
		DefaultBinaryOperationIntImpl()}
}

type FloatMultiplication struct {
	GrammarElement
	BinaryOperationFloat
}

func (f *FloatMultiplication) CalculateFloat() (result float64, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat()
	rightVal, errR := f.GetRightArg().CalculateFloat()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal * rightVal
	}

	return
}

func NewFloatMultiplication(name string) *FloatMultiplication {
	return &FloatMultiplication{
		&grammarElementImpl{name, sheet_logic_types.FloatMultiplication},
		DefaultBinaryOperationFloatImpl()}
}
