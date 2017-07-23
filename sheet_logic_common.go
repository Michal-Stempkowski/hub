package sheet_logic

import (
	"fmt"
	"hub/sheet_logic_types"
)

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
	*grammarElementImpl
}

func (e *EmptyIntExpression) CalculateInt() (int64, error) {
	return 0, fmt.Errorf("%T.CalculateInt", e, e.CalculateInt)
}

func NewEmptyIntExpression() *EmptyIntExpression {
	return &EmptyIntExpression{
		&grammarElementImpl{"<none>", sheet_logic_types.EmptyIntExpression}}
}

type FloatExpresion interface {
	CalculateFloat() (float64, error)
}

type StringExpresion interface {
	CalculateString() (string, error)
}

type BoolExpresion interface {
	CalculateBool() (bool, error)
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
