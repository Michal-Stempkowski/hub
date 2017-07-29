package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntEquals(t *testing.T) {
	uut := NewIntEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntEquals)

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewIntConstant(variableName, 3))
	assertCalculatesToBool(
		t,
		uut,
		false,
		"IntEquals.CalculateBool")

	uut.SetRightArg(NewIntConstant(variableName, 2))
	assertCalculatesToBool(
		t,
		uut,
		true,
		"IntEquals.CalculateBool")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"IntEquals.CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"IntEquals.CalculateBool fails when argRight missing")
}
