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
