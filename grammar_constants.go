package sheet_logic

import "hub/sheet_logic_types"

type IntConstant struct {
	*grammarElementImpl
	value int64
}

func (i *IntConstant) CalculateInt() (int64, error) {
	return i.value, nil
}

func NewIntConstant(name string, value int64) *IntConstant {
	return &IntConstant{&grammarElementImpl{name, sheet_logic_types.IntConstant}, value}
}

type StringConstant struct {
	*grammarElementImpl
	value string
}

func (s *StringConstant) CalculateString() (string, error) {
	return s.value, nil
}

func NewStringConstant(name string, value string) *StringConstant {
	return &StringConstant{&grammarElementImpl{name, sheet_logic_types.StringConstant}, value}
}

type FloatConstant struct {
	*grammarElementImpl
	value float64
}

func (f *FloatConstant) CalculateFloat() (float64, error) {
	return f.value, nil
}

func NewFloatConstant(name string, value float64) *FloatConstant {
	return &FloatConstant{&grammarElementImpl{name, sheet_logic_types.FloatConstant}, value}
}

type BoolConstant struct {
	*grammarElementImpl
	value bool
}

func (b *BoolConstant) CalculateBool() (bool, error) {
	return b.value, nil
}

func NewBoolConstant(name string, value bool) *BoolConstant {
	return &BoolConstant{&grammarElementImpl{name, sheet_logic_types.BoolConstant}, value}
}
