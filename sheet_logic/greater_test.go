package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntGreater(t *testing.T) {
	uut := NewIntGreater(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntGreater)

	lesser, equal, greater := false, false, true
	IntComparatorScenario(
		t,
		(*IntComparator)(uut),
		"IntGreater",
		lesser,
		equal,
		greater)
}

func TestFloatGreater(t *testing.T) {
	uut := NewFloatGreater(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatGreater)

	lesser, equal, greater := false, false, true
	FloatComparatorScenario(
		t,
		(*FloatComparator)(uut),
		"FloatGreater",
		lesser,
		equal,
		greater)
}

func TestStringGreater(t *testing.T) {
	uut := NewStringGreater(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringGreater)

	lesser, equal, greater := false, false, true
	StringComparatorScenario(
		t,
		(*StringComparator)(uut),
		"StringGreater",
		lesser,
		equal,
		greater)
}
