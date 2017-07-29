package sheet_logic

import (
	"hub/framework"
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
		result = framework.FloatEq(leftVal, rightVal)
	}

	return
}

func NewFloatEquals(name string) *FloatEquals {
	return &FloatEquals{
		&grammarElementImpl{name, sheet_logic_types.FloatEquals},
		DefaultBinaryOperationFloatImpl()}
}

type BoolEquals struct {
	GrammarElement
	BinaryOperationBool
}

func (b *BoolEquals) CalculateBool() (result bool, err error) {
	leftVal, errL := b.GetLeftArg().CalculateBool()
	rightVal, errR := b.GetRightArg().CalculateBool()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal == rightVal
	}

	return
}

func NewBoolEquals(name string) *BoolEquals {
	return &BoolEquals{
		&grammarElementImpl{name, sheet_logic_types.BoolEquals},
		DefaultBinaryOperationBoolImpl()}
}

type StringEquals struct {
	GrammarElement
	BinaryOperationString
}

func (s *StringEquals) CalculateBool() (result bool, err error) {
	leftVal, errL := s.GetLeftArg().CalculateString()
	rightVal, errR := s.GetRightArg().CalculateString()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal == rightVal
	}

	return
}

func NewStringEquals(name string) *StringEquals {
	return &StringEquals{
		&grammarElementImpl{name, sheet_logic_types.StringEquals},
		DefaultBinaryOperationStringImpl()}
}
