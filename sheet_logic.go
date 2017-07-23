package sheet_logic

import "hub/sheet_logic_types"

type GrammarElement interface {
	GetType() sheet_logic_types.T
	GetName() string
	SetName(string)
}

// Because there will be also volatile sources (formulas, textfields etc.)
type IntExpresion interface {
	CalculateInt() int
}

type FloatExpresion interface {
	CalculateFloat() float32
}

type StringExpresion interface {
	CalculateString() string
}

type grammarElementImpl struct {
	name         string
	grammar_type sheet_logic_types.T
}

func (g *grammarElementImpl) GetName() string {
	return g.name
}

func (g *grammarElementImpl) SetName(newName string) {
	g.name = newName
}

func (g *grammarElementImpl) GetType() sheet_logic_types.T {
	return g.grammar_type
}

type IntConstant struct {
	*grammarElementImpl
	value int
}

func (i *IntConstant) CalculateInt() int {
	return i.value
}

func NewIntConstant(name string, value int) *IntConstant {
	return &IntConstant{&grammarElementImpl{name, sheet_logic_types.IntConstant}, value}
}

type StringConstant struct {
	*grammarElementImpl
	value string
}

func (s *StringConstant) CalculateString() string {
	return s.value
}

func NewStringConstant(name string, value string) *StringConstant {
	return &StringConstant{&grammarElementImpl{name, sheet_logic_types.StringConstant}, value}
}

type FloatConstant struct {
	*grammarElementImpl
	value float32
}

func (f *FloatConstant) CalculateFloat() float32 {
	return f.value
}

func NewFloatConstant(name string, value float32) *FloatConstant {
	return &FloatConstant{&grammarElementImpl{name, sheet_logic_types.FloatConstant}, value}
}
