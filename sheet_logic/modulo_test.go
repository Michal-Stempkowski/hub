package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntModulo(t *testing.T) {
	uut := NewIntModulo(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntModulo)

	uut.SetLeftArg(NewIntConstant(variableName, 5))
	uut.SetRightArg(NewIntConstant(variableName, 2))
	assertCalculatesToInt(
		t,
		uut,
		1,
		"IntModulo.CalculateInt")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		"IntModulo.CalculateInt fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 5))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		"IntModulo.CalculateInt fails when argRight missing")

	uut.SetRightArg(NewIntConstant(variableName, 0))
	assertCalculatesToIntFails(
		t,
		uut,
		"IntModulo.CalculateInt fails when dividing by zero")
}
