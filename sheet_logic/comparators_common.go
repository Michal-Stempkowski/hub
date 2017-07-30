package sheet_logic

import "hub/sheet_logic/sheet_logic_types"

type IntComparator struct {
	GrammarElement
	BinaryOperationInt
	comparatorFunc func(int64, int64) bool
}

func (i *IntComparator) CalculateBool() (result bool, err error) {
	leftVal, errL := i.GetLeftArg().CalculateInt()
	rightVal, errR := i.GetRightArg().CalculateInt()
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

func (f *FloatComparator) CalculateBool() (result bool, err error) {
	leftVal, errL := f.GetLeftArg().CalculateFloat()
	rightVal, errR := f.GetRightArg().CalculateFloat()
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

func (b *BoolComparator) CalculateBool() (result bool, err error) {
	leftVal, errL := b.GetLeftArg().CalculateBool()
	rightVal, errR := b.GetRightArg().CalculateBool()
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
