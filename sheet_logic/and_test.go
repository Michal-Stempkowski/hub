package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestAnd(t *testing.T) {
	uut := NewAnd(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.And)

	BoolComparatorScenario(
		t,
		(*BoolComparator)(uut),
		"And",
		[]bool{false, false, false, true})
}
