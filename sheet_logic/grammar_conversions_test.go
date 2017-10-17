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
		noGrammarContext,
		"IntToStringConversion.CalculateString fails on uninitialized")

	uut.SetArg(NewIntConstant(variableName, exampleIntValue))
	assertCalculatesToString(
		t,
		uut,
		exampleIntValueAsString,
		noGrammarContext,
		"IntToStringConversion.CalculateString")
}

func TestIntToFloatConversion(t *testing.T) {
	uut := NewIntToFloatConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntToFloatConversion)

	assertCalculatesToFloatFails(
		t,
		uut,
		noGrammarContext,
		"IntToFloatConversion.CalculateFloat fails on uninitialized")

	uut.SetArg(NewIntConstant(variableName, exampleIntValue))
	assertCalculatesToFloat(
		t,
		uut,
		exampleIntValue,
		noGrammarContext,
		"IntToFloatConversion.CalculateFloat")
}

func TestFloatToIntConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToIntConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"FloatToIntConversion.CalculateInt fails on uninitialized")

	uut.SetArg(floatArg)
	assertCalculatesToInt(
		t,
		uut,
		piRoundedDown,
		noGrammarContext,
		"FloatToIntConversion.CalculateInt for pi")

	floatArg.Value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedUp,
		noGrammarContext,
		"FloatToIntConversion.CalculateInt for e")

	floatArg.Value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"FloatToIntConversion.CalculateInt for big number")
}

func TestFloatToIntRoundDownConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntRoundDownConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToIntRoundDownConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"FloatToIntRoundDownConversion.CalculateInt fails on uninitialized")

	uut.SetArg(floatArg)
	assertCalculatesToInt(
		t,
		uut,
		piRoundedDown,
		noGrammarContext,
		"FloatToIntRoundDownConversion.CalculateInt for pi")

	floatArg.Value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedDown,
		noGrammarContext,
		"FloatToIntRoundDownConversion.CalculateInt for e")

	floatArg.Value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"FloatToIntRoundDownConversion.CalculateInt for big number")
}

func TestFloatToIntRoundUpConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, piTwoDigitPrecision)
	uut := NewFloatToIntRoundUpConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToIntRoundUpConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"FloatToIntRoundUpConversion.CalculateInt fails on uninitialized")

	uut.SetArg(floatArg)
	assertCalculatesToInt(
		t,
		uut,
		piRoundedUp,
		noGrammarContext,
		"FloatToIntRoundUpConversion.CalculateInt for pi")

	floatArg.Value = floatE
	assertCalculatesToInt(
		t,
		uut,
		eRoundedUp,
		noGrammarContext,
		"FloatToIntRoundUpConversion.CalculateInt for e")

	floatArg.Value = bigNumber
	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"FloatToIntRoundUpConversion.CalculateInt for big number")
}

func TestStringToIntConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, exampleIntValueAsString)
	uut := NewStringToIntConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringToIntConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"StringToIntConversion.CalculateInt fails on uninitialized")

	uut.SetArg(stringArg)
	assertCalculatesToInt(
		t,
		uut,
		exampleIntValue,
		noGrammarContext,
		"StringToIntConversion.CalculateInt some integer")

	stringArg.Value = bigNumberAsString
	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"StringToIntConversion.CalculateInt for big number")

	stringArg.Value = stringPi
	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"StringToIntConversion.CalculateInt for float number")
}

func TestStringToFloatConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, stringPiTwoDigitPrecision)
	uut := NewStringToFloatConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringToFloatConversion)

	assertCalculatesToFloatFails(
		t,
		uut,
		noGrammarContext,
		"StringToFloatConversion.CalculateFloat fails on uninitialized")

	uut.SetArg(stringArg)
	assertCalculatesToFloat(
		t,
		uut,
		piTwoDigitPrecision,
		noGrammarContext,
		"StringToFloatConversion.CalculateFloat for pi")

	stringArg.Value = exampleIntValueAsString
	assertCalculatesToFloat(
		t,
		uut,
		exampleIntValue,
		noGrammarContext,
		"StringToFloatConversion.CalculateFloat for int")

	stringArg.Value = exampleStringValue
	assertCalculatesToFloatFails(
		t,
		uut,
		noGrammarContext,
		"StringToFloatConversion.CalculateFloat for no number")
}

func TestFloatToStringConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, pi)
	uut := NewFloatToStringConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToStringConversion)

	assertCalculatesToStringFails(
		t,
		uut,
		noGrammarContext,
		"FloatToStringConversion.CalculateInt fails on uninitialized")

	uut.SetArg(floatArg)
	uut.Precision = 4
	assertCalculatesToString(
		t,
		uut,
		stringPi,
		noGrammarContext,
		"FloatToStringConversion.CalculateInt for pi")

	uut.Precision = 2
	assertCalculatesToString(
		t,
		uut,
		stringPiTwoDigitPrecision,
		noGrammarContext,
		"FloatToStringConversion.CalculateInt for e")

	uut.Precision = 10
	assertCalculatesToString(
		t,
		uut,
		stringPi,
		noGrammarContext,
		"FloatToStringConversion.CalculateInt for pi")

	uut.Precision = -1
	assertCalculatesToStringFails(
		t,
		uut,
		noGrammarContext,
		"FloatToStringConversion.CalculateInt for big number")
}

