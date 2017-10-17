package sheet_logic

import "testing"
import "hub/sheet_logic/sheet_logic_types"

func TestShouldBeAbleToCreateIntConstant(t *testing.T) {
	uut := NewIntConstant(variableName, exampleIntValue)

	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntConstant)
	assertCalculatesToInt(
		t,
		uut,
		exampleIntValue,
		noGrammarContext,
		"IntConstant.CalculateInt")
}

func TestShouldBeAbleToCreateStringConstant(t *testing.T) {
	uut := NewStringConstant(variableName, exampleStringValue)

	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringConstant)
	assertCalculatesToString(
		t,
		uut,
		exampleStringValue,
		noGrammarContext,
		"StringConstant.CalculateString")
}

func TestShouldBeAbleToCreateFloatConstant(t *testing.T) {
	uut := NewFloatConstant(variableName, exampleFloatValue)

	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatConstant)
	assertCalculatesToFloat(
		t,
		uut,
		exampleFloatValue,
		noGrammarContext,
		"FloatConstant.CalculateFloat")
}

func TestShouldBeAbleToCreateBoolConstant(t *testing.T) {
	uut := NewBoolConstant(variableName, true)

	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolConstant)
	assertCalculatesToBool(
		t,
		uut,
		true,
		noGrammarContext,
		"BoolConstant.CalculateBool")
}
