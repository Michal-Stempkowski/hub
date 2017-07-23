package sheet_logic

import (
	"hub/sheet_logic_types"
	"math"
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

func TestFloatToIntRoundDownConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntRoundDownConversion(variableName, floatArg)

	grammarElementScenario(
		t, uut.grammarElementImpl, sheet_logic_types.FloatToIntRoundDownConversion)

	assertCalculatesToInt(
		t,
		uut,
		piRoundedDown,
		"FloatToIntRoundDownConversion.CalculateInt for pi")

	floatArg.value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedDown,
		"FloatToIntRoundDownConversion.CalculateInt for e")

	floatArg.value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntRoundDownConversion.CalculateInt for big number")
}

func TestFloatToIntRoundUpConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntRoundUpConversion(variableName, floatArg)

	grammarElementScenario(
		t, uut.grammarElementImpl, sheet_logic_types.FloatToIntRoundUpConversion)

	assertCalculatesToInt(
		t,
		uut,
		piRoundedUp,
		"FloatToIntRoundUpConversion.CalculateInt for pi")

	floatArg.value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedUp,
		"FloatToIntRoundUpConversion.CalculateInt for e")

	floatArg.value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntRoundUpConversion.CalculateInt for big number")
}

func TestStringToIntConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, exampleIntValueAsString)
	uut := NewStringToIntConversion(variableName, stringArg)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.StringToIntConversion)

	assertCalculatesToInt(
		t,
		uut,
		exampleIntValue,
		"StringToIntConversion.CalculateInt some integer")

	stringArg.value = bigNumberAsString
	assertCalculatesToIntFails(
		t,
		uut,
		"StringToIntConversion.CalculateInt for big number")

	stringArg.value = stringPi
	assertCalculatesToIntFails(
		t,
		uut,
		"StringToIntConversion.CalculateInt for float number")
}

func TestStringToFloatConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, stringPiTwoDigitPrecision)
	uut := NewStringToFloatConversion(variableName, stringArg)

	grammarElementScenario(
		t, uut.grammarElementImpl, sheet_logic_types.StringToFloatConversion)

	assertCalculatesToFloat(
		t,
		uut,
		piTwoDigitPrecision,
		"StringToFloatConversion.CalculateFloat for pi")

	stringArg.value = exampleIntValueAsString
	assertCalculatesToFloat(
		t,
		uut,
		exampleIntValue,
		"StringToFloatConversion.CalculateFloat for int")

	stringArg.value = exampleStringValue
	assertCalculatesToFloatFails(
		t,
		uut,
		"StringToFloatConversion.CalculateFloat for no number")
}

func TestFloatToStringConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, pi)
	uut := NewFloatToStringConversion(variableName, floatArg, 4)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.FloatToStringConversion)

	assertCalculatesToString(
		t,
		uut,
		stringPi,
		"FloatToStringConversion.CalculateInt for pi")

	uut.Precision = 2
	assertCalculatesToString(
		t,
		uut,
		stringPiTwoDigitPrecision,
		"FloatToStringConversion.CalculateInt for e")

	uut.Precision = 10
	assertCalculatesToString(
		t,
		uut,
		stringPi,
		"FloatToStringConversion.CalculateInt for pi")

	uut.Precision = -1
	assertCalculatesToStringFails(
		t,
		uut,
		"FloatToStringConversion.CalculateInt for big number")
}

func TestIntToBoolConversion(t *testing.T) {
	intArg := NewIntConstant(variableName, 5)
	uut := NewIntToBoolConversion(variableName, intArg)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.IntToBoolConversion)

	assertCalculatesToBool(
		t,
		uut,
		true,
		"IntToBoolConversion.CalculateBool for non zero number")

	intArg.value = 0
	assertCalculatesToBool(
		t,
		uut,
		false,
		"IntToBoolConversion.CalculateBool for zero")
}

func TestFloatToBoolConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, 12.)
	uut := NewFloatToBoolConversion(variableName, floatArg)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.FloatToBoolConversion)

	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatToBoolConversion.CalculateBool for non zero number")

	floatArg.value = 0.
	assertCalculatesToBool(
		t,
		uut,
		false,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.value = math.NaN()
	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.value = math.Inf(1)
	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.value = math.Inf(-1)
	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatToBoolConversion.CalculateBool for zero")
}

func TestStringToBoolConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, "true")
	uut := NewStringToBoolConversion(variableName, stringArg)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.StringToBoolConversion)

	assertCalculatesToBool(
		t,
		uut,
		true,
		"StringToBoolConversion.CalculateBool for 'true'")

	stringArg.value = "false"
	assertCalculatesToBool(
		t,
		uut,
		false,
		"StringToBoolConversion.CalculateBool for 'false'")

	stringArg.value = "Anything else"
	assertCalculatesToBoolFails(
		t,
		uut,
		"StringToBoolConversion.CalculateBool for anything else fails")
}
