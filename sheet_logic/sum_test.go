package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntSum(t *testing.T) {
	uut := NewIntSum(variableName)
	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewIntConstant(variableName, 3))

	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntSum)

	assertCalculatesToInt(
		t,
		uut,
		5,
		"IntSum.CalculateInt")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		"IntSum.CalculateInt fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		"IntSum.CalculateInt fails when argRight missing")
}
