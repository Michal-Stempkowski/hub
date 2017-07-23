package sheet_logic

import (
	"hub/sheet_logic_types"
	"testing"
)

func TestIntToStringConversion(t *testing.T) {
	intArg := NewIntConstant(variableName, exampleIntValue)
	uut := NewIntToStringConversion(variableName, intArg)
	assertHasName(t, uut, variableName)
	assertCalculatesToString(
		t,
		uut,
		exampleIntValueAsString,
		"IntToStringConversion.CalculateString")

	assertHasType(t, uut, sheet_logic_types.IntToStringConversion)

	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}

func TestIntToFloatConversion(t *testing.T) {
	intArg := NewIntConstant(variableName, exampleIntValue)
	uut := NewIntToFloatConversion(variableName, intArg)
	assertHasName(t, uut, variableName)
	assertCalculatesToFloat(
		t,
		uut,
		exampleIntValue,
		"IntToFloatConversion.CalculateFloat")

	assertHasType(t, uut, sheet_logic_types.IntToFloatConversion)

	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}
