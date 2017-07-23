package sheet_logic

import "testing"
import "hub/sheet_logic_types"

const (
	variableName            = "VariableName"
	anotherVariableName     = "AnotherVariableName"
	exampleIntValue         = 5
	exampleIntValueAsString = "5"
	exampleFloatValue       = 3.14
	exampleStringValue      = "StringValue"
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

func assertCalculatesToInt(t *testing.T, expr IntExpresion, expected int, trail string) {
	val, err := expr.CalculateInt()
	if err != nil {
		t.Errorf(trail, ":", "Int calculation has failed: ", err)
	}
	if val != expected {
		t.Errorf(trail, ":", expected, "expected, received", val)
	}
}

func assertCalculatesToString(t *testing.T, expr StringExpresion, expected string, trail string) {
	val, err := expr.CalculateString()
	if err != nil {
		t.Errorf(trail, ":", "String calculation has failed: ", err)
	}
	if val != expected {
		t.Errorf(trail, ":", expected, "expected, received", val)
	}
}

func assertCalculatesToFloat(t *testing.T, expr FloatExpresion, expected float64, trail string) {
	val, err := expr.CalculateFloat()
	if err != nil {
		t.Errorf(trail, ":", "Float calculation has failed: ", err)
	}
	if val != expected {
		t.Errorf(trail, ":", expected, "expected, received", val)
	}
}
