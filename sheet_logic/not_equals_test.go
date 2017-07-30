package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntNotEquals(t *testing.T) {
	uut := NewIntNotEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntNotEquals)

	lesser, equal, greater := true, false, true
	IntComparatorScenario(
		t,
		(*IntComparator)(uut),
		"IntNotEquals",
		lesser,
		equal,
		greater)
}

func TestFloatNotEquals(t *testing.T) {
	uut := NewFloatNotEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatNotEquals)

	lesser, equal, greater := true, false, true
	FloatComparatorScenario(
		t,
		(*FloatComparator)(uut),
		"FloatNotEquals",
		lesser,
		equal,
		greater)
}

func TestBoolNotEquals(t *testing.T) {
	uut := NewBoolNotEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolNotEquals)

	BoolComparatorScenario(
		t,
		(*BoolComparator)(uut),
		"BoolNotEquals",
		[]bool{false, true, true, false})
}

func TestStringNotEquals(t *testing.T) {
	uut := NewStringNotEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringNotEquals)

	lesser, equal, greater := true, false, true
	StringComparatorScenario(
		t,
		(*StringComparator)(uut),
		"StringNotEquals",
		lesser,
		equal,
		greater)
}
