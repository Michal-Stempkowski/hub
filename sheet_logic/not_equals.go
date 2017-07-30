package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
)

type IntNotEquals IntComparator

func NewIntNotEquals(name string) *IntNotEquals {
	tmp := NewIntComparator(
		name,
		sheet_logic_types.IntNotEquals,
		func(a int64, b int64) bool { return a != b })
	return (*IntNotEquals)(tmp)
}

type FloatNotEquals FloatComparator

func NewFloatNotEquals(name string) *FloatNotEquals {
	tmp := NewFloatComparator(
		name,
		sheet_logic_types.FloatNotEquals,
		func(a float64, b float64) bool { return !framework.FloatEq(a, b) })
	return (*FloatNotEquals)(tmp)
}

type BoolNotEquals BoolComparator

func NewBoolNotEquals(name string) *BoolNotEquals {
	tmp := NewBoolComparator(
		name,
		sheet_logic_types.BoolNotEquals,
		func(a bool, b bool) bool { return a != b })
	return (*BoolNotEquals)(tmp)
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
