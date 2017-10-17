package sheet_logic

import "testing"

const (
	emptyIdentifier            = ""
	inputIdentifier            = "Attributes.Strength.Base"
	inputIdentifierNotExisting = "Attributes.Strength.Modifier"
)

func TestIntInput(t *testing.T) {
	uut := NewIntInput(variableName)

	c := noGrammarContext
	uut.Identifier = inputIdentifier
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

func TestFloatInput(t *testing.T) {
	uut := NewFloatInput(variableName)

	c := noGrammarContext
	uut.Identifier = inputIdentifier
	assertCalculatesToFloatFails(
		t, uut, c, "TestFloatInput (grammar context is needed)")

	c = newDummyGrammarContext()
	// grammar context may or may not handle empty identifiers - FloatInput verifies
	c.expectedFloats[emptyIdentifier] = 2.0
	c.expectedFloats[inputIdentifier] = 5.0

	uut.Identifier = inputIdentifierNotExisting
	assertCalculatesToFloatFails(
		t, uut, c, "TestFloatInput (unknown identifier)")

	uut.Identifier = emptyIdentifier
	assertCalculatesToFloatFails(
		t, uut, c, "TestFloatInput (empty identifier)")

	uut.Identifier = inputIdentifier
	assertCalculatesToFloat(t, uut, 5.0, c, "TestFloatInput (ok case)")
}
