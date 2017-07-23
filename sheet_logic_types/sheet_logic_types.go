package sheet_logic_types

type T int64

const (
	// Constants
	IntConstant T = iota
	StringConstant
	FloatConstant
	BoolConstant
	// Conversions
	IntToStringConversion
	IntToFloatConversion
	StringToIntConversion
	FloatToIntConversion
	FloatToIntRoundDownConversion
	FloatToIntRoundUpConversion
)
