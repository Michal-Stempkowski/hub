package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
)

type IntNotEquals struct {
	GrammarElement
	BinaryOperationInt
}

func (i *IntNotEquals) CalculateBool() (result bool, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt()
	rightVal, errR := i.GetRightArg().CalculateInt()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal != rightVal
	}

	return
}

func NewIntNotEquals(name string) *IntNotEquals {
	return &IntNotEquals{
		&grammarElementImpl{name, sheet_logic_types.IntNotEquals},
		DefaultBinaryOperationIntImpl()}
}

type FloatNotEquals struct {
	GrammarElement
	BinaryOperationFloat
}

func (f *FloatNotEquals) CalculateBool() (result bool, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat()
	rightVal, errR := f.GetRightArg().CalculateFloat()
	if err = getFirstError(errL, errR); err == nil {
		result = !framework.FloatEq(leftVal, rightVal)
	}

	return
}

func NewFloatNotEquals(name string) *FloatNotEquals {
	return &FloatNotEquals{
		&grammarElementImpl{name, sheet_logic_types.FloatNotEquals},
		DefaultBinaryOperationFloatImpl()}
}

type BoolNotEquals struct {
	GrammarElement
	BinaryOperationBool
}

func (b *BoolNotEquals) CalculateBool() (result bool, err error) {
	leftVal, errL := b.GetLeftArg().CalculateBool()
	rightVal, errR := b.GetRightArg().CalculateBool()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal != rightVal
	}

	return
}

func NewBoolNotEquals(name string) *BoolNotEquals {
	return &BoolNotEquals{
		&grammarElementImpl{name, sheet_logic_types.BoolNotEquals},
		DefaultBinaryOperationBoolImpl()}
}

type StringNotEquals struct {
	GrammarElement
	BinaryOperationString
}

func (s *StringNotEquals) CalculateBool() (result bool, err error) {
	leftVal, errL := s.GetLeftArg().CalculateString()
	rightVal, errR := s.GetRightArg().CalculateString()
	if err = getFirstError(errL, errR); err == nil {
		result = leftVal != rightVal
	}

	return
}

func NewStringNotEquals(name string) *StringNotEquals {
	return &StringNotEquals{
		&grammarElementImpl{name, sheet_logic_types.StringNotEquals},
		DefaultBinaryOperationStringImpl()}
}
