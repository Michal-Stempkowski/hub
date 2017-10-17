package sheet_logic

import (
	"fmt"
	"hub/sheet_logic/sheet_logic_types"
)

const EmptyExpressionName string = "<none>"

type GrammarContext interface {
	GetIntValue(string) (int64, error)
	GetFloatValue(string) (float64, error)
	GetStringValue(string) (string, error)
	GetBoolValue(string) (bool, error)
}

type GrammarElement interface {
	GetType() sheet_logic_types.T
	GetName() string
	SetName(string)
}

type UnaryOperationInt interface {
	GetArg() IntExpresion
	SetArg(IntExpresion)
}

type UnaryOperationFloat interface {
	GetArg() FloatExpresion
	SetArg(FloatExpresion)
}

type UnaryOperationString interface {
	GetArg() StringExpresion
	SetArg(StringExpresion)
}

type UnaryOperationBool interface {
	GetArg() BoolExpresion
	SetArg(BoolExpresion)
}

type BinaryOperationInt interface {
	GetLeftArg() IntExpresion
	SetLeftArg(IntExpresion)

	GetRightArg() IntExpresion
	SetRightArg(IntExpresion)
}

type BinaryOperationFloat interface {
	GetLeftArg() FloatExpresion
	SetLeftArg(FloatExpresion)

	GetRightArg() FloatExpresion
	SetRightArg(FloatExpresion)
}

type BinaryOperationString interface {
	GetLeftArg() StringExpresion
	SetLeftArg(StringExpresion)

	GetRightArg() StringExpresion
	SetRightArg(StringExpresion)
}

type BinaryOperationBool interface {
	GetLeftArg() BoolExpresion
	SetLeftArg(BoolExpresion)

	GetRightArg() BoolExpresion
	SetRightArg(BoolExpresion)
}

// Because there will be also volatile sources (formulas, textfields etc.)
type IntExpresion interface {
	CalculateInt(GrammarContext) (int64, error)
}

type EmptyIntExpression struct {
	GrammarElement
}

func (e *EmptyIntExpression) CalculateInt(_ GrammarContext) (int64, error) {
	return 0, fmt.Errorf("%T.CalculateInt", e)
}

func NewEmptyIntExpression() *EmptyIntExpression {
	return &EmptyIntExpression{
		&emptyGrammarElementImpl{sheet_logic_types.EmptyIntExpression}}
}

type FloatExpresion interface {
	CalculateFloat(GrammarContext) (float64, error)
}

type EmptyFloatExpression struct {
	GrammarElement
}

func (e *EmptyFloatExpression) CalculateFloat(_ GrammarContext) (float64, error) {
	return 0, fmt.Errorf("%T.CalculateFloat", e)
}

func NewEmptyFloatExpression() *EmptyFloatExpression {
	return &EmptyFloatExpression{
		&emptyGrammarElementImpl{sheet_logic_types.EmptyFloatExpression}}
}

type StringExpresion interface {
	CalculateString(GrammarContext) (string, error)
}

type EmptyStringExpression struct {
	GrammarElement
}

func (e *EmptyStringExpression) CalculateString(_ GrammarContext) (string, error) {
	return "", fmt.Errorf("%T.CalculateString", e)
}

func NewEmptyStringExpression() *EmptyStringExpression {
	return &EmptyStringExpression{
		&emptyGrammarElementImpl{sheet_logic_types.EmptyStringExpression}}
}

type BoolExpresion interface {
	CalculateBool(GrammarContext) (bool, error)
}

type EmptyBoolExpression struct {
	GrammarElement
}

