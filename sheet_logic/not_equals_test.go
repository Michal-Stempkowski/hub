package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntNotEquals(t *testing.T) {
	uut := NewIntNotEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntNotEquals)
	lesser, equal, greater := true, false, true
	IntComparatorScenario(t, (*IntComparator)(uut), "IntNotEquals", lesser, equal, greater)
}

func TestFloatNotEquals(t *testing.T) {
	uut := NewFloatNotEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatNotEquals)
	lesser, equal, greater := true, false, true
	FloatComparatorScenario(t, (*FloatComparator)(uut), "FloatNotEquals", lesser, equal, greater)
}

func TestBoolNotEquals(t *testing.T) {
	uut := NewBoolNotEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.BoolNotEquals)

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewBoolConstant(variableName, false))
	assertCalculatesToBool(
		t,
		uut,
		true,
		"BoolNotEquals.CalculateBool different parameters")

	uut.SetRightArg(NewBoolConstant(variableName, true))
	assertCalculatesToBool(
		t,
		uut,
		false,
		"BoolNotEquals.CalculateBool same parameters")

	uut.SetLeftArg(NewEmptyBoolExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"BoolNotEquals.CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewEmptyBoolExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"BoolNotEquals.CalculateBool fails when argRight missing")
}

func TestStringNotEquals(t *testing.T) {
	uut := NewStringNotEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.StringNotEquals)

	uut.SetLeftArg(NewStringConstant(variableName, "abc"))
	uut.SetRightArg(NewStringConstant(variableName, "a"))
	assertCalculatesToBool(
		t,
		uut,
		true,
		"StringNotEquals.CalculateBool different parameters")

	uut.SetRightArg(NewStringConstant(variableName, "abc"))
	assertCalculatesToBool(
		t,
		uut,
		false,
		"StringNotEquals.CalculateBool same parameters")

	uut.SetLeftArg(NewEmptyStringExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"StringNotEquals.CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewStringConstant(variableName, "abc"))
	uut.SetRightArg(NewEmptyStringExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"StringNotEquals.CalculateBool fails when argRight missing")
}