func TestIntToBoolConversion(t *testing.T) {
	intArg := NewIntConstant(variableName, 5)
	uut := NewIntToBoolConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntToBoolConversion)

	assertCalculatesToBoolFails(
		t,
		uut,
		noGrammarContext,
		"IntToBoolConversion.CalculateBool fails on uninitialized")

	uut.SetArg(intArg)
	assertCalculatesToBool(
		t,
		uut,
		true,
		noGrammarContext,
		"IntToBoolConversion.CalculateBool for non zero number")

	intArg.Value = 0
	assertCalculatesToBool(
		t,
		uut,
		false,
		noGrammarContext,
		"IntToBoolConversion.CalculateBool for zero")
}

func TestFloatToBoolConversion(t *testing.T) {
	floatArg := NewFloatConstant(variableName, 12.)
	uut := NewFloatToBoolConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatToBoolConversion)

	assertCalculatesToBoolFails(
		t,
		uut,
		noGrammarContext,
		"FloatToBoolConversion.CalculateBool fails on uninitialized")

	uut.SetArg(floatArg)
	assertCalculatesToBool(
		t,
		uut,
		true,
		noGrammarContext,
		"FloatToBoolConversion.CalculateBool for non zero number")

	floatArg.Value = 0.
	assertCalculatesToBool(
		t,
		uut,
		false,
		noGrammarContext,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.Value = math.NaN()
	assertCalculatesToBool(
		t,
		uut,
		true,
		noGrammarContext,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.Value = math.Inf(1)
	assertCalculatesToBool(
		t,
		uut,
		true,
		noGrammarContext,
		"FloatToBoolConversion.CalculateBool for zero")

	floatArg.Value = math.Inf(-1)
	assertCalculatesToBool(
		t,
		uut,
		true,
		noGrammarContext,
		"FloatToBoolConversion.CalculateBool for zero")
}

func TestStringToBoolConversion(t *testing.T) {
	stringArg := NewStringConstant(variableName, "true")
	uut := NewStringToBoolConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringToBoolConversion)

	assertCalculatesToBoolFails(
		t,
		uut,
		noGrammarContext,
		"StringToBoolConversion.CalculateBool fails on uninitialized")

	uut.SetArg(stringArg)
	assertCalculatesToBool(
		t,
		uut,
		true,
		noGrammarContext,
		"StringToBoolConversion.CalculateBool for 'true'")

	stringArg.Value = "false"
	assertCalculatesToBool(
		t,
		uut,
		false,
		noGrammarContext,
		"StringToBoolConversion.CalculateBool for 'false'")

	stringArg.Value = "Anything else"
	assertCalculatesToBoolFails(
		t,
		uut,
		noGrammarContext,
		"StringToBoolConversion.CalculateBool for anything else fails")
}

func TestBoolToIntConversion(t *testing.T) {
	boolArg := NewBoolConstant(variableName, true)
	uut := NewBoolToIntConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolToIntConversion)

	assertCalculatesToIntFails(
		t,
		uut,
		noGrammarContext,
		"IntToStringConversion.CalculateInt fails on uninitialized")

	uut.SetArg(boolArg)
	assertCalculatesToInt(
		t,
		uut,
		1,
		noGrammarContext,
		"IntToStringConversion.CalculateInt")

	boolArg.Value = false
	assertCalculatesToInt(
		t,
		uut,
		0,
		noGrammarContext,
		"IntToStringConversion.CalculateInt")
}

func TestBoolToFloatConversion(t *testing.T) {
	boolArg := NewBoolConstant(variableName, true)
	uut := NewBoolToFloatConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolToFloatConversion)

	assertCalculatesToFloatFails(
		t,
		uut,
		noGrammarContext,
		"BoolToFloatConversion.CalculateFloat fails on uninitialized")

	uut.SetArg(boolArg)
	assertCalculatesToFloat(
		t,
		uut,
		1.,
		noGrammarContext,
		"BoolToFloatConversion.CalculateFloat")

	boolArg.Value = false
	assertCalculatesToFloat(
		t,
		uut,
		0.,
		noGrammarContext,
		"BoolToFloatConversion.CalculateFloat")
}

func TestBoolToStringConversion(t *testing.T) {
	boolArg := NewBoolConstant(variableName, true)
	uut := NewBoolToStringConversion(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolToStringConversion)

	assertCalculatesToStringFails(
		t,
		uut,
		noGrammarContext,
		"BoolToStringConversion.CalculateString fails on uninitialized")

	uut.SetArg(boolArg)
	assertCalculatesToString(
		t,
		uut,
		"true",
		noGrammarContext,
		"BoolToStringConversion.CalculateString")

	boolArg.Value = false
	assertCalculatesToString(
		t,
		uut,
		"false",
		noGrammarContext,
		"BoolToStringConversion.CalculateString")
}
