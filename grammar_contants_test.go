package sheet_logic

import "testing"
import "hub/sheet_logic_types"

func TestShouldBeAbleToCreateIntConstant(t *testing.T) {
	uut := NewIntConstant(variableName, exampleIntValue)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.IntConstant)
	assertCalculatesToInt(t, uut, exampleIntValue, "IntConstant.CalculateInt")
}

func TestShouldBeAbleToCreateStringConstant(t *testing.T) {
	uut := NewStringConstant(variableName, exampleStringValue)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.StringConstant)
	assertCalculatesToString(t, uut, exampleStringValue, "StringConstant.CalculateString")
}

func TestShouldBeAbleToCreateFloatConstant(t *testing.T) {
	uut := NewFloatConstant(variableName, exampleFloatValue)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.FloatConstant)
	assertCalculatesToFloat(t, uut, exampleFloatValue, "FloatConstant.CalculateFloat")
}

func TestShouldBeAbleToCreateBoolConstant(t *testing.T) {
	uut := NewBoolConstant(variableName, true)

	grammarElementScenario(t, uut.grammarElementImpl, sheet_logic_types.BoolConstant)
	assertCalculatesToBool(t, uut, true, "BoolConstant.CalculateBool")
}
