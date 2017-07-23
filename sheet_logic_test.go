package sheet_logic

import "testing"
import "hub/sheet_logic_types"

const (
	variableName        = "VariableName"
	anotherVariableName = "AnotherVariableName"
	exampleIntValue     = 5
	exampleStringValue  = "StringValue"
)

func assertHasType(t *testing.T, el GrammarElement, expectedType sheet_logic_types.T) {
	grammar_type := el.GetType()
	if grammar_type != expectedType {
		t.Errorf("Grammar element type", expectedType, "expected, received", grammar_type)
	}
}

func assertHasName(t *testing.T, el GrammarElement, expectedName string) {
	grammar_name := el.GetName()
	if grammar_name != expectedName {
		t.Errorf("Grammar element name", expectedName, "expected, received", grammar_name)
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, trail string) {
	if a != b {
		t.Errorf(trail, ":", a, "expected, received", b)
	}
}

func TestShouldBeAbleToCreateIntConstant(t *testing.T) {
	uut := NewIntConstant(variableName, exampleIntValue)
	assertHasName(t, uut, variableName)
	assertEqual(t, uut.Value, exampleIntValue, "IntConstant value")
	assertHasType(t, uut, sheet_logic_types.IntConstant)

	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}

func TestShouldBeAbleToCreateStringConstant(t *testing.T) {
	uut := NewStringConstant(variableName, exampleStringValue)
	assertHasName(t, uut, variableName)
	assertEqual(t, uut.Value, exampleStringValue, "StringConstant value")
	assertHasType(t, uut, sheet_logic_types.StringConstant)

	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}
