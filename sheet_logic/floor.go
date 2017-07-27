package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"math"
)

type FloatFloor struct {
	GrammarElement
	UnaryOperationFloat
}

func (f *FloatFloor) CalculateFloat() (result float64, err error) {
	var arg float64

	if arg, err = f.GetArg().CalculateFloat(); err == nil {
		result = math.Floor(arg)
	}

	return
}

func NewFloatFloor(name string) *FloatFloor {
	return &FloatFloor{
		&grammarElementImpl{name, sheet_logic_types.FloatFloor},
		DefaultUnaryOperationFloatImpl()}
}
