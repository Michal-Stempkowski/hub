package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
)

type FloatRound struct {
	GrammarElement
	UnaryOperationFloat
}

func (f *FloatRound) CalculateFloat(g GrammarContext) (result float64, err error) {
	var arg float64

	if arg, err = f.GetArg().CalculateFloat(g); err == nil {
		result = framework.Round(arg)
	}

	return
}

func NewFloatRound(name string) *FloatRound {
	return &FloatRound{
		&grammarElementImpl{name, sheet_logic_types.FloatRound},
		DefaultUnaryOperationFloatImpl()}
}
