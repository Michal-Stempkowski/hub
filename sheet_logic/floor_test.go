package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestFloatFloor(t *testing.T) {
	uut := NewFloatFloor(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatFloor)

	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatFloor.CalculateFloat fails on uninitialized")

	uut.SetArg(NewFloatConstant(variableName, 3.6))
	assertCalculatesToFloat(
		t,
		uut,
		3.0,
		"FloatFloor.CalculateFloat")

	uut.SetArg(NewFloatConstant(variableName, 3.1))
	assertCalculatesToFloat(
		t,
		uut,
		3.0,
		"FloatFloor.CalculateFloat")
}
