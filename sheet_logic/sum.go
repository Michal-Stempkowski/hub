package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type IntSum struct {
	GrammarElement
	BinaryOperationInt
}

func (i *IntSum) CalculateInt() (result int64, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt()
	rightVal, errR := i.GetRightArg().CalculateInt()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal + rightVal
	}

	return
}

func NewIntSum(name string) *IntSum {
	return &IntSum{
		&grammarElementImpl{name, sheet_logic_types.IntSum},
		DefaultBinaryOperationIntImpl()}
}
