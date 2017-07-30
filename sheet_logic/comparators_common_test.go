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

func FloatComparatorScenario(
	t *testing.T, uut *FloatComparator, uutTypeName string, forLesser, forEqual, forGreater bool) {
	uut.SetLeftArg(NewFloatConstant(variableName, 2.0))
	uut.SetRightArg(NewFloatConstant(variableName, 2.1))
	assertCalculatesToBool(
		t,
		uut,
		forLesser,
		uutTypeName+".CalculateBool for lesser")

	uut.SetRightArg(NewFloatConstant(variableName, 1.9+0.1))
	assertCalculatesToBool(
		t,
		uut,
		forEqual,
		uutTypeName+".CalculateBool same parameters")

	uut.SetRightArg(NewFloatConstant(variableName, 1.9))
	assertCalculatesToBool(
		t,
		uut,
		forGreater,
		uutTypeName+".CalculateBool greater")

	uut.SetLeftArg(NewEmptyFloatExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		uutTypeName+".CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewFloatConstant(variableName, 2.0))
	uut.SetRightArg(NewEmptyFloatExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		uutTypeName+".CalculateBool fails when argRight missing")
}

func BoolComparatorScenario(
	t *testing.T, uut *BoolComparator, uutTypeName string, expectedResults []bool) {
	if len(expectedResults) != 4 {
		t.Errorf("%s: BoolComparatorScenario - expectedResults should have 4 entries!")
	}

	uut.SetLeftArg(NewBoolConstant(variableName, false))
	uut.SetRightArg(NewBoolConstant(variableName, false))
	assertCalculatesToBool(
		t,
		uut,
		expectedResults[0],
		uutTypeName+".CalculateBool for false - false")

	uut.SetLeftArg(NewBoolConstant(variableName, false))
	uut.SetRightArg(NewBoolConstant(variableName, true))
	assertCalculatesToBool(
		t,
		uut,
		expectedResults[1],
		uutTypeName+".CalculateBool for false - true")

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewBoolConstant(variableName, false))
	assertCalculatesToBool(
		t,
		uut,
		expectedResults[2],
		uutTypeName+".CalculateBool for true - false")

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewBoolConstant(variableName, true))
	assertCalculatesToBool(
		t,
		uut,
		expectedResults[3],
		uutTypeName+".CalculateBool for true - true")

	uut.SetLeftArg(NewEmptyBoolExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		uutTypeName+".CalculateBool fails when argLeft missing")

	uut.SetLeftArg(NewBoolConstant(variableName, true))
	uut.SetRightArg(NewEmptyBoolExpression())
	assertCalculatesToBoolFails(
		t,
		uut,
		uutTypeName+".CalculateBool fails when argRight missing")
}
