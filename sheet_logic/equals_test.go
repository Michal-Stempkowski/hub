package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntEquals(t *testing.T) {
	uut := NewIntEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntEquals)

	lesser, equal, greater := false, true, false
	IntComparatorScenario(
		t,
		(*IntComparator)(uut),
		"IntEquals",
		lesser,
		equal,
		greater)
}

func TestFloatEquals(t *testing.T) {
	uut := NewFloatEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatEquals)

	lesser, equal, greater := false, true, false
	FloatComparatorScenario(
		t,
		(*FloatComparator)(uut),
		"FloatEquals",
		lesser,
		equal,
		greater)
}

func TestBoolEquals(t *testing.T) {
	uut := NewBoolEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolEquals)

	BoolComparatorScenario(
		t,
		(*BoolComparator)(uut),
		"BoolEquals",
		[]bool{true, false, false, true})
}

func TestStringEquals(t *testing.T) {
	uut := NewStringEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringEquals)

	lesser, equal, greater := false, true, false
	StringComparatorScenario(
		t,
		(*StringComparator)(uut),
		"StringEquals",
		lesser,
		equal,
		greater)
}
