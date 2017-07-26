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

type FloatSum struct {
	GrammarElement
	BinaryOperationFloat
}

func (f *FloatSum) CalculateFloat() (result float64, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat()
	rightVal, errR := f.GetRightArg().CalculateFloat()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal + rightVal
	}

	return
}

func NewFloatSum(name string) *FloatSum {
	return &FloatSum{
		&grammarElementImpl{name, sheet_logic_types.FloatSum},
		DefaultBinaryOperationFloatImpl()}
}

type StringSum struct {
	GrammarElement
	BinaryOperationString
}

func (s *StringSum) CalculateString() (result string, err error) {
	leftVal, errL := s.GetLeftArg().CalculateString()
	rightVal, errR := s.GetRightArg().CalculateString()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal + rightVal
	}

	return
}

func NewStringSum(name string) *StringSum {
	return &StringSum{
		&grammarElementImpl{name, sheet_logic_types.StringSum},
		DefaultBinaryOperationStringImpl()}
}
