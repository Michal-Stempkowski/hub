package sheet_logic

import "testing"

const (
	emptyIdentifier            = ""
	inputIdentifier            = "Attributes.Strength.Base"
	inputIdentifierNotExisting = "Attributes.Strength.Modifier"
)

func TestIntInput(t *testing.T) {
	c := noGrammarContext

	uut := NewIntInput(variableName)
	assertCalculatesToIntFails(
		t, uut, c, "TestIntInput (grammar context is needed)")

	c = newDummyGrammarContext()
	// grammar context may or may not handle empty identifiers - IntInput verifies
	c.expectedInts[emptyIdentifier] = 2
	c.expectedInts[inputIdentifier] = 5

	uut.Identifier = inputIdentifierNotExisting
	assertCalculatesToIntFails(
		t, uut, c, "TestIntInput (unknown identifier)")

	uut.Identifier = emptyIdentifier
	assertCalculatesToIntFails(
		t, uut, c, "TestIntInput (empty identifier)")

	uut.Identifier = inputIdentifier
	assertCalculatesToInt(t, uut, 5, c, "TestIntInput (ok case)")
}
