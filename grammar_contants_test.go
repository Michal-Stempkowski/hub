package sheet_logic

import "testing"
import "hub/sheet_logic_types"

func TestShouldBeAbleToCreateIntConstant(t *testing.T) {
	uut := NewIntConstant(variableName, exampleIntValue)
	assertHasName(t, uut, variableName)
	assertCalculatesToInt(t, uut, exampleIntValue, "IntConstant.CalculateInt")
	assertHasType(t, uut, sheet_logic_types.IntConstant)

	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}

func TestShouldBeAbleToCreateStringConstant(t *testing.T) {
	uut := NewStringConstant(variableName, exampleStringValue)
	assertHasName(t, uut, variableName)
	assertCalculatesToString(t, uut, exampleStringValue, "StringConstant.CalculateString")
	assertHasType(t, uut, sheet_logic_types.StringConstant)

	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}

func TestShouldBeAbleToCreateFloatConstant(t *testing.T) {
	uut := NewFloatConstant(variableName, exampleFloatValue)
	assertHasName(t, uut, variableName)
	assertCalculatesToFloat(t, uut, exampleFloatValue, "FloatConstant.CalculateFloat")
	assertHasType(t, uut, sheet_logic_types.FloatConstant)

	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}
