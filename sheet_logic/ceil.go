package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"math"
)

type FloatCeil struct {
	GrammarElement
	UnaryOperationFloat
}

func (f *FloatCeil) CalculateFloat() (result float64, err error) {
	var arg float64

	if arg, err = f.GetArg().CalculateFloat(); err == nil {
		result = math.Ceil(arg)
	}

	return
}

func NewFloatCeil(name string) *FloatCeil {
	return &FloatCeil{
		&grammarElementImpl{name, sheet_logic_types.FloatCeil},
		DefaultUnaryOperationFloatImpl()}
}
