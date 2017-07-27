package sheet_logic

import (
	"fmt"
	"hub/sheet_logic/sheet_logic_types"
)

type IntModulo struct {
	GrammarElement
	BinaryOperationInt
}

func (i *IntModulo) CalculateInt() (result int64, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt()
	rightVal, errR := i.GetRightArg().CalculateInt()
	var errDivisionByZero error
	if rightVal == 0 {
		errDivisionByZero = fmt.Errorf("IntDivision: Division by zero!")
	}
	if err = getFirstError(errL, errR, errDivisionByZero); err == nil {
		result = leftVal % rightVal
	}

	return
}

func NewIntModulo(name string) *IntModulo {
	return &IntModulo{
		&grammarElementImpl{name, sheet_logic_types.IntModulo},
		DefaultBinaryOperationIntImpl()}
}
