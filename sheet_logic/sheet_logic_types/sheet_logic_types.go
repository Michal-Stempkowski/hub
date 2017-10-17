package sheet_logic_types

type T int64

const (
	// Empty types
	EmptyIntExpression T = iota
	EmptyFloatExpression
	EmptyStringExpression
	EmptyBoolExpression

	// Constants
	IntConstant
	StringConstant
	FloatConstant
	BoolConstant

	// Conversions
	IntToStringConversion
	IntToFloatConversion
	IntToBoolConversion

	StringToIntConversion
	StringToFloatConversion
	StringToBoolConversion

	FloatToIntConversion
	FloatToIntRoundDownConversion
	FloatToIntRoundUpConversion
	FloatToStringConversion
	FloatToBoolConversion

	BoolToIntConversion
	BoolToStringConversion
	BoolToFloatConversion

	// Sum
	IntSum
	FloatSum
	StringSum

	// Difference
	IntDifference
	FloatDifference

	// Division
	IntDivision
	FloatDivision

	// Multiplication
	IntMultiplication
	FloatMultiplication

	// Modulo
	IntModulo

	// Floor
	FloatFloor

	// Ceil
	FloatCeil

	// Round
	FloatRound

	// Equals
	IntEquals
	FloatEquals
	BoolEquals
	StringEquals

	// Not equals
	IntNotEquals
	FloatNotEquals
	BoolNotEquals
	StringNotEquals

	// Lesser
	IntLesser
	FloatLesser
	StringLesser

	// Lesser or equal
	IntLesserEqual
	FloatLesserEqual
	StringLesserEqual

	// Greater
	IntGreater
	FloatGreater
	StringGreater

	// Greater or Equal
	IntGreaterEqual
	FloatGreaterEqual
	StringGreaterEqual

	// Logical
	Negation
	And
	Or

	// Input
	IntInput
)