func (e *EmptyBoolExpression) CalculateBool(_ GrammarContext) (bool, error) {
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

type UnaryOperationIntImpl struct {
	arg IntExpresion
}

func (u *UnaryOperationIntImpl) GetArg() IntExpresion {
	return u.arg
}

func (u *UnaryOperationIntImpl) SetArg(newExpr IntExpresion) {
	u.arg = newExpr
}

func DefaultUnaryOperationIntImpl() UnaryOperationInt {
	return &UnaryOperationIntImpl{
		NewEmptyIntExpression()}
}

type UnaryOperationFloatImpl struct {
	arg FloatExpresion
}

func (u *UnaryOperationFloatImpl) GetArg() FloatExpresion {
	return u.arg
}

func (u *UnaryOperationFloatImpl) SetArg(newExpr FloatExpresion) {
	u.arg = newExpr
}

func DefaultUnaryOperationFloatImpl() UnaryOperationFloat {
	return &UnaryOperationFloatImpl{
		NewEmptyFloatExpression()}
}

type UnaryOperationStringImpl struct {
	arg StringExpresion
}

func (u *UnaryOperationStringImpl) GetArg() StringExpresion {
	return u.arg
}

func (u *UnaryOperationStringImpl) SetArg(newExpr StringExpresion) {
	u.arg = newExpr
}

func DefaultUnaryOperationStringImpl() UnaryOperationString {
	return &UnaryOperationStringImpl{
		NewEmptyStringExpression()}
}

type UnaryOperationBoolImpl struct {
	arg BoolExpresion
}

func (u *UnaryOperationBoolImpl) GetArg() BoolExpresion {
	return u.arg
}

func (u *UnaryOperationBoolImpl) SetArg(newExpr BoolExpresion) {
	u.arg = newExpr
}

func DefaultUnaryOperationBoolImpl() UnaryOperationBool {
	return &UnaryOperationBoolImpl{
		NewEmptyBoolExpression()}
}

type BinaryOperationIntImpl struct {
	argLeft  IntExpresion
	argRight IntExpresion
}

func (b *BinaryOperationIntImpl) GetLeftArg() IntExpresion {
	return b.argLeft
}

func (b *BinaryOperationIntImpl) GetRightArg() IntExpresion {
	return b.argRight
}

func (b *BinaryOperationIntImpl) SetLeftArg(newExpr IntExpresion) {
	b.argLeft = newExpr
}

func (b *BinaryOperationIntImpl) SetRightArg(newExpr IntExpresion) {
	b.argRight = newExpr
}

func DefaultBinaryOperationIntImpl() BinaryOperationInt {
	return &BinaryOperationIntImpl{
		NewEmptyIntExpression(),
		NewEmptyIntExpression()}
}

type BinaryOperationFloatImpl struct {
	argLeft  FloatExpresion
	argRight FloatExpresion
}

func (b *BinaryOperationFloatImpl) GetLeftArg() FloatExpresion {
	return b.argLeft
}

func (b *BinaryOperationFloatImpl) GetRightArg() FloatExpresion {
	return b.argRight
}

func (b *BinaryOperationFloatImpl) SetLeftArg(newExpr FloatExpresion) {
	b.argLeft = newExpr
}

func (b *BinaryOperationFloatImpl) SetRightArg(newExpr FloatExpresion) {
	b.argRight = newExpr
}

func DefaultBinaryOperationFloatImpl() BinaryOperationFloat {
	return &BinaryOperationFloatImpl{
		NewEmptyFloatExpression(),
		NewEmptyFloatExpression()}
}

type BinaryOperationStringImpl struct {
	argLeft  StringExpresion
	argRight StringExpresion
}

func (b *BinaryOperationStringImpl) GetLeftArg() StringExpresion {
	return b.argLeft
}

func (b *BinaryOperationStringImpl) GetRightArg() StringExpresion {
	return b.argRight
}

func (b *BinaryOperationStringImpl) SetLeftArg(newExpr StringExpresion) {
	b.argLeft = newExpr
}

func (b *BinaryOperationStringImpl) SetRightArg(newExpr StringExpresion) {
	b.argRight = newExpr
}

func DefaultBinaryOperationStringImpl() BinaryOperationString {
	return &BinaryOperationStringImpl{
		NewEmptyStringExpression(),
		NewEmptyStringExpression()}
}

type BinaryOperationBoolImpl struct {
	argLeft  BoolExpresion
	argRight BoolExpresion
}

func (b *BinaryOperationBoolImpl) GetLeftArg() BoolExpresion {
	return b.argLeft
}

func (b *BinaryOperationBoolImpl) GetRightArg() BoolExpresion {
	return b.argRight
}

func (b *BinaryOperationBoolImpl) SetLeftArg(newExpr BoolExpresion) {
	b.argLeft = newExpr
}

func (b *BinaryOperationBoolImpl) SetRightArg(newExpr BoolExpresion) {
	b.argRight = newExpr
}

func DefaultBinaryOperationBoolImpl() BinaryOperationBool {
	return &BinaryOperationBoolImpl{
		NewEmptyBoolExpression(),
		NewEmptyBoolExpression()}
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
