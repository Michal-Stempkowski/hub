package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntLesserEqual(t *testing.T) {
	uut := NewIntLesserEqual(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntLesserEqual)

	lesser, equal, greater := true, true, false
	IntComparatorScenario(
		t,
		(*IntComparator)(uut),
		"IntLesserEqual",
		lesser,
		equal,
		greater)
}

func TestFloatLesserEqual(t *testing.T) {
	uut := NewFloatLesserEqual(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatLesserEqual)

	lesser, equal, greater := true, true, false
	FloatComparatorScenario(
		t,
		(*FloatComparator)(uut),
		"FloatLesserEqual",
		lesser,
		equal,
		greater)
}

func TestStringLesserEqual(t *testing.T) {
	uut := NewStringLesserEqual(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringLesserEqual)

	lesser, equal, greater := true, true, false
	StringComparatorScenario(
		t,
		(*StringComparator)(uut),
		"StringLesserEqual",
		lesser,
		equal,
		greater)
}
