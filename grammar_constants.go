package sheet_logic

import "hub/sheet_logic_types"

type IntConstant struct {
	*grammarElementImpl
	value int
}

func (i *IntConstant) CalculateInt() (int, error) {
	return i.value, nil
}

func NewIntConstant(name string, value int) *IntConstant {
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
	value float32
}

func (f *FloatConstant) CalculateFloat() (float32, error) {
	return f.value, nil
}

func NewFloatConstant(name string, value float32) *FloatConstant {
	return &FloatConstant{&grammarElementImpl{name, sheet_logic_types.FloatConstant}, value}
}
