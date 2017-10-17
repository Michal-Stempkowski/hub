package sheet_logic

import "hub/sheet_logic/sheet_logic_types"

type IntComparator struct {
	GrammarElement
	BinaryOperationInt
	comparatorFunc func(int64, int64) bool
}

func (i *IntComparator) CalculateBool(g GrammarContext) (result bool, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt(g)
	rightVal, errR := i.GetRightArg().CalculateInt(g)
	if err = getFirstError(errL, errR); err == nil {
		result = i.comparatorFunc(leftVal, rightVal)
	}

	return
}

func NewIntComparator(
	name string, t sheet_logic_types.T, cmp func(int64, int64) bool) *IntComparator {
	return &IntComparator{
		&grammarElementImpl{name, t},
		DefaultBinaryOperationIntImpl(),
		cmp}
}

type FloatComparator struct {
	GrammarElement
	BinaryOperationFloat
	comparatorFunc func(float64, float64) bool
}

func (f *FloatComparator) CalculateBool(g GrammarContext) (result bool, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat(g)
	rightVal, errR := f.GetRightArg().CalculateFloat(g)
	if err = getFirstError(errL, errR); err == nil {
		result = f.comparatorFunc(leftVal, rightVal)
	}

	return
}

func NewFloatComparator(
	name string, t sheet_logic_types.T, cmp func(float64, float64) bool) *FloatComparator {
	return &FloatComparator{
		&grammarElementImpl{name, t},
		DefaultBinaryOperationFloatImpl(),
		cmp}
}

type BoolComparator struct {
	GrammarElement
	BinaryOperationBool
	comparatorFunc func(bool, bool) bool
}

func (b *BoolComparator) CalculateBool(g GrammarContext) (result bool, err error) {
	leftVal, errL := b.GetLeftArg().CalculateBool(g)
	rightVal, errR := b.GetRightArg().CalculateBool(g)
	if err = getFirstError(errL, errR); err == nil {
		result = b.comparatorFunc(leftVal, rightVal)
	}

	return
}

func NewBoolComparator(
	name string, t sheet_logic_types.T, cmp func(bool, bool) bool) *BoolComparator {
	return &BoolComparator{
		&grammarElementImpl{name, t},
		DefaultBinaryOperationBoolImpl(),
		cmp}
}

type StringComparator struct {
	GrammarElement
	BinaryOperationString
	comparatorFunc func(string, string) bool
}

func (s *StringComparator) CalculateBool(g GrammarContext) (result bool, err error) {
	leftVal, errL := s.GetLeftArg().CalculateString(g)
	rightVal, errR := s.GetRightArg().CalculateString(g)
	if err = getFirstError(errL, errR); err == nil {
		result = s.comparatorFunc(leftVal, rightVal)
	}

	return
}

func NewStringComparator(
	name string, t sheet_logic_types.T, cmp func(string, string) bool) *StringComparator {
	return &StringComparator{
		&grammarElementImpl{name, t},
		DefaultBinaryOperationStringImpl(),
		cmp}
}
