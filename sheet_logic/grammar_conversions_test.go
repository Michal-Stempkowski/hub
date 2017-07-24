package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"math"
	"testing"
)

func TestIntToStringConversion(t *testing.T) {
	uut := NewIntToStringConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntToStringConversion)

	assertCalculatesToStringFails(
		t,
		uut,
		"IntToStringConversion.CalculateString fails on uninitialized")

	uut.SetArg(NewIntConstant(variableName, exampleIntValue))
	assertCalculatesToString(
		t,
		uut,
		exampleIntValueAsString,
		"IntToStringConversion.CalculateString")
}

func TestIntToFloatConversion(t *testing.T) {
	uut := NewIntToFloatConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntToFloatConversion)

	assertCalculatesToFloatFails(
		t,
		uut,
		"IntToFloatConversion.CalculateFloat fails on uninitialized")

	uut.SetArg(NewIntConstant(variableName, exampleIntValue))
	assertCalculatesToFloat(
		t,
		uut,
		exampleIntValue,
		"IntToFloatConversion.CalculateFloat")
}

func TestFloatToIntConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToIntConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntConversion.CalculateInt fails on uninitialized")

	uut.SetArg(floatArg)
	assertCalculatesToInt(
		t,
		uut,
		piRoundedDown,
		"FloatToIntConversion.CalculateInt for pi")

	floatArg.Value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedUp,
		"FloatToIntConversion.CalculateInt for e")

	floatArg.Value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntConversion.CalculateInt for big number")
}

func TestFloatToIntRoundDownConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntRoundDownConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToIntRoundDownConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntRoundDownConversion.CalculateInt fails on uninitialized")

	uut.SetArg(floatArg)
	assertCalculatesToInt(
		t,
		uut,
		piRoundedDown,
		"FloatToIntRoundDownConversion.CalculateInt for pi")

	floatArg.Value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedDown,
		"FloatToIntRoundDownConversion.CalculateInt for e")

	floatArg.Value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntRoundDownConversion.CalculateInt for big number")
}

func TestFloatToIntRoundUpConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntRoundUpConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToIntRoundUpConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntRoundUpConversion.CalculateInt fails on uninitialized")

	uut.SetArg(floatArg)
	assertCalculatesToInt(
		t,
		uut,
		piRoundedUp,
		"FloatToIntRoundUpConversion.CalculateInt for pi")

	floatArg.Value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedUp,
		"FloatToIntRoundUpConversion.CalculateInt for e")

	floatArg.Value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		"FloatToIntRoundUpConversion.CalculateInt for big number")
}

func TestStringToIntConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, exampleIntValueAsString)
	uut := NewStringToIntConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringToIntConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		"StringToIntConversion.CalculateInt fails on uninitialized")

	uut.SetArg(stringArg)
	assertCalculatesToInt(
		t,
		uut,
		exampleIntValue,
		"StringToIntConversion.CalculateInt some integer")

	stringArg.Value = bigNumberAsString
	assertCalculatesToIntFails(
		t,
		uut,
		"StringToIntConversion.CalculateInt for big number")

	stringArg.Value = stringPi
	assertCalculatesToIntFails(
		t,
		uut,
		"StringToIntConversion.CalculateInt for float number")
}

func TestStringToFloatConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, stringPiTwoDigitPrecision)
	uut := NewStringToFloatConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringToFloatConversion)

	assertCalculatesToFloatFails(
		t,
		uut,
		"StringToFloatConversion.CalculateFloat fails on uninitialized")

	uut.SetArg(stringArg)
	assertCalculatesToFloat(
		t,
		uut,
		piTwoDigitPrecision,
		"StringToFloatConversion.CalculateFloat for pi")

	stringArg.Value = exampleIntValueAsString
	assertCalculatesToFloat(
		t,
		uut,
		exampleIntValue,
		"StringToFloatConversion.CalculateFloat for int")

	stringArg.Value = exampleStringValue
	assertCalculatesToFloatFails(
		t,
		uut,
		"StringToFloatConversion.CalculateFloat for no number")
}

func TestFloatToStringConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, pi)
	uut := NewFloatToStringConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToStringConversion)

	assertCalculatesToStringFails(
		t,
		uut,
		"FloatToStringConversion.CalculateInt fails on uninitialized")

	uut.SetArg(floatArg)
	uut.Precision = 4
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
	uut := NewIntToBoolConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntToBoolConversion)

	assertCalculatesToBoolFails(
		t,
		uut,
		"IntToBoolConversion.CalculateBool fails on uninitialized")

	uut.SetArg(intArg)
	assertCalculatesToBool(
		t,
		uut,
		true,
		"IntToBoolConversion.CalculateBool for non zero number")

	intArg.Value = 0
	assertCalculatesToBool(
		t,
		uut,
		false,
		"IntToBoolConversion.CalculateBool for zero")
}

func TestFloatToBoolConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, 12.)
	uut := NewFloatToBoolConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToBoolConversion)

	assertCalculatesToBoolFails(
		t,
		uut,
		"FloatToBoolConversion.CalculateBool fails on uninitialized")

	uut.SetArg(floatArg)
	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatToBoolConversion.CalculateBool for non zero number")

	floatArg.Value = 0.
	assertCalculatesToBool(
		t,
		uut,
		false,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.Value = math.NaN()
	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.Value = math.Inf(1)
	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.Value = math.Inf(-1)
	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatToBoolConversion.CalculateBool for zero")
}

func TestStringToBoolConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, "true")
	uut := NewStringToBoolConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringToBoolConversion)

	assertCalculatesToBoolFails(
		t,
		uut,
		"StringToBoolConversion.CalculateBool fails on uninitialized")

	uut.SetArg(stringArg)
	assertCalculatesToBool(
		t,
		uut,
		true,
		"StringToBoolConversion.CalculateBool for 'true'")

	stringArg.Value = "false"
	assertCalculatesToBool(
		t,
		uut,
		false,
		"StringToBoolConversion.CalculateBool for 'false'")

	stringArg.Value = "Anything else"
	assertCalculatesToBoolFails(
		t,
		uut,
		"StringToBoolConversion.CalculateBool for anything else fails")
}

func TestBoolToIntConversion(t *testing.T) {
	boolArg := NewBoolConstant(variableName, true)
	uut := NewBoolToIntConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolToIntConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		"IntToStringConversion.CalculateInt fails on uninitialized")

	uut.SetArg(boolArg)
	assertCalculatesToInt(
		t,
		uut,
		1,
		"IntToStringConversion.CalculateInt")

	boolArg.Value = false
	assertCalculatesToInt(
		t,
		uut,
		0,
		"IntToStringConversion.CalculateInt")
}

func TestBoolToFloatConversion(t *testing.T) {
	boolArg := NewBoolConstant(variableName, true)
	uut := NewBoolToFloatConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolToFloatConversion)

	assertCalculatesToFloatFails(
		t,
		uut,
		"BoolToFloatConversion.CalculateFloat fails on uninitialized")

	uut.SetArg(boolArg)
	assertCalculatesToFloat(
		t,
		uut,
		1.,
		"BoolToFloatConversion.CalculateFloat")

	boolArg.Value = false
	assertCalculatesToFloat(
		t,
		uut,
		0.,
		"BoolToFloatConversion.CalculateFloat")
}

func TestBoolToStringConversion(t *testing.T) {
	boolArg := NewBoolConstant(variableName, true)
	uut := NewBoolToStringConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolToStringConversion)

	assertCalculatesToStringFails(
		t,
		uut,
		"BoolToStringConversion.CalculateString fails on uninitialized")

	uut.SetArg(boolArg)
	assertCalculatesToString(
		t,
		uut,
		"true",
		"BoolToStringConversion.CalculateString")

	boolArg.Value = false
	assertCalculatesToString(
		t,
		uut,
		"false",
		"BoolToStringConversion.CalculateString")
}
