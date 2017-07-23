package sheet_logic_types

type T int64

const (
	IntConstant T = iota
	StringConstant
	FloatConstant
	IntToStringConversion
	IntToFloatConversion
	StringToIntConversion
	FloatToIntConversion
)
