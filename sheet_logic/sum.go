package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type IntSum struct {
	GrammarElement
	BinaryOperationInt
}

func (i *IntSum) CalculateInt(g GrammarContext) (result int64, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt(g)
	rightVal, errR := i.GetRightArg().CalculateInt(g)
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

func (f *FloatSum) CalculateFloat(g GrammarContext) (result float64, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat(g)
	rightVal, errR := f.GetRightArg().CalculateFloat(g)
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

func (s *StringSum) CalculateString(g GrammarContext) (result string, err error) {
	leftVal, errL := s.GetLeftArg().CalculateString(g)
	rightVal, errR := s.GetRightArg().CalculateString(g)
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
