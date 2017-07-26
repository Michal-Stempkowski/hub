package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntDivision(t *testing.T) {
	uut := NewIntDivision(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntDivision)

	uut.SetLeftArg(NewIntConstant(variableName, 3))
	uut.SetRightArg(NewIntConstant(variableName, 2))
	assertCalculatesToInt(
		t,
		uut,
		1,
		"IntDivision.CalculateInt")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		"IntDivision.CalculateInt fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 3))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		"IntDivision.CalculateInt fails when argRight missing")

	uut.SetRightArg(NewIntConstant(variableName, 0))
	assertCalculatesToIntFails(
		t,
		uut,
		"IntDivision.CalculateInt fails when dividing by zero")
}

func TestFloatDivision(t *testing.T) {
	uut := NewFloatDivision(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatDivision)

	uut.SetLeftArg(NewFloatConstant(variableName, 7.5))
	uut.SetRightArg(NewFloatConstant(variableName, 2.5))
	assertCalculatesToFloat(
		t,
		uut,
		3.0,
		"FloatDivision.CalculateFloat")

	uut.SetLeftArg(NewEmptyFloatExpression())
	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatDivision.CalculateFloat fails when argLeft missing")

	uut.SetLeftArg(NewFloatConstant(variableName, 7.5))
	uut.SetRightArg(NewEmptyFloatExpression())
	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatDivision.CalculateFloat fails when argRight missing")

	uut.SetRightArg(NewFloatConstant(variableName, 0.0))
	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatDivision.CalculateFloat fails when dividing by zero")
}
