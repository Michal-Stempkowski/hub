package sheet_logic

import (
	"hub/framework"
	"hub/sheet_logic/sheet_logic_types"
	"testing"
)

const (
	variableName               = "VariableName"
	anotherVariableName        = "AnotherVariableName"
	exampleIntValue            = 5
	exampleIntValueAsString    = "5"
	exampleFloatValue          = 3.14
	exampleStringValue         = "StringValue"
	stringPi                   = "3.1416"
	stringPiTwoDigitPrecision  = "3.14"
	stringPiZeroDigitPrecision = "3."
	piRoundedDown              = 3
	piRoundedUp                = 4
	piTwoDigitPrecision        = 3.14
	pi                         = 3.1416
	stringE                    = "2.7182"
	floatE                     = 2.7182
	eRoundedUp                 = 3
	eRoundedDown               = 2
	bigNumber                  = 1 << 80
	bigNumberAsString          = "1.0e80"
)

func assertHasType(t *testing.T, el GrammarElement, expectedType sheet_logic_types.T) {
	grammar_type := el.GetType()
	if grammar_type != expectedType {
		t.Errorf("Grammar element type %v expected, received %v", expectedType, grammar_type)
	}
}

func assertHasName(t *testing.T, el GrammarElement, expectedName string) {
	grammar_name := el.GetName()
	if grammar_name != expectedName {
		t.Errorf("Grammar element name %v expected, received %v", expectedName, grammar_name)
	}
}

func assertCalculatesToInt(t *testing.T, expr IntExpresion, expected int64, trail string) {
	val, err := expr.CalculateInt()
	if err != nil {
		t.Errorf("%v: Int calculation has failed: %v", trail, err)
	}
	if val != expected {
		t.Errorf("%v: %v expected, received %v", trail, expected, val)
	}
}

func assertCalculatesToFloat(t *testing.T, expr FloatExpresion, expected float64, trail string) {
	val, err := expr.CalculateFloat()
	if err != nil {
		t.Errorf("%v: Float calculation has failed: %v", trail, err)
	}
	if !framework.FloatEq(val, expected) {
		t.Errorf("%v: %v expected, received %v", trail, expected, val)
	}
}

func assertCalculatesToIntFails(t *testing.T, expr IntExpresion, trail string) {
	_, err := expr.CalculateInt()
	if err == nil {
		t.Errorf("%v: Int calculation should fail!", trail)
	}
}

func assertCalculatesToFloatFails(t *testing.T, expr FloatExpresion, trail string) {
	_, err := expr.CalculateFloat()
	if err == nil {
		t.Errorf("%v: Float calculation should fail!", trail)
	}
}

func assertCalculatesToStringFails(t *testing.T, expr StringExpresion, trail string) {
	_, err := expr.CalculateString()
	if err == nil {
		t.Errorf("%v: String calculation should fail!", trail)
	}
}

func assertCalculatesToBoolFails(t *testing.T, expr BoolExpresion, trail string) {
	_, err := expr.CalculateBool()
	if err == nil {
		t.Errorf("%v: Bool calculation should fail!", trail)
	}
}

func assertCalculatesToString(t *testing.T, expr StringExpresion, expected string, trail string) {
	val, err := expr.CalculateString()
	if err != nil {
		t.Errorf("%v: String calculation has failed: %v", trail, err)
	}
	if val != expected {
		t.Errorf("%v: %v expected, received %v", trail, expected, val)
	}
}

func assertCalculatesToBool(t *testing.T, expr BoolExpresion, expected bool, trail string) {
	val, err := expr.CalculateBool()
	if err != nil {
		t.Errorf("%v: Bool calculation has failed: %v", trail, err)
	}
	if val != expected {
		t.Errorf("%v: %v expected, received %v", trail, expected, val)
	}
}

func grammarElementScenario(
	t *testing.T, uut GrammarElement, expectedType sheet_logic_types.T) {
	assertHasType(t, uut, expectedType)

	assertHasName(t, uut, variableName)
	uut.SetName(anotherVariableName)
	assertHasName(t, uut, anotherVariableName)
}

func emptyGrammarElementScenario(
	t *testing.T, uut GrammarElement, expectedType sheet_logic_types.T) {
	assertHasType(t, uut, expectedType)

	assertHasName(t, uut, EmptyExpressionName)
	uut.SetName(anotherVariableName)
	assertHasName(t, uut, EmptyExpressionName)
}

func TestEmptyIntExpression(t *testing.T) {
	uut := NewEmptyIntExpression()

	emptyGrammarElementScenario(t, uut.GrammarElement, sheet_logic_types.EmptyIntExpression)
	assertCalculatesToIntFails(t, uut, "EmptyIntExpression.CalculateInt")
}

func TestEmptyFloatExpression(t *testing.T) {
	uut := NewEmptyFloatExpression()

	emptyGrammarElementScenario(t, uut.GrammarElement, sheet_logic_types.EmptyFloatExpression)
	assertCalculatesToFloatFails(t, uut, "EmptyFloatExpression.CalculateFloat")
}

func TestEmptyStringExpression(t *testing.T) {
	uut := NewEmptyStringExpression()

	emptyGrammarElementScenario(t, uut.GrammarElement, sheet_logic_types.EmptyStringExpression)
	assertCalculatesToStringFails(t, uut, "EmptyStringExpression.CalculateString")
}

func TestEmptyBoolExpression(t *testing.T) {
	uut := NewEmptyBoolExpression()

	emptyGrammarElementScenario(t, uut.GrammarElement, sheet_logic_types.EmptyBoolExpression)
	assertCalculatesToBoolFails(t, uut, "EmptyBoolExpression.CalculateBool")
}
