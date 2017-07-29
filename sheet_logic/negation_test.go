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
		"Negation.CalculateBool fails on uninitialized")

	uut.SetArg(boolArg)
	assertCalculatesToBool(
		t,
		uut,
		false,
		"Negation.CalculateBool")

	boolArg.Value = false
	assertCalculatesToBool(
		t,
		uut,
		true,
		"Negation.CalculateBool")
}
