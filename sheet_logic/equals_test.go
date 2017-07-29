package sheet_logic

import (
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

func TestIntEquals(t *testing.T) {
	uut := NewIntEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.IntEquals)

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewIntConstant(variableName, 3))
	assertCalculatesToBool(
		t,
		uut,
		false,
		"IntEquals.CalculateBool different parameters")

	uut.SetRightArg(NewIntConstant(variableName, 2))
	assertCalculatesToBool(
		t,
		uut,
		true,
		"IntEquals.CalculateBool same parameters")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"IntEquals.CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"IntEquals.CalculateBool fails when argRight missing")
}

func TestFloatEquals(t *testing.T) {
	uut := NewFloatEquals(variableName)
	grammarElementScenario(t, uut.GrammarElement, sheet_logic_types.FloatEquals)

	uut.SetLeftArg(NewFloatConstant(variableName, 2.1))
	uut.SetRightArg(NewFloatConstant(variableName, 2.0))
	assertCalculatesToBool(
		t,
		uut,
		false,
		"FloatEquals.CalculateBool different parameters")

	uut.SetRightArg(NewFloatConstant(variableName, 2.1))
	assertCalculatesToBool(
		t,
		uut,
		true,
		"FloatEquals.CalculateBool same parameters")

	uut.SetLeftArg(NewEmptyFloatExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"FloatEquals.CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewFloatConstant(variableName, 2))
	uut.SetRightArg(NewEmptyFloatExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		"FloatEquals.CalculateBool fails when argRight missing")
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
