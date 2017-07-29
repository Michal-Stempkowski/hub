package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type IntEquals struct {
	GrammarElement
	BinaryOperationInt
}

func (i *IntEquals) CalculateBool() (result bool, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt()
	rightVal, errR := i.GetRightArg().CalculateInt()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal == rightVal
	}

	return
}

func NewIntEquals(name string) *IntEquals {
	return &IntEquals{
		&grammarElementImpl{name, sheet_logic_types.IntEquals},
		DefaultBinaryOperationIntImpl()}
}

type FloatEquals struct {
	GrammarElement
	BinaryOperationFloat
}

func (f *FloatEquals) CalculateBool() (result bool, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat()
	rightVal, errR := f.GetRightArg().CalculateFloat()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal == rightVal
	}

	return
}

func NewFloatEquals(name string) *FloatEquals {
	return &FloatEquals{
		&grammarElementImpl{name, sheet_logic_types.FloatEquals},
		DefaultBinaryOperationFloatImpl()}
}
