package sheet_logic_types

type T int64

const FloatPrecision = 0.000001

const (
	// Constants
	IntConstant T = iota
	StringConstant
	FloatConstant
	BoolConstant
	// Conversions
	IntToStringConversion
	IntToFloatConversion
	IntToBoolConversion

	StringToIntConversion
	StringToFloatConversion

	FloatToIntConversion
	FloatToIntRoundDownConversion
	FloatToIntRoundUpConversion
	FloatToStringConversion
)
