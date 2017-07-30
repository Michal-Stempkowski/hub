package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
)

type IntEquals IntComparator

func NewIntEquals(name string) *IntEquals {
	tmp := NewIntComparator(
		name,
		sheet_logic_types.IntEquals,
		func(a int64, b int64) bool { return a == b })
	return (*IntEquals)(tmp)
}

type FloatEquals FloatComparator

func NewFloatEquals(name string) *FloatEquals {
	tmp := NewFloatComparator(
		name,
		sheet_logic_types.FloatEquals,
		framework.FloatEq)
	return (*FloatEquals)(tmp)
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
