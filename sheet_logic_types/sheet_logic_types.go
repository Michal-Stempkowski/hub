package sheet_logic_types

type T int64

const FloatPrecision = 0.000001

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
)
