package sheet_logic

import (
	"hub/sheet_logic_types"
	"testing"
)

func TestIntToStringConversion(t *testing.T) {
	intArg := NewIntConstant(variableName, exampleIntValue)
	uut := NewIntToStringConversion(variableName, intArg)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.IntToStringConversion)

	assertCalculatesToString(
		t,
		uut,
		exampleIntValueAsString,
		"IntToStringConversion.CalculateString")
}

func TestIntToFloatConversion(t *testing.T) {
	intArg := NewIntConstant(variableName, exampleIntValue)
	uut := NewIntToFloatConversion(variableName, intArg)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.IntToFloatConversion)

	assertCalculatesToFloat(
		t,
		uut,
		exampleIntValue,
		"IntToFloatConversion.CalculateFloat")
}

func TestFloatToIntConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntConversion(variableName, floatArg)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.FloatToIntConversion)

	assertCalculatesToInt(
		t,
		uut,
		piRoundedDown,
		"FloatToIntConversion.CalculateInt for pi")

	floatArg.value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedUp,
		"FloatToIntConversion.CalculateInt for e")

	floatArg.value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntConversion.CalculateInt for big number")
}
