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

func TestStringInput(t *testing.T) {
	uut := NewStringInput(variableName)

	c := noGrammarContext
	uut.Identifier = inputIdentifier
	assertCalculatesToStringFails(
		t, uut, c, "TestStringInput (grammar context is needed)")

	c = newDummyGrammarContext()
	// grammar context may or may not handle empty identifiers - StringInput verifies
	c.expectedStrings[emptyIdentifier] = "a"
	c.expectedStrings[inputIdentifier] = "b"

	uut.Identifier = inputIdentifierNotExisting
	assertCalculatesToStringFails(
		t, uut, c, "TestStringInput (unknown identifier)")

	uut.Identifier = emptyIdentifier
	assertCalculatesToStringFails(
		t, uut, c, "TestStringInput (empty identifier)")

	uut.Identifier = inputIdentifier
	assertCalculatesToString(t, uut, "b", c, "TestStringInput (ok case)")
}

func TestBoolInput(t *testing.T) {
	uut := NewBoolInput(variableName)

	c := noGrammarContext
	uut.Identifier = inputIdentifier
	assertCalculatesToBoolFails(
		t, uut, c, "TestBoolInput (grammar context is needed)")

	c = newDummyGrammarContext()
	// grammar context may or may not handle empty identifiers - BoolInput verifies
	c.expectedBools[emptyIdentifier] = false
	c.expectedBools[inputIdentifier] = true

	uut.Identifier = inputIdentifierNotExisting
	assertCalculatesToBoolFails(
		t, uut, c, "TestBoolInput (unknown identifier)")

	uut.Identifier = emptyIdentifier
	assertCalculatesToBoolFails(
		t, uut, c, "TestBoolInput (empty identifier)")

	uut.Identifier = inputIdentifier
	assertCalculatesToBool(t, uut, true, c, "TestBoolInput (ok case)")
}
