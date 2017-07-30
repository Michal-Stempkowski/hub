package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntEquals(t *testing.T) {
	uut := NewIntEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntEquals)
	lesser, equal, greater := false, true, false
	IntComparatorScenario(t, (*IntComparator)(uut), "IntEquals", lesser, equal, greater)
}

func TestFloatEquals(t *testing.T) {
	uut := NewFloatEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatEquals)
	lesser, equal, greater := false, true, false
	FloatComparatorScenario(t, (*FloatComparator)(uut), "FloatEquals", lesser, equal, greater)
}

func TestBoolEquals(t *testing.T) {
	uut := NewBoolEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolEquals)

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewBoolConstant(variableName, false))
	assertCalculatesToBool(
		t,
		uut,
		false,
		"BoolEquals.CalculateBool different parameters")

	uut.SetRightArg(NewBoolConstant(variableName, true))
	assertCalculatesToBool(
		t,
		uut,
		true,
		"BoolEquals.CalculateBool same parameters")

	uut.SetLeftArg(NewEmptyBoolExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"BoolEquals.CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewEmptyBoolExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"BoolEquals.CalculateBool fails when argRight missing")
}

func TestStringEquals(t *testing.T) {
	uut := NewStringEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringEquals)

	uut.SetLeftArg(NewStringConstant(variableName, "abc"))
	uut.SetRightArg(NewStringConstant(variableName, "a"))
	assertCalculatesToBool(
		t,
		uut,
		false,
		"StringEquals.CalculateBool different parameters")

	uut.SetRightArg(NewStringConstant(variableName, "abc"))
	assertCalculatesToBool(
		t,
		uut,
		true,
		"StringEquals.CalculateBool same parameters")

	uut.SetLeftArg(NewEmptyStringExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"StringEquals.CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewStringConstant(variableName, "abc"))
	uut.SetRightArg(NewEmptyStringExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"StringEquals.CalculateBool fails when argRight missing")
}
