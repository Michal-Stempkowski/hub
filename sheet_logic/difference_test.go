package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntDifference(t *testing.T) {
	uut := NewIntDifference(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntDifference)

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewIntConstant(variableName, 3))
	assertCalculatesToInt(
		t,
		uut,
		-1,
		"IntDifference.CalculateInt")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		"IntDifference.CalculateInt fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		"IntDifference.CalculateInt fails when argRight missing")
}

func TestFloatDifference(t *testing.T) {
	uut := NewFloatDifference(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatDifference)

	uut.SetLeftArg(NewFloatConstant(variableName, 2.5))
	uut.SetRightArg(NewFloatConstant(variableName, 3.2))
	assertCalculatesToFloat(
		t,
		uut,
		-0.7,
		"FloatDifference.CalculateFloat")

	uut.SetLeftArg(NewEmptyFloatExpression())
	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatDifference.CalculateFloat fails when argLeft missing")

	uut.SetLeftArg(NewFloatConstant(variableName, 2.5))
	uut.SetRightArg(NewEmptyFloatExpression())
	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatDifference.CalculateFloat fails when argRight missing")
}
