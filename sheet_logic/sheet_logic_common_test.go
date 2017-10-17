package sheet_logic

import (
	"fmt"
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
	noContext                  = "No grammar context passed"
)

var noGrammarContext *dummyGrammarContext

type dummyGrammarContext struct{}

func (_ *dummyGrammarContext) GetIntValue(string) (int64, error) {
	return 0, fmt.Errorf(noContext)
}

func (_ *dummyGrammarContext) GetFloatValue(string) (float64, error) {
	return 0., fmt.Errorf(noContext)
}

func (_ *dummyGrammarContext) GetStringValue(string) (string, error) {
	return "", fmt.Errorf(noContext)
}

func (_ *dummyGrammarContext) GetBoolValue(string) (bool, error) {
	return false, fmt.Errorf(noContext)
}

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

func assertCalculatesToInt(
	t *testing.T, expr IntExpresion, expected int64, c GrammarContext, trail string) {

	val, err := expr.CalculateInt(c)
	if err != nil {
		t.Errorf("%v: Int calculation has failed: %v", trail, err)
	}
	if val != expected {
		t.Errorf("%v: %v expected, received %v", trail, expected, val)
	}
}

func assertCalculatesToFloat(
	t *testing.T, expr FloatExpresion, expected float64, c GrammarContext, trail string) {

	val, err := expr.CalculateFloat(c)
	if err != nil {
		t.Errorf("%v: Float calculation has failed: %v", trail, err)
	}
	if !framework.FloatEq(val, expected) {
		t.Errorf("%v: %v expected, received %v", trail, expected, val)
	}
}

func assertCalculatesToIntFails(
	t *testing.T, expr IntExpresion, c GrammarContext, trail string) {

	_, err := expr.CalculateInt(c)
	if err == nil {
		t.Errorf("%v: Int calculation should fail!", trail)
	}
}

func assertCalculatesToFloatFails(
	t *testing.T, expr FloatExpresion, c GrammarContext, trail string) {

	_, err := expr.CalculateFloat(c)
	if err == nil {
		t.Errorf("%v: Float calculation should fail!", trail)
	}
}

func assertCalculatesToStringFails(
	t *testing.T, expr StringExpresion, c GrammarContext, trail string) {

	_, err := expr.CalculateString(c)
	if err == nil {
		t.Errorf("%v: String calculation should fail!", trail)
	}
}

func assertCalculatesToBoolFails(
	t *testing.T, expr BoolExpresion, c GrammarContext, trail string) {

	_, err := expr.CalculateBool(c)
	if err == nil {
		t.Errorf("%v: Bool calculation should fail!", trail)
	}
}

func assertCalculatesToString(
	t *testing.T, expr StringExpresion, expected string, c GrammarContext, trail string) {

	val, err := expr.CalculateString(c)
	if err != nil {
		t.Errorf("%v: String calculation has failed: %v", trail, err)
	}
	if val != expected {
		t.Errorf("%v: %v expected, received %v", trail, expected, val)
	}
}

func assertCalculatesToBool(
	t *testing.T, expr BoolExpresion, expected bool, c GrammarContext, trail string) {

	val, err := expr.CalculateBool(c)
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
	assertCalculatesToIntFails(
		t, uut, noGrammarContext, "EmptyIntExpression.CalculateInt")
}

func TestEmptyFloatExpression(t *testing.T) {
	uut := NewEmptyFloatExpression()

	emptyGrammarElementScenario(t, uut.GrammarElement, sheet_logic_types.EmptyFloatExpression)
	assertCalculatesToFloatFails(
		t, uut, noGrammarContext, "EmptyFloatExpression.CalculateFloat")
}

func TestEmptyStringExpression(t *testing.T) {
	uut := NewEmptyStringExpression()

	emptyGrammarElementScenario(t, uut.GrammarElement, sheet_logic_types.EmptyStringExpression)
	assertCalculatesToStringFails(
		t, uut, noGrammarContext, "EmptyStringExpression.CalculateString")
}

func TestEmptyBoolExpression(t *testing.T) {
	uut := NewEmptyBoolExpression()

	emptyGrammarElementScenario(t, uut.GrammarElement, sheet_logic_types.EmptyBoolExpression)
	assertCalculatesToBoolFails(
		t, uut, noGrammarContext, "EmptyBoolExpression.CalculateBool")
}
