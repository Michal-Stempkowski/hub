package sheet_logic

import (
	"hub/sheet_logic_types"
	"math"
	"testing"
)

const (
	variableName              = "VariableName"
	anotherVariableName       = "AnotherVariableName"
	exampleIntValue           = 5
	exampleIntValueAsString   = "5"
	exampleFloatValue         = 3.14
	exampleStringValue        = "StringValue"
	stringPi                  = "3.1416"
	stringPiTwoDigitPrecision = "3.14"
	piRoundedDown             = 3
	piRoundedUp               = 4
	piTwoDigitPrecision       = 3.14
	stringE                   = "2.7182"
	floatE                    = 2.7182
	eRoundedUp                = 3
	eRoundedDown              = 2
	bigNumber                 = 1 << 80
	bigNumberAsString         = "1.0e80"
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

func assertCalculatesToInt(t *testing.T, expr IntExpresion, expected int64, trail string) {
	val, err := expr.CalculateInt()
	if err != nil {
		t.Errorf(trail, ":", "Int calculation has failed: ", err)
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
	if diff := math.Abs(val - expected); diff >= sheet_logic_types.FloatPrecision {
		t.Errorf(trail, ":", expected, "expected, received", val)
	}
}

func assertCalculatesToIntFails(t *testing.T, expr IntExpresion, trail string) {
	_, err := expr.CalculateInt()
	if err == nil {
		t.Errorf(trail, ":", "Int calculation should fail!")
	}
}

func assertCalculatesToFloatFails(t *testing.T, expr FloatExpresion, trail string) {
	_, err := expr.CalculateFloat()
	if err == nil {
		t.Errorf(trail, ":", "Int calculation should fail!")
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

func assertCalculatesToBool(t *testing.T, expr BoolExpresion, expected bool, trail string) {
	val, err := expr.CalculateBool()
	if err != nil {
		t.Errorf(trail, ":", "Bool calculation has failed: ", err)
	}
	if val != expected {
		t.Errorf(trail, ":", expected, "expected, received", val)
	}
}

func grammarElementScenario(
	t *testing.T, uut *grammarElementImpl, expectedType sheet_logic_types.T) {
	assertHasType(t, uut, expectedType)

	assertHasName(t, uut, variableName)
	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}
