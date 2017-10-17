package sheet_logic

import (
	"fmt"
	"hub/sheet_logic/sheet_logic_types"
)

type IntInput struct {
	GrammarElement
	Identifier string
}

func (i *IntInput) CalculateInt(g GrammarContext) (int64, error) {
	if i.Identifier == "" {
		return 0, fmt.Errorf("IntInput: Empty identifier not allowed!")
	}
	return g.GetIntValue(i.Identifier)
}

func NewIntInput(name string) *IntInput {
	return &IntInput{
		&grammarElementImpl{name, sheet_logic_types.IntInput},
		""}
}

type FloatInput struct {
	GrammarElement
	Identifier string
}

func (f *FloatInput) CalculateFloat(g GrammarContext) (float64, error) {
	if f.Identifier == "" {
		return 0.0, fmt.Errorf("FloatInput: Empty identifier not allowed!")
	}
	return g.GetFloatValue(f.Identifier)
}

func NewFloatInput(name string) *FloatInput {
	return &FloatInput{
		&grammarElementImpl{name, sheet_logic_types.FloatInput},
		""}
}

type StringInput struct {
	GrammarElement
	Identifier string
}

func (f *StringInput) CalculateString(g GrammarContext) (string, error) {
	if f.Identifier == "" {
		return "", fmt.Errorf("StringInput: Empty identifier not allowed!")
	}
	return g.GetStringValue(f.Identifier)
}

func NewStringInput(name string) *StringInput {
	return &StringInput{
		&grammarElementImpl{name, sheet_logic_types.StringInput},
		""}
}
