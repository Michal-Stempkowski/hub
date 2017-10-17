package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntMultiplication(t *testing.T) {
	uut := NewIntMultiplication(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntMultiplication)

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewIntConstant(variableName, 3))
	assertCalculatesToInt(
		t,
		uut,
		6,
		noGrammarContext,
		"IntMultiplication.CalculateInt")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"IntMultiplication.CalculateInt fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"IntMultiplication.CalculateInt fails when argRight missing")
}

func TestFloatMultiplication(t *testing.T) {
	uut := NewFloatMultiplication(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatMultiplication)

	uut.SetLeftArg(NewFloatConstant(variableName, 2.5))
	uut.SetRightArg(NewFloatConstant(variableName, 1.5))
	assertCalculatesToFloat(
		t,
		uut,
		3.75,
		noGrammarContext,
		"FloatMultiplication.CalculateFloat")

	uut.SetLeftArg(NewEmptyFloatExpression())
	assertCalculatesToFloatFails(
		t,
		uut,
		noGrammarContext,
		"FloatMultiplication.CalculateFloat fails when argLeft missing")

	uut.SetLeftArg(NewFloatConstant(variableName, 2.5))
	uut.SetRightArg(NewEmptyFloatExpression())
	assertCalculatesToFloatFails(
		t,
		uut,
		noGrammarContext,
		"FloatMultiplication.CalculateFloat fails when argRight missing")
}
