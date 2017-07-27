package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestFloatRound(t *testing.T) {
	uut := NewFloatRound(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatRound)

	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatRound.CalculateFloat fails on uninitialized")

	uut.SetArg(NewFloatConstant(variableName, 3.6))
	assertCalculatesToFloat(
		t,
		uut,
		4.0,
		"FloatRound.CalculateFloat")

	uut.SetArg(NewFloatConstant(variableName, 3.1))
	assertCalculatesToFloat(
		t,
		uut,
		3.0,
		"FloatRound.CalculateFloat")
}
