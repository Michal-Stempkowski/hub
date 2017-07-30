package sheet_logic

import "testing"

func IntComparatorScenario(
	t *testing.T, uut *IntComparator, uutTypeName string, forLesser, forEqual, forGreater bool) {
	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewIntConstant(variableName, 3))
	assertCalculatesToBool(
		t,
		uut,
		forLesser,
		uutTypeName+".CalculateBool for lesser")

	uut.SetRightArg(NewIntConstant(variableName, 2))
	assertCalculatesToBool(
		t,
		uut,
		forEqual,
		uutTypeName+".CalculateBool same parameters")

	uut.SetRightArg(NewIntConstant(variableName, 1))
	assertCalculatesToBool(
		t,
		uut,
		forGreater,
		uutTypeName+".CalculateBool greater")

	uut.SetLeftArg(NewEmptyIntExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		uutTypeName+".CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewIntConstant(variableName, 2))
	uut.SetRightArg(NewEmptyIntExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		uutTypeName+".CalculateBool fails when argRight missing")
}
