package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
)

type Negation struct {
	GrammarElement
	UnaryOperationBool
}

func (n *Negation) CalculateBool() (result bool, err error) {
	var boolVal bool

	if boolVal, err = n.GetArg().CalculateBool(); err == nil {
		result = !boolVal
	}

	return
}

func NewNegation(name string) *Negation {
	return &Negation{
		&grammarElementImpl{name, sheet_logic_types.Negation},
		DefaultUnaryOperationBoolImpl()}
}
