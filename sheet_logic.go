package sheet_logic

import "hub/sheet_logic_types"

type GrammarElement interface {
	GetType() sheet_logic_types.T
	GetName() string
	SetName(string)
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
	Value int
}

func NewIntConstant(name string, value int) *IntConstant {
	return &IntConstant{&grammarElementImpl{name, sheet_logic_types.IntConstant}, value}
}

type StringConstant struct {
	*grammarElementImpl
	Value string
}

func NewStringConstant(name string, value string) *StringConstant {
	return &StringConstant{&grammarElementImpl{name, sheet_logic_types.StringConstant}, value}
}

type FloatConstant struct {
	*grammarElementImpl
	Value float32
}

func NewFloatConstant(name string, value float32) *FloatConstant {
	return &FloatConstant{&grammarElementImpl{name, sheet_logic_types.FloatConstant}, value}
}
