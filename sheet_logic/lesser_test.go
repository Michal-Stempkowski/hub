package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntLesser(t *testing.T) {
	uut := NewIntLesser(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntLesser)

	lesser, equal, greater := true, false, false
	IntComparatorScenario(
		t,
		(*IntComparator)(uut),
		"IntLesser",
		lesser,
		equal,
		greater)
}

func TestFloatLesser(t *testing.T) {
	uut := NewFloatLesser(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatLesser)

	lesser, equal, greater := true, false, false
	FloatComparatorScenario(
		t,
		(*FloatComparator)(uut),
		"FloatLesser",
		lesser,
		equal,
		greater)
}
