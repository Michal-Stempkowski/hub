package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntSum(t *testing.T) {
	argLeft := NewIntConstant(variableName, 2)
	argRight := NewIntConstant(variableName, 3)
	uut := NewIntSum(variableName, argLeft, argRight)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.IntSum)

	assertCalculatesToInt(
		t,
		uut,
		5,
		"IntSum.CalculateInt")

	// uut.argLeft = nil
	// assertCalculatesToIntFails(
	// 	t,
	// 	uut,
	// 	"IntSum.CalculateInt fails when argLeft missing")

	// uut.argLeft = NewIntConstant(variableName, 2)
	// uut.argRight = nil
	// assertCalculatesToIntFails(
	// 	t,
	// 	uut,
	// 	"IntSum.CalculateInt fails when argRight missing")
}
