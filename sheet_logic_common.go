package sheet_logic

import (
	"fmt"
	"hub/sheet_logic_types"
)

const EmptyExpressionName string = "<none>"

type GrammarElement interface {
	GetType() sheet_logic_types.T
	GetName() string
	SetName(string)
}

// Because there will be also volatile sources (formulas, textfields etc.)
type IntExpresion interface {
	CalculateInt() (int64, error)
}

type EmptyIntExpression struct {
	GrammarElement
}

func (e *EmptyIntExpression) CalculateInt() (int64, error) {
	return 0, fmt.Errorf("%T.CalculateInt", e)
}

func NewEmptyIntExpression() *EmptyIntExpression {
	return &EmptyIntExpression{
		&emptyGrammarElementImpl{sheet_logic_types.EmptyIntExpression}}
}

type FloatExpresion interface {
	CalculateFloat() (float64, error)
}

type EmptyFloatExpression struct {
	GrammarElement
}

func (e *EmptyFloatExpression) CalculateFloat() (float64, error) {
	return 0, fmt.Errorf("%T.CalculateFloat", e)
}

func NewEmptyFloatExpression() *EmptyFloatExpression {
	return &EmptyFloatExpression{
		&emptyGrammarElementImpl{sheet_logic_types.EmptyFloatExpression}}
}

type StringExpresion interface {
	CalculateString() (string, error)
}

type EmptyStringExpression struct {
	GrammarElement
}

func (e *EmptyStringExpression) CalculateString() (string, error) {
	return "", fmt.Errorf("%T.CalculateString", e)
}

func NewEmptyStringExpression() *EmptyStringExpression {
	return &EmptyStringExpression{
		&emptyGrammarElementImpl{sheet_logic_types.EmptyStringExpression}}
}

type BoolExpresion interface {
	CalculateBool() (bool, error)
}

type EmptyBoolExpression struct {
	GrammarElement
}

func (e *EmptyBoolExpression) CalculateBool() (bool, error) {
	return false, fmt.Errorf("%T.CalculateBool", e)
}

func NewEmptyBoolExpression() *EmptyBoolExpression {
	return &EmptyBoolExpression{
		&emptyGrammarElementImpl{sheet_logic_types.EmptyBoolExpression}}
}

type grammarElementImpl struct {
	name         string
	grammar_type sheet_logic_types.T
}

func (g *grammarElementImpl) GetName() string {
	return g.name
}

func (g *grammarElementImpl) SetName(newName string) {
	g.name = newName
}

func (g *grammarElementImpl) GetType() sheet_logic_types.T {
	return g.grammar_type
}

func getFirstError(errors ...error) error {
	for _, e := range errors {
		if e != nil {
			return e
		}
	}

	return nil
}

type emptyGrammarElementImpl struct {
	grammar_type sheet_logic_types.T
}

func (g *emptyGrammarElementImpl) GetName() string {
	return EmptyExpressionName
}

func (g *emptyGrammarElementImpl) SetName(string) {
	// Not doing anything conciously, no need to treat this call as error
}

func (g *emptyGrammarElementImpl) GetType() sheet_logic_types.T {
	return g.grammar_type
}
