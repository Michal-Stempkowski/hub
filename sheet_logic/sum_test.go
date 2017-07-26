package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntSum(t *testing.T) {
	uut := NewIntSum(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntSum)

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewIntConstant(variableName, 3))
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

func TestFloatSum(t *testing.T) {
	uut := NewFloatSum(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatSum)

	uut.SetLeftArg(NewFloatConstant(variableName, 2.5))
	uut.SetRightArg(NewFloatConstant(variableName, 3.2))
	assertCalculatesToFloat(
		t,
		uut,
		5.7,
		"FloatSum.CalculateFloat")

	uut.SetLeftArg(NewEmptyFloatExpression())
	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatSum.CalculateFloat fails when argLeft missing")

	uut.SetLeftArg(NewFloatConstant(variableName, 2.5))
	uut.SetRightArg(NewEmptyFloatExpression())
	assertCalculatesToFloatFails(
		t,
		uut,
		"FloatSum.CalculateFloat fails when argRight missing")
}

func TestStringSum(t *testing.T) {
	uut := NewStringSum(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringSum)

	uut.SetLeftArg(NewStringConstant(variableName, "ab"))
	uut.SetRightArg(NewStringConstant(variableName, "c"))
	assertCalculatesToString(
		t,
		uut,
		"abc",
		"StringSum.CalculateString")

	uut.SetLeftArg(NewEmptyStringExpression())
	assertCalculatesToStringFails(
		t,
		uut,
		"StringSum.CalculateString fails when argLeft missing")

	uut.SetLeftArg(NewStringConstant(variableName, "ab"))
	uut.SetRightArg(NewEmptyStringExpression())
	assertCalculatesToStringFails(
		t,
		uut,
		"StringSum.CalculateString fails when argRight missing")
}
