package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type IntSum struct {
	*grammarElementImpl
	argLeft  IntExpresion
	argRight IntExpresion
}

func (i *IntSum) CalculateInt() (result int64, err error) {
	leftVal, errL := i.argLeft.CalculateInt()
	rightVal, errR := i.argRight.CalculateInt()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal + rightVal
	}

	return
}

func NewIntSum(name string, argLeft, argRight IntExpresion) *IntSum {
	return &IntSum{
		&grammarElementImpl{name, sheet_logic_types.IntSum},
		argLeft,
		argRight}
}
