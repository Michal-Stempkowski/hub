package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestOr(t *testing.T) {
	uut := NewOr(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.Or)

	BoolComparatorScenario(
		t,
		(*BoolComparator)(uut),
		"Or",
		[]bool{false, true, true, true})
}
