package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestFloatCeil(t *testing.T) {
	uut := NewFloatCeil(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatCeil)

	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatCeil.CalculateFloat fails on uninitialized")

	uut.SetArg(NewFloatConstant(variableName, 3.6))
	assertCalculatesToFloat(
		t,
		uut,
		4.0,
		"FloatCeil.CalculateFloat")

	uut.SetArg(NewFloatConstant(variableName, 3.1))
	assertCalculatesToFloat(
		t,
		uut,
		4.0,
		"FloatCeil.CalculateFloat")
}
