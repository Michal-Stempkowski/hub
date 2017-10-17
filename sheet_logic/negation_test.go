package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestNegation(t *testing.T) {
	boolArg := NewBoolConstant(variableName, true)
	uut := NewNegation(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.Negation)

	assertCalculatesToBoolFails(
		t,
		uut,
		noGrammarContext,
		"Negation.CalculateBool fails on uninitialized")

	uut.SetArg(boolArg)
	assertCalculatesToBool(
		t,
		uut,
		false,
		noGrammarContext,
		"Negation.CalculateBool")

	boolArg.Value = false
	assertCalculatesToBool(
		t,
		uut,
		true,
		noGrammarContext,
		"Negation.CalculateBool")
}
