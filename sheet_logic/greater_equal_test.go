package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntGreaterEqual(t *testing.T) {
	uut := NewIntGreaterEqual(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntGreaterEqual)

	lesser, equal, greater := false, true, true
	IntComparatorScenario(
		t,
		(*IntComparator)(uut),
		noGrammarContext,
		"IntGreaterEqual",
		lesser,
		equal,
		greater)
}

func TestFloatGreaterEqual(t *testing.T) {
	uut := NewFloatGreaterEqual(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatGreaterEqual)

	lesser, equal, greater := false, true, true
	FloatComparatorScenario(
		t,
		(*FloatComparator)(uut),
		noGrammarContext,
		"FloatGreaterEqual",
		lesser,
		equal,
		greater)
}

func TestStringGreaterEqual(t *testing.T) {
	uut := NewStringGreaterEqual(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringGreaterEqual)

	lesser, equal, greater := false, true, true
	StringComparatorScenario(
		t,
		(*StringComparator)(uut),
		noGrammarContext,
		"StringGreaterEqual",
		lesser,
		equal,
		greater)
}
