package sheet_logic

import "testing"

func IntComparatorScenario(
	t *testing.T,
	uut *IntComparator,
	g GrammarContext,
	uutTypeName string,
	forLesser,
	forEqual,
	forGreater bool) {

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewIntConstant(variableName, 3))
	assertCalculatesToBool(
		t,
		uut,
		forLesser,
		g,
		uutTypeName+".CalculateBool for lesser")

	uut.SetRightArg(NewIntConstant(variableName, 2))
	assertCalculatesToBool(
		t,
		uut,
		forEqual,
		g,
		uutTypeName+".CalculateBool same parameters")

	uut.SetRightArg(NewIntConstant(variableName, 1))
	assertCalculatesToBool(
		t,
		uut,
		forGreater,
		g,
		uutTypeName+".CalculateBool greater")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		g,
		uutTypeName+".CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		g,
		uutTypeName+".CalculateBool fails when argRight missing")
}

func FloatComparatorScenario(
	t *testing.T,
	uut *FloatComparator,
	g GrammarContext,
	uutTypeName string,
	forLesser,
	forEqual,
	forGreater bool) {

	uut.SetLeftArg(NewFloatConstant(variableName, 2.0))
	uut.SetRightArg(NewFloatConstant(variableName, 2.1))
	assertCalculatesToBool(
		t,
		uut,
		forLesser,
		g,
		uutTypeName+".CalculateBool for lesser")

	uut.SetRightArg(NewFloatConstant(variableName, 1.9+0.1))
	assertCalculatesToBool(
		t,
		uut,
		forEqual,
		g,
		uutTypeName+".CalculateBool same parameters")

	uut.SetRightArg(NewFloatConstant(variableName, 1.9))
	assertCalculatesToBool(
		t,
		uut,
		forGreater,
		g,
		uutTypeName+".CalculateBool greater")

	uut.SetLeftArg(NewEmptyFloatExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		g,
		uutTypeName+".CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewFloatConstant(variableName, 2.0))
	uut.SetRightArg(NewEmptyFloatExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		g,
		uutTypeName+".CalculateBool fails when argRight missing")
}

func BoolComparatorScenario(
	t *testing.T,
	uut *BoolComparator,
	g GrammarContext,
	uutTypeName string,
	expectedResults []bool) {
	if len(expectedResults) != 4 {
		t.Errorf("%s: BoolComparatorScenario - expectedResults should have 4 entries!")
	}

	uut.SetLeftArg(NewBoolConstant(variableName, false))
	uut.SetRightArg(NewBoolConstant(variableName, false))
	assertCalculatesToBool(
		t,
		uut,
		expectedResults[0],
		g,
		uutTypeName+".CalculateBool for false - false")

	uut.SetLeftArg(NewBoolConstant(variableName, false))
	uut.SetRightArg(NewBoolConstant(variableName, true))
	assertCalculatesToBool(
		t,
		uut,
		expectedResults[1],
		g,
		uutTypeName+".CalculateBool for false - true")

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewBoolConstant(variableName, false))
	assertCalculatesToBool(
		t,
		uut,
		expectedResults[2],
		g,
		uutTypeName+".CalculateBool for true - false")

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewBoolConstant(variableName, true))
	assertCalculatesToBool(
		t,
		uut,
		expectedResults[3],
		g,
		uutTypeName+".CalculateBool for true - true")

	uut.SetLeftArg(NewEmptyBoolExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		g,
		uutTypeName+".CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewEmptyBoolExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		g,
		uutTypeName+".CalculateBool fails when argRight missing")
}

func StringComparatorScenario(
	t *testing.T,
	uut *StringComparator,
	g GrammarContext,
	uutTypeName string,
	forLesser,
	forEqual,
	forGreater bool) {

	uut.SetLeftArg(NewStringConstant(variableName, "abc"))
	uut.SetRightArg(NewStringConstant(variableName, "abd"))
	assertCalculatesToBool(
		t,
		uut,
		forLesser,
		g,
		uutTypeName+".CalculateBool for lesser")

	uut.SetRightArg(NewStringConstant(variableName, "abc"))
	assertCalculatesToBool(
		t,
		uut,
		forEqual,
		g,
		uutTypeName+".CalculateBool same parameters")

	uut.SetRightArg(NewStringConstant(variableName, "abb"))
	assertCalculatesToBool(
		t,
		uut,
		forGreater,
		g,
		uutTypeName+".CalculateBool greater")

	uut.SetLeftArg(NewEmptyStringExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		g,
		uutTypeName+".CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewStringConstant(variableName, "abc"))
	uut.SetRightArg(NewEmptyStringExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		g,
		uutTypeName+".CalculateBool fails when argRight missing")
}
